package util

// SSHConfig SSH连接配置
type SSHConfig struct {
	IP        string // 主机IP
	Port      int    // SSH端口
	Type      int    // 认证类型:1->密码,2->密钥
	Username  string // 用户名
	Password  string // 密码(type=1时使用)
	PublicKey string // 公钥(type=2时使用)
}
