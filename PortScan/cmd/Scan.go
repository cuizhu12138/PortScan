/*
Copyright © 2023 sjx <540643428@qq.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ScanCmd represents the Scan command
var ScanCmd = &cobra.Command{
	Use:   "Scan",
	Short: "进行端口扫描",
	Long: `以多种方式进行端口扫描，目前支持以下几种方式
	Connect连接扫描
	Syn扫描`,
	Run: /*nil,*/ func(cmd *cobra.Command, args []string) {
		fmt.Println("您没有选择扫描模式,将默认以connect模式进行扫描,修改扫描模式请输入 'PortScan Scan -h' 获取更多帮助")
		cCmd.Run(cmd, args)
	},
	Args: cobra.RangeArgs(1, 2),
}

func init() {
	rootCmd.AddCommand(ScanCmd)

	ScanCmd.Aliases = []string{"S"}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ScanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ScanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 处理没有flag的情况
/*
// 判断扫描模式
		if flagc != "" {
			WayToScan = "Connect连接"
			IPNeedToScan = flagc
		} else if flags != "" {
			WayToScan = "Syn扫描"
			IPNeedToScan = flags
		} else if flagx != "" {
			WayToScan = "Xmas Tree扫描"
			IPNeedToScan = flagx
		} else if flagf != "" {
			WayToScan = "Fin扫描"
			IPNeedToScan = flagf
		} else if flagn != "" {
			WayToScan = "Null扫描"
			IPNeedToScan = flagn
		} else if flaga != "" {
			WayToScan = "Ack扫描"
			IPNeedToScan = flaga
		} else if flagw != "" {
			WayToScan = "Window窗口扫描"
			IPNeedToScan = flagw
		} else if flagu != "" {
			WayToScan = "UDP扫描"
			IPNeedToScan = flagu
		} else {
			WayToScan = "Connect连接(默认)"
			if len(args) != 1 {
				fmt.Println("IP参数个数错误")
				return
			} else {
				IPNeedToScan = args[0]
			}
		}

		// 解析IP
		parseIP := net.ParseIP(IPNeedToScan)

		if parseIP == nil {
			fmt.Println("无效的IP地址")
			return
		}

		IPArray := parseIP.To4()
		if IPArray == nil {
			fmt.Println("IP地址转化错误")
			return
		}


ScanCmd.Flags().StringVarP(&flagc, "Connect", "c", "", "进行连接扫描")
	ScanCmd.Flags().StringVarP(&flags, "SYN", "s", "", "进行Syn扫描")
	ScanCmd.Flags().StringVarP(&flagc, "Xmas Tree", "x", "", "进行Xmas Tree扫描")
	ScanCmd.Flags().StringVarP(&flagc, "Fin", "f", "", "进行Fin扫描")
	ScanCmd.Flags().StringVarP(&flagc, "Null", "n", "", "进行Null扫描")
	ScanCmd.Flags().StringVarP(&flagc, "Ack", "a", "", "进行Ack扫描")
	ScanCmd.Flags().StringVarP(&flagc, "Window", "w", "", "进行Window窗口扫描")
	ScanCmd.Flags().StringVarP(&flagc, "UDP", "u", "", "进行UDP扫描")

var (
	flagc string // Connect连接扫描
	flags string // Syn扫描
	flagx string // Xmas Tree扫描
	flagf string // Fin扫描
	flagn string // Null扫描
	flaga string // Ack扫描
	flagw string // Window窗口扫描
	flagu string // UDP扫描

	WayToScan    string
	IPNeedToScan string
)


var AllScanMethod = []string{"Connect连接", "Syn扫描", "Xmas Tree扫描", "Fin扫描", "Null扫描", "Ack扫描", "Window窗口扫描", "UDP扫描"}

*/
