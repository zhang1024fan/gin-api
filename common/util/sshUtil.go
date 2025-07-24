package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHUtil struct{}

// NewSSHUtil 创建SSHUtil实例
func NewSSHUtil() *SSHUtil {
	return &SSHUtil{}
}

// ExecuteRemoteCommand 执行远程命令并返回输出
func (s *SSHUtil) ExecuteRemoteCommand(auth *SSHConfig, command string) (string, error) {
	config, err := s.getSSHConfig(auth)
	if err != nil {
		return "", fmt.Errorf("failed to create SSH config: %v", err)
	}

	// 建立SSH连接
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", auth.IP, auth.Port), config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	// 创建会话
	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// 设置输出缓冲区
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// 执行命令
	err = session.Run(command)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	return stdoutBuf.String(), nil
}

// getSSHConfig 根据认证类型创建SSH配置
func (s *SSHUtil) getSSHConfig(auth *SSHConfig) (*ssh.ClientConfig, error) {
	var authMethod ssh.AuthMethod
	var err error

	switch auth.Type {
	case 1: // 密码认证
		authMethod = ssh.Password(auth.Password)
	case 2: // 密钥认证
		authMethod, err = s.publicKeyAuth(auth.PublicKey)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported authentication type")
	}
	
	return &ssh.ClientConfig{
		User: auth.Username,
		Auth: []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil // 忽略主机密钥检查
		}),
		Timeout: 30 * time.Second,
	}, nil
}

// publicKeyAuth 使用公钥创建认证方法
func (s *SSHUtil) publicKeyAuth(publicKey string) (ssh.AuthMethod, error) {
	// 尝试解析私钥
	signer, err := ssh.ParsePrivateKey([]byte(publicKey))
	if err != nil {
		// 尝试添加密码为空的情况
		signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(publicKey), []byte(""))
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key (with/without passphrase): %v", err)
		}
	}
	return ssh.PublicKeys(signer), nil
}

// ExecuteScript 执行远程脚本
func (s *SSHUtil) ExecuteScript(auth *SSHConfig, script string) (string, error) {
	return s.ExecuteRemoteCommand(auth, script)
}

// TerminalLogin 建立SSH终端连接
func (s *SSHUtil) TerminalLogin(auth *SSHConfig) (*ssh.Client, error) {
	config, err := s.getSSHConfig(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to create SSH config: %v", err)
	}

	return ssh.Dial("tcp", fmt.Sprintf("%s:%d", auth.IP, auth.Port), config)
}

// UploadFile 上传文件到远程主机
func (s *SSHUtil) UploadFile(auth *SSHConfig, localPath, remotePath string) error {
	client, err := s.TerminalLogin(auth)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// 使用scp命令上传文件
	cmd := fmt.Sprintf("scp %s %s@%s:%s", localPath, auth.Username, auth.IP, remotePath)
	_, err = s.ExecuteRemoteCommand(auth, cmd)
	return err
}

// GetSystemInfo 获取远程主机系统信息
func (s *SSHUtil) GetSystemInfo(auth *SSHConfig) (map[string]string, error) {
	// 1. 测试基本连接
	if _, err := s.ExecuteRemoteCommand(auth, "echo 'Testing SSH connection...'"); err != nil {
		return nil, fmt.Errorf("SSH connection test failed: %v", err)
	}

	// 2. 创建并执行脚本
	output, err := s.ExecuteRemoteCommand(auth, getScriptContent())
	if err != nil {
		return nil, fmt.Errorf("failed to execute script: %v", err)
	}

	// 3. 解析结果
	var info map[string]string
	if err := json.Unmarshal([]byte(output), &info); err != nil {
		return nil, fmt.Errorf("invalid script output: %v (output: %s)", err, output)
	}

	// 打印调试信息
	fmt.Printf("Script output: %s\nParsed info: %+v\n", output, info)

	// 4. 验证必要字段
	requiredFields := []string{"privateIp", "os", "cpu", "memory", "disk", "name"}
	var missingFields []string
	for _, field := range requiredFields {
		if _, ok := info[field]; !ok {
			missingFields = append(missingFields, field)
		}
	}
	if len(missingFields) > 0 {
		return info, fmt.Errorf("missing required fields: %v (partial info: %+v)", missingFields, info)
	}

	return info, nil
}

// getScriptContent 返回可直接执行的脚本命令
func getScriptContent() string {
	return `#!/bin/bash

# 获取系统信息
privateIp=$(hostname -I | awk '{print $1}' || echo "unknown")
publicIp=$(curl -s ipinfo.io/ip 2>/dev/null || echo "")
os=$(cat /etc/os-release 2>/dev/null | grep PRETTY_NAME | cut -d= -f2 | tr -d '"' | sed 's/ LTS//;s/ //g' || echo "unknown")

# 获取CPU核心数(去掉单位)
cpu=$(nproc 2>/dev/null || echo "unknown")

# 获取内存大小(去掉单位)
memory=$(free -m 2>/dev/null | awk '/^Mem:/{printf "%.0f\n",$2/1024}' || echo "unknown")

# 获取磁盘总容量(去掉单位)
disk=$(df -h 2>/dev/null | awk '/\/$/ {print $2}' | sed 's/G//' || echo "unknown")

# 确保获取主机名称
name=$(hostname 2>/dev/null)
if [ -z "$name" ]; then
    name=$(uname -n 2>/dev/null || echo "unknown")
fi

# 输出JSON结果
echo '{"privateIp":"'$privateIp'","publicIp":"'$publicIp'","os":"'$os'","cpu":"'$cpu'","memory":"'$memory'","disk":"'$disk'","name":"'$name'"}'
`
}
