package gtkui

import (
	"fmt"
	"github.com/Ericwyn/EzeFormat/ipc"
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/EzeFormat/utils/gotkutils"
	"github.com/Ericwyn/EzeFormat/utils/pathutils"
	"github.com/Ericwyn/EzeFormat/utils/strutils"
	"github.com/gotk3/gotk3/gtk"
	"strings"
	"time"
)

var textView *gtk.TextView

func StartApp() {
	log.I("open gui by gtk")
	if trySendMessage(ipc.IpcMessagePing) {
		// 如果已经有其他翻译进程的话, 就发送一下消息，然后退出就好了
		sendSocketMessage(false)
		return
	}

	// 开启 server 监听来自其他进程的翻译请求
	startUnixSocketServer()
	// 此处需要异步，需要等 app 界面起来之后再去做消息发送
	go sendSocketMessage(true)

	OpenNewApp()
}

func OpenNewApp() {
	gtk.Init(nil)
	win := initWindows("EzeFormat")

	// 输入框
	inputBox, tv := initInputBox()
	textView = tv

	//btn 框框
	btnBoxLineSmart := initBtnBoxWithWrapBox([]gotkutils.BtnDefine{
		{
			Name:    "SmartFormatBtn",
			Label:   "智能解析",
			OnClick: FormatSmartFunc,
			Width:   150,
		},
		{
			Name:    "CompressFormatBtn",
			Label:   "智能压缩",
			OnClick: CompressSmartFunc,
			Width:   150,
		},
		{
			Name:    "TimeNowBtn",
			Label:   "当前时间",
			OnClick: TimeNowFunc,
			Width:   150,
		},
	})

	// btn 框框
	btnBoxLineJson := initBtnBox([]gotkutils.BtnDefine{
		{
			Name:    "FormatJsonBtn",
			Label:   "JSON 解析",
			OnClick: FormatJsonFunc,
			Width:   150,
		},
		{
			Name:    "CompressJsonBtn",
			Label:   "JSON 压缩",
			OnClick: CompressJsonFunc,
			Width:   150,
		},
		{
			Name:    "FormatXmlBtn",
			Label:   "XML 解析",
			OnClick: FormatJsonFunc,
			Width:   150,
		},
		{
			Name:    "CompressXmlBtn",
			Label:   "XML 压缩",
			OnClick: CompressJsonFunc,
			Width:   150,
		},
	})

	// 整个 box
	containerBox := gotkutils.NewBox(gtk.ORIENTATION_VERTICAL)
	containerBox.SetName("appContainer")
	containerBox.Add(inputBox)
	containerBox.Add(btnBoxLineSmart)
	containerBox.Add(btnBoxLineJson)
	//containerBox.Add(btnBoxLineXml)
	win.Add(containerBox)

	win.ShowAll()

	err := runCss()
	if err != nil {
		log.E("run css error")
		log.E(err)
	}

	gtk.StatusbarNew()

	gtk.Main()
}

func initWindows(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.E("Unable to create window:", err)
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetSizeRequest(800, 600)
	win.SetResizable(false)
	win.SetPosition(gtk.WIN_POS_CENTER)

	iconPath := pathutils.GetRunnerPath() + "/res-static/icon/json.ico"
	log.I("runPath: ", pathutils.GetRunnerPath(), ", icon: "+iconPath)

	err = win.SetIconFromFile(iconPath)

	//iconImg, err := gdk.PixbufNewFromFile(iconPath)
	if err != nil {
		log.E("set icon error", err)
	}
	//win.SetIcon(iconImg)
	win.SetIconName("EzeFormat")
	win.SetName("EzeFormat")

	return win
}

func initInputBox() (*gtk.Box, *gtk.TextView) {
	// 输入框
	// 外层的 inputScrollWin
	inputScrollWin, err := gtk.ScrolledWindowNew(nil, nil)
	inputScrollWin.SetName("inputScrollWin")
	if err != nil {
		log.E("create inputScrollWin error")
	}
	inputScrollWin.SetSizeRequest(800, 600)
	// 里面的 textview
	tv, err := gtk.TextViewNew()
	tv.SetName("inputTextView")

	if err != nil {
		log.E("Unable to create TextView:", err)
	}
	inputScrollWin.Add(tv)

	inputBox := gotkutils.NewBox(gtk.ORIENTATION_HORIZONTAL)
	inputBox.Add(inputScrollWin)

	return inputBox, tv
}

func initBtnBoxWithWrapBox(btnList []gotkutils.BtnDefine) *gtk.Box {
	btnBox := gotkutils.NewBox(gtk.ORIENTATION_HORIZONTAL)
	for _, btn := range btnList {
		btnBox.Add(gotkutils.NewBtn(btn))
	}

	wrapCheckBtn := gotkutils.NewCheckBtn(gotkutils.CheckBtnDefine{
		Name:  "",
		Label: "换行",
		OnClick: func(wrap bool) {
			log.I("change wrap set:", wrap)
			if wrap {
				setTextViewWrap(gtk.WRAP_WORD_CHAR)
			} else {
				setTextViewWrap(gtk.WRAP_NONE)
			}
		},
	})
	btnBox.Add(wrapCheckBtn)

	return btnBox
}

func initBtnBox(btnList []gotkutils.BtnDefine) *gtk.Box {
	btnBox := gotkutils.NewBox(gtk.ORIENTATION_HORIZONTAL)
	for _, btn := range btnList {
		btnBox.Add(gotkutils.NewBtn(btn))
	}
	return btnBox
}

// SetInputFunc 设置输入框
func SetInputFunc(str string) {
	if textView == nil {
		return
	}
	gotkutils.SetTextViewInput(textView, str)
}

// FormatSmartFunc 智能格式化
func FormatSmartFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.FormatSmart(input))
}

// CompressSmartFunc 智能格式化
func CompressSmartFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.CompressSmart(input))
}

// TimeNowFunc 展示当前时间戳
func TimeNowFunc() {
	if textView == nil {
		return
	}
	//input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.FormatType(fmt.Sprint(time.Now().UnixMilli()), strutils.TypeTimeStampMills))
}

// FormatJsonFunc 格式化
func FormatJsonFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.FormatType(input, strutils.TypeJson))
}

// CompressJsonFunc 压缩
func CompressJsonFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.CompressType(input, strutils.TypeJson))
}

// FormatXmlFunc 格式化
func FormatXmlFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.FormatType(input, strutils.TypeXml))
}

// CompressXmlFunc 压缩
func CompressXmlFunc() {
	if textView == nil {
		return
	}
	input := gotkutils.GetTextViewInput(textView)
	gotkutils.SetTextViewInput(textView, strutils.CompressType(input, strutils.TypeXml))
}

// 清空
func cleanTextViewFunc() {
	if textView == nil {
		return
	}
	gotkutils.SetTextViewInput(textView, "")
}

func jsonPreCheck(input string) bool {
	if textView == nil {
		return false
	}
	input = strings.Trim(input, " ")
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\r\n")
	if input == "" || strings.HasPrefix(input, "【错误】") {
		return false
	}
	if (strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]")) ||
		(strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}")) {
		return true
	} else {
		gotkutils.SetTextViewInput(textView, "【错误】: 这看着不太像 JSON ? \n\n"+input)
		return false
	}
}

func setTextViewWrap(warpMode gtk.WrapMode) {
	if textView == nil {
		return
	}
	textView.SetWrapMode(warpMode)
}
