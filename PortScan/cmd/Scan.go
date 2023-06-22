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
	Long:  `以多种方式进行端口扫描，目前只支持Connect连接扫描`,
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
