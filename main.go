package main

import (
	"bufio"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/a9886907/sillyGirl/develop/core"
	"github.com/a9886907/sillyGirl/utils"
)

var sillyGirl = core.MakeBucket("sillyGirl")

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc
	core.Init()
	if sillyGirl.GetBool("anti_kasi") {
		go utils.MonitorGoroutine()
	}
	port := sillyGirl.GetString("port", "8080")
	if core.HttpPort != "" {
		port = core.HttpPort
	}
	logs.Info("Http服务已运行(%s)。", sillyGirl.GetString("port", "8080"))
	go core.Server.Run("0.0.0.0:" + port)
	logs.Info("命令怎么失效了？关注频道 https://t.me/kczz2021 获取最新消息，发言关注微信公众号送喜官获取暗号。")
	d := false
	for _, arg := range os.Args {
		if arg == "-d" {
			d = true
		}
	}
	if !d {
		t := false
		for _, arg := range os.Args {
			if arg == "-t" {
				t = true
			}
		}
		if t {
			logs.Info("终端交互已启用。")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				data := scanner.Text()
				f := &core.Faker{
					Type:    "terminal",
					Message: string(data),
					Admin:   true,
				}
				core.Senders <- f
			}
			logs.Info("终端交互异常,请检查运行环境设置,如果是docker环境,请附加-it参数")
		} else {
			logs.Info("终端交互不可用，运行带-t参数即可启用。")
		}
	}

	select {}
}
