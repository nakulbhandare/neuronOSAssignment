package cmd

import (
	"net"
	"os"
	"os/exec"
	"time"
)

// Commander defined methods for commanding
type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

// PingResult structure for Ping commmad response
type PingResult struct {
	Successful bool
	Time       time.Duration
}

// SystemInfo structure for systeminfo command response
type SystemInfo struct {
	Hostname  string
	IPAddress string
}

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(host string) (PingResult, error) {
	start := time.Now()
	err := exec.Command("ping", "-c", "1", host).Run()
	if err != nil {
		return PingResult{Successful: false}, err
	}
	return PingResult{Successful: true, Time: time.Since(start)}, nil
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return SystemInfo{}, err
	}

	var ip string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip = ipNet.IP.String()
			break
		}
	}

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: ip,
	}, nil
}
