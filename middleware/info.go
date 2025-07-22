package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	PrivateIp string `json:"privateIp"`
	PublicIp  string `json:"publicIp"`
	Os        string `json:"os"`
	Cpu       string `json:"cpu"`
	Memory    string `json:"memory"`
	Disk      string `json:"disk"`
}

// 获取一个有效的 IPv4 私有 IP（优先 192.168.x.x，其次是 172.x.x.x）
func getPrivateIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	var ips []string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	// 自定义排序：192.168.x.x 优先，其次 172.x.x.x
	sort.Slice(ips, func(i, j int) bool {
		a := ips[i]
		b := ips[j]

		a192 := strings.HasPrefix(a, "192.168.")
		b192 := strings.HasPrefix(b, "192.168.")
		a172 := strings.HasPrefix(a, "172.")
		b172 := strings.HasPrefix(b, "172.")

		if a192 && !b192 {
			return true
		}
		if !a192 && b192 {
			return false
		}
		if a172 && !b172 {
			return true
		}
		if !a172 && b172 {
			return false
		}
		return false
	})

	if len(ips) > 0 {
		return ips[0]
	}
	return ""
}

// 获取公网 IP（使用 ipify）
func getPublicIP() string {
	cmd := exec.Command("curl", "-s", "http://ifconfig.io")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// 获取系统信息（如 Ubuntu22.04.2）
func getOSInfo() string {
	cmd := exec.Command("lsb_release", "-d")
	out, err := cmd.Output()
	if err != nil {
		return runtime.GOOS
	}
	output := strings.TrimSpace(string(out))
	parts := strings.SplitN(output, ":", 2)
	if len(parts) == 2 {
		desc := strings.TrimSpace(parts[1])
		desc = strings.ReplaceAll(desc, "Ubuntu ", "Ubuntu")
		desc = strings.ReplaceAll(desc, " LTS", "")
		return desc
	}
	return "unknown"
}

// 获取 CPU 核心数（逻辑核心）
func getCpuInfo() string {
	cores, _ := cpu.Counts(true)
	return fmt.Sprintf("%dC", cores)
}

// 获取内存总量（单位 GB，按 1000 进制）
func getMemInfo() string {
	vmStat, _ := mem.VirtualMemory()
	totalGB := math.Round(float64(vmStat.Total) / (1000 * 1000 * 1000))
	return fmt.Sprintf("%dG", int64(totalGB))
}
// 获取磁盘总容量（根分区，单位 GB，按 1000 进制）
func getDiskInfo() string {
	cmd := exec.Command("grep", "vda$", "/proc/partitions")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	// 输出示例： "  252       0 62914560 vda"
	var major, minor, blocks int64
	var name string
	fmt.Sscanf(string(out), "%d %d %d %s", &major, &minor, &blocks, &name)

	// blocks 是以 KB 为单位的
	totalGB := blocks / 1048576 // 转换为 GB
	return fmt.Sprintf("%dG", totalGB)
}
func main() {
	info := SystemInfo{
		PrivateIp: getPrivateIP(),
		PublicIp:  getPublicIP(),
		Os:        getOSInfo(),
		Cpu:       getCpuInfo(),
		Memory:    getMemInfo(),
		Disk:      getDiskInfo(),
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}

	fmt.Println(string(jsonData))
}
