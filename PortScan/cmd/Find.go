/*
Copyright © 2023 sjx <540643428@qq.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	flagi  string // ICMP主机发现
	flagaa string // ARP主机发现
)

// FindCmd represents the Find command
var FindCmd = &cobra.Command{
	Use:   "Find",
	Short: "进行主机发现",
	Long: `以多种方式进行主机发现，目前支持以下几种方式
	ICMP主机发现
	ARP主机发现`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在进行主机发现，默认以ICMP模式扫描，修改主机发现模式请输入 'PortScan Find -h' 获取帮助")
	},
}

func init() {
	rootCmd.AddCommand(FindCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// FindCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// FindCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
