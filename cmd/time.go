package cmd

import (
	"github.com/spf13/cobra"
	"ht/internal/timer"
	"log"
)

var timeCmd = &cobra.Command {
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run:func(cmd *cobra.Command,args []string){
		nowTime := timer.GetNowTime()
		log.Printf("当前时间：%s 时间戳：%d",nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())
	},
}