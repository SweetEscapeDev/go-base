package pkg

import (
	"fmt"
	"net"
)

func GetInstanceIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Errorf("fatal error: can't read instance IP: %w", err))
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	panic(fmt.Errorf("fatal error: can't read instance IP: %w", err))
}
