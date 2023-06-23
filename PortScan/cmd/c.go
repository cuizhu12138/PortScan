/*
Copyright © 2023 sjx <540643428@qq.com>
*/
package cmd

import (
	"PortScan/DataProcessing"
	"PortScan/ScanMethod"
	"PortScan/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "进行connect连接扫描",
	Long:  `用法：PortScan S c IP起点-IP终点 端口起点-端口终点(默认0 - 1024)`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp2 := config.MaxGoroutine
		if OccurSimultaneously != "" {
			tmp, err := strconv.ParseInt(OccurSimultaneously, 10, 32)

			if err != nil {
				fmt.Println("无效的并发数")
				os.Exit(114516)
			}
			config.MaxGoroutine = int(tmp)
		}
		defer func() { config.MaxGoroutine = tmp2 }()
		// 解析IP和端口
		if err := ScanMethod.ConnectScan(DataProcessing.ParseIPPort(args)); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

	},
	Args: cobra.RangeArgs(1, 2),
}

func init() {
	ScanCmd.AddCommand(cCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
