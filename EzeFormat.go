package main

import (
	"flag"
	"github.com/Ericwyn/EzeFormat/fyneui"
)

var xclipFlag = flag.Bool("x", false, "从剪贴板获取输入数据")

func main() {
	flag.Parse()
	fyneui.StartApp(*xclipFlag)
}
