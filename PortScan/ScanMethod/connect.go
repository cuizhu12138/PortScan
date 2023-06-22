package ScanMethod

import (
	"PortScan/DataProcessing"
	"PortScan/FindMethod"
	"PortScan/config"
	"fmt"
	"net"
	"sort"
)

var (
	CtotalIP   []net.IP
	CtotalPort []int

	CportNum int
	CipNum   int
)

// Scanner 扫描器
func CScanner(ports, results chan int, IP string) {
	for p := range ports {
		address := fmt.Sprintf(IP+":%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0 //端口关闭将发送0
			continue
		}
		conn.Close()
		results <- p //端口开启将发送端口号
	}
}

func CRealScan(ip net.IP) {

	var FinAns []int

	// 限制最大并发数
	PortPool := make(chan int, config.MaxGoroutine)
	Results := make(chan int)

	// 创造扫描器
	for i := 0; i < cap(PortPool); i++ {
		go CScanner(PortPool, Results, ip.String())
	}

	// 开一个协程把需要扫描的端口号放进端口池
	go func() {
		for _, port := range CtotalPort {
			PortPool <- port
		}
	}()

	// 记录打开的端口
	for i := 0; i < CportNum; i++ {
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

func ConnectScan(IPRange []string, PortRange []string) error {

	CtotalIP = DataProcessing.EnumerateIP(IPRange)
	CtotalPort = DataProcessing.EnumeratePort(PortRange)

	CportNum = len(CtotalPort)
	CipNum = len(CtotalIP)

	fmt.Printf("共需要扫描%d个IP,每个IP%d个端口\n", CipNum, CportNum)

	for _, IP := range CtotalIP {
		fmt.Println("正在扫描" + IP.String())

		if FindMethod.IcmpFindHost(IP) {
			CRealScan(IP)
		}
	}

	return nil
}
