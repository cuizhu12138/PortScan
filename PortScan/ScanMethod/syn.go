package ScanMethod

import (
	"PortScan/DataProcessing"
	"PortScan/config"
	"fmt"
	"net"
	"sort"
	"syscall"
	"time"
)

const (
	protocolICMP = 1
	protocolTCP  = 6
)

var (
	StotalIP   []net.IP
	StotalPort []int

	SportNum int
	SipNum   int
)

func getLocalIP(targetIP string) (net.IP, error) {
	conn, err := net.Dial("udp", targetIP+":80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func sendSYN(targetIP string, port int, timeout time.Duration) error {
	// 创建原始套接字
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_IP)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	// 设置超时时间
	tv := syscall.NsecToTimeval(timeout.Nanoseconds())
	if err := syscall.SetsockoptTimeval(fd, syscall.SOL_SOCKET, 0x1006, &tv); err != nil {
		return err
	}

	// 获取本地 IP 地址
	localIP, err := getLocalIP(targetIP)
	if err != nil {
		return err
	}

	// 构建 TCP 头部
	tcpHeader := make([]byte, 20)
	tcpHeader[0] = localIP[12] // 源 IP 地址
	tcpHeader[1] = localIP[13]
	tcpHeader[2] = localIP[14]
	tcpHeader[3] = localIP[15]
	tcpHeader[4] = 0                // 目标 IP 地址
	tcpHeader[5] = 0                // 目标 IP 地址
	tcpHeader[6] = 0                // 目标 IP 地址
	tcpHeader[7] = 0                // 目标 IP 地址
	tcpHeader[8] = 0x00             // 伪首部协议类型（TCP）
	tcpHeader[9] = 0x06             // 伪首部协议类型（TCP）
	tcpHeader[10] = 0               // 源端口
	tcpHeader[11] = 0               // 源端口
	tcpHeader[12] = byte(port >> 8) // 目标端口
	tcpHeader[13] = byte(port)
	tcpHeader[14] = 0    // 序列号
	tcpHeader[15] = 0    // 序列号
	tcpHeader[16] = 0    // 序列号
	tcpHeader[17] = 0    // 序列号
	tcpHeader[18] = 0    // 确认号
	tcpHeader[19] = 0    // 确认号
	tcpHeader[20] = 0x02 // SYN 标志位
	tcpHeader[21] = 0    // 窗口大小
	tcpHeader[22] = 0    // 窗口大小
	tcpHeader[23] = 0    // 校验和
	tcpHeader[24] = 0    // 校验和
	tcpHeader[25] = 0    // 紧急指针
	tcpHeader[26] = 0    // 紧急指针

	// 发送 SYN 数据包
	remoteAddr := syscall.SockaddrInet4{Port: port}
	copy(remoteAddr.Addr[:], net.ParseIP(targetIP).To4())
	err = syscall.Sendto(fd, tcpHeader, 0, &remoteAddr)
	if err != nil {
		return err
	}

	// 等待响应
	recvBuf := make([]byte, 1024)
	_, _, err = syscall.Recvfrom(fd, recvBuf, 0)
	if err != nil {
		return err
	}

	return nil
}

// Scanner 扫描器
func SScanner(ports, results chan int, IP string) {
	//for p := range ports {
	//address := fmt.Sprintf(IP+":%d", p)
	//conn, err := net.Dial("tcp", address)
	//if err != nil {
	//	results <- 0 //端口关闭将发送0
	//	continue
	//}
	//conn.Close()
	//results <- p //端口开启将发送端口号

	//}
	for p := range ports {
		targetIP := IP
		port := p
		timeout := 3 * time.Second

		err := sendSYN(targetIP, port, timeout)
		if err != nil {
			results <- 0
			fmt.Println(err)
			//fmt.Printf("端口 %d 关闭\n", port)
		} else {
			results <- p
			//fmt.Printf("端口 %d 开放\n", port)
		}
	}

}

func SRealScan(ip net.IP) {

	var FinAns []int

	// 限制最大并发数
	PortPool := make(chan int, config.MaxGoroutine)
	Results := make(chan int)

	// 创造扫描器
	for i := 0; i < cap(PortPool); i++ {
		go SScanner(PortPool, Results, ip.String())
	}

	// 开一个协程把需要扫描的端口号放进端口池
	go func() {
		for _, port := range StotalPort {
			PortPool <- port
		}
	}()

	// 记录打开的端口
	for i := 0; i < SportNum; i++ {
		port := <-Results
		if port != 0 {
			FinAns = append(FinAns, port)
		}
	}

	sort.Ints(FinAns)

	fmt.Printf("共有%d个端口开放\n", len(FinAns))

	for _, ans := range FinAns {
		fmt.Printf("%s:%d is open !\n", ip, ans)
	}
}

func SynScan(IPRange []string, PortRange []string) error {

	StotalIP = DataProcessing.EnumerateIP(IPRange)
	StotalPort = DataProcessing.EnumeratePort(PortRange)

	SportNum = len(StotalPort)
	SipNum = len(StotalIP)

	fmt.Printf("共需要扫描%d个IP,每个IP%d个端口\n", SipNum, SportNum)

	for _, IP := range StotalIP {
		fmt.Println("正在扫描" + IP.String())
		SRealScan(IP)
	}

	return nil
}
