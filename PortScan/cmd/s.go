/*
Copyright © 2023 sjx <540643428@qq.com>
*/
package cmd

import (
	"PortScan/DataProcessing"
	"PortScan/ScanMethod"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "进行syn连接扫描",
	Long:  `用法：PortScan S s IP起点-IP终点 端口起点-端口终点(默认0 - 1024)`,
	Run: func(cmd *cobra.Command, args []string) {

		// 解析IP和端口
		if err := ScanMethod.SynScan(DataProcessing.ParseIPPort(args)); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
	Args: cobra.RangeArgs(1, 2),
}

func init() {
	ScanCmd.AddCommand(sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
