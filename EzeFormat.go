package main

import (
	"flag"
	"fmt"
	"github.com/Ericwyn/EzeFormat/cmd"
	"github.com/Ericwyn/EzeFormat/conf"
	"github.com/Ericwyn/EzeFormat/fyneui"
	"os"
)

var xclipFlag = flag.Bool("x", false, "从剪贴板获取输入数据")

var versionFlag = flag.Bool("v", false, "查看版本号")

var inputTxt = flag.String("txt", "", "命令行模式, 直接输入的文本进行格式化")
var inputType = flag.String("type", "smart", "命令行模式, 直接输入的文本类型：json, xml, smart(智能推测)")

func main() {
	flag.Parse()

	if *inputTxt != "" {
		cmd.Format(*inputTxt, *inputType)
		return
	}

	if *versionFlag {
		fmt.Println(conf.Version)
		os.Exit(0)
	}

	fyneui.StartApp(*xclipFlag)
}
