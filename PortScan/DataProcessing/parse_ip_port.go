package DataProcessing

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func IsValidIP(IP string) bool {
	ip := net.ParseIP(IP)
	if ip == nil {
		return false
	} else {
		return true
	}
}

func IsValidPort(Port string) bool {
	num, err := strconv.ParseInt(Port, 10, 32)
	if err != nil {
		fmt.Println("无效的端口号:", err)
	} else {
		if num < 0 || num > 65535 {
			return false
		}
	}
	return true
}

// ParseIPPort 提供参数进行IP和Port的解析
func ParseIPPort(args []string) (IPArray, PortArray []string) {
	// 解析IP地址
	IPArray = strings.Split(args[0], "-")

	// 判断IP不是范围的情况
	if len(IPArray) != 2 {
		IPArray = append(IPArray, IPArray[0])
	}

	// 检查IP合法性
	for _, IP := range IPArray {
		if !IsValidIP(IP) {
			fmt.Println("无效的IP地址")
			os.Exit(114514)
		}
	}

	// 解析端口号
	if len(args) == 2 {
		PortArray = strings.Split(args[1], "-")

		// 判断端口号不是范围的情况
		if len(PortArray) != 2 {
			PortArray = append(PortArray, PortArray[0])
		}
	} else {
		// 没有提供端口号则默认0 - 1024
		PortArray = []string{"0", "1024"}
	}

	// 检查端口号合法性
	for _, Port := range PortArray {
		if !IsValidPort(Port) {
			fmt.Println("无效的端口号")
			os.Exit(114515)
		}
	}

	return IPArray, PortArray
}
