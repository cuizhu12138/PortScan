package FindMethod

import (
	"net"
	"time"
)

const (
	protocolICMP   = 1
	icmpEcho       = 8
	icmpTimestamp  = 13
	icmpHeaderSize = 8
)

func checkSum(msg []byte) uint16 {
	sum := 0

	// 将报文按16位进行拆分求和
	for i := 0; i < len(msg)-1; i += 2 {
		sum += int(msg[i])<<8 + int(msg[i+1])
	}

	// 如果报文长度为奇数，处理最后一个字节
	if len(msg)%2 == 1 {
		sum += int(msg[len(msg)-1]) << 8
	}

	// 将进位相加
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16

	// 取反得到校验和
	checksum := uint16(^sum)
	return checksum
}

func IcmpFindHost(IP net.IP) bool {

	conn, err := net.Dial("ip4:icmp", IP.String())
	if err != nil {
		return false
	}
	defer conn.Close()

	// 构造ICMP报文
	icmpMsg := make([]byte, icmpHeaderSize)
	icmpMsg[0] = icmpTimestamp // 类型
	icmpMsg[1] = 0             // 代码
	icmpMsg[2] = 0             // 校验和（先填充0）
	icmpMsg[3] = 0             // 校验和（先填充0）
	icmpMsg[4] = 0             // 标识符
	icmpMsg[5] = 0             // 标识符
	icmpMsg[6] = 0             // 序列号
	icmpMsg[7] = 0             // 序列号

	// 计算校验和
	checksum := checkSum(icmpMsg)
	icmpMsg[2] = byte(checksum >> 8)
	icmpMsg[3] = byte(checksum & 0xff)

	// 发送ICMP报文
	_, err = conn.Write(icmpMsg)
	if err != nil {
		return false
	}

	// 设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// 接收响应
	recvBuf := make([]byte, 1024)
	_, err = conn.Read(recvBuf)
	if err != nil {
		return false
	}

	return true
}
