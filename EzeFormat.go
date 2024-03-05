package main

import (
	"flag"
	"fmt"
	"github.com/Ericwyn/EzeFormat/conf"
	"github.com/Ericwyn/EzeFormat/fyneui"
	"os"
)

var xclipFlag = flag.Bool("x", false, "从剪贴板获取输入数据")

var versionFlag = flag.Bool("v", false, "查看版本号")

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(conf.Version)
		os.Exit(0)
	}

	fyneui.StartApp(*xclipFlag)
}
