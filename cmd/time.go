package cmd

import (
	"github.com/spf13/cobra"
	"ht/internal/timer"
	"log"
	"strconv"
	"time"
)

var location, _ = time.LoadLocation("Asia/Shanghai")
func FormatToDate(t time.Time,l *time.Location) string {
	return t.In(l).Format("2006-01-02 15:04:05")
}
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		ft, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatalf("错误：%v", err)
		}
		switch ft {
		case "ts": // time stamp
			intput := args[0]
			layout := "2006-01-02 15:04:05"
			t, err := time.ParseInLocation(layout, intput, location)
			if err != nil {
				log.Fatalf("入参错误: %v", err)
			}
			log.Printf("解析日期: %s 时间戳: %d ", FormatToDate(time.Unix(t.Unix(),0), location), t.Unix())
		case "dt":
			input := args[0]
			sec, err := strconv.ParseInt(input,10,64)
			if err != nil {
				log.Fatalf("时间戳入参错误：%v", err)
			}
			log.Printf("解析日期为: %s",FormatToDate(time.Unix(sec, 0),location))
		default:
			log.Printf("------当前时间戳: %d", nowTime.Unix())
		}
	},
}

func init() {
	timeCmd.Flags().StringP("format", "f", "", "时间戳")
}
