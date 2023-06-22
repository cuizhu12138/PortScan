package DataProcessing

import (
	"net"
	"strconv"
)

// 枚举下一个有效的IP地址
func nextIP(ip net.IP) net.IP {
	nextIP := make(net.IP, len(ip))
	copy(nextIP, ip)

	for i := len(nextIP) - 1; i >= 0; i-- {
		nextIP[i]++
		if nextIP[i] > 0 {
			break
		}
	}

	return nextIP
}

// EnumerateIP 枚举所有IP
func EnumerateIP(IPRange []string) (TotalIP []net.IP) {

	// 防止倒置输入死循环
	if IPRange[0] > IPRange[1] {
		IPRange[0], IPRange[1] = IPRange[1], IPRange[0]
	}

	StartIP := net.ParseIP(IPRange[0])
	EndIP := net.ParseIP(IPRange[1])

	for {
		//fmt.Println(StartIP)
		TotalIP = append(TotalIP, StartIP)
		if StartIP.Equal(EndIP) {
			break
		}
		StartIP = nextIP(StartIP)
	}

	return TotalIP
}

// EnumeratePort 枚举端口号
func EnumeratePort(PortRange []string) (TotalPort []int) {
	StartPort, _ := strconv.ParseInt(PortRange[0], 10, 32)
	EndPort, _ := strconv.ParseInt(PortRange[1], 10, 32)

	// 防止倒置输入死循环
	if StartPort > EndPort {
		StartPort, EndPort = EndPort, StartPort
	}

	for {
		//fmt.Println(StartPort)
		TotalPort = append(TotalPort, int(StartPort))
		if StartPort == EndPort {
			break
		}
		StartPort++
	}

	return TotalPort
}
