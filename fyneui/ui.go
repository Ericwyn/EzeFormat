package fyneui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Ericwyn/EzeFormat/fyneui/resource"
	"github.com/Ericwyn/EzeFormat/ipc"
	"github.com/Ericwyn/EzeFormat/utils/format"
	"github.com/Ericwyn/EzeFormat/utils/strutils"
	"github.com/Ericwyn/EzeFormat/utils/xclip"
	"time"
)

var version = "V1.0.5"

var mainApp fyne.App

var homeWindow fyne.Window

var homeNoteLabel *widget.Label

var historyTextLabel *widget.Label

var homeInputBox *EzeMultiLineEntry

var useSocket = true

// 编辑历史
var historyList = make([]string, 0)

func StartApp(useXclipData bool) {

	if useSocket {
		if trySendMessage(ipc.IpcMessagePing) {
			// 如果已经有其他翻译进程的话, 就发送一下消息，然后退出就好了
			sendSocketMessage(false)
			return
		}

		// 开启 server 监听来自其他进程的翻译请求
		startUnixSocketServer()
		// 此处需要异步，需要等 app 界面起来之后再去做消息发送
		go sendSocketMessage(true)
	} else {
		if useXclipData {
			go func() {
				// sleep 0.5s
				time.Sleep(time.Millisecond * 300)
				// 获取滑词然后格式化翻译
				getSelectTextAndSmartFormat()
			}()
		}
	}

	ShowMainUi()

}

func ShowMainUi() {
	mainApp = app.New()
	//mainApp.SetIcon(resource.ResourceIcon())
	mainApp.Settings().SetTheme(&resource.CustomerTheme{})

	homeWindow = mainApp.NewWindow("EzeFormat")

	homeWindow.Resize(fyne.Size{
		Width: 550,
		//Height: 600,
	})
	homeWindow.CenterOnScreen()

	// --------------------
	homeInputBox = NewEzeMultiLineEntry(fyne.Size{Height: 550})
	//homeInputBox.Resize()

	homeNoteLabel = widget.NewLabel("")

	//authorNote := widget.NewLabel(" Source: https://github.com/Ericwyn/EzeFormat 【 " + version + "】")
	authorNote := widget.NewRichTextFromMarkdown("[https://github.com/Ericwyn/EzeFormat](https://github.com/Ericwyn/EzeFormat)     **" + version + "**")

	//hello := widget.NewLabel("Hello Fyne!")
	homeWindow.SetContent(container.NewVBox(
		BtnPanelLineHistory(),
		homeInputBox,
		homeNoteLabel,
		BtnPanelLine1(),
		BtnPanelLine2(),
		authorNote,
	))

	homeWindow.ShowAndRun()
}

// BtnPanelLineHistory
// 智能解析, 智能压缩,当前时间, 换行
func BtnPanelLineHistory() *fyne.Container {
	leftSpacer := canvas.NewRectangle(theme.BackgroundColor())
	leftSpacer.SetMinSize(fyne.NewSize(3, 0)) // 设置固定宽度为20，高度为0表示不限制

	historyTextLabel = widget.NewLabel("0/0")

	return container.NewHBox(
		leftSpacer,
		newEzeButtonWithMargin("后退", func() {
			// 展示前一个历史记录
			changePreviousHistory()
		}),

		newEzeButtonWithMargin("前进", func() {
			// 展示下一个历史记录
			changeNextHistory()
		}),
		leftSpacer,
		historyTextLabel,
	)

}

// BtnPanelLine1
// 智能解析, 智能压缩,当前时间, 换行
func BtnPanelLine1() *fyne.Container {
	leftSpacer := canvas.NewRectangle(theme.BackgroundColor())
	leftSpacer.SetMinSize(fyne.NewSize(3, 0)) // 设置固定宽度为20，高度为0表示不限制

	return container.NewHBox(
		leftSpacer,
		newEzeButtonWithMargin("智能解析", func() {
			smartFormatFunc()
		}),

		newEzeButtonWithMargin("智能压缩", func() {
			smartCompressFunc()
		}),

		newEzeButtonWithMargin("当前时间", func() {
			timeNowFunc()
		}),

		newEzeButtonWithMargin("复制结果", func() {
			copyTextFunc()
		}),

		newEzeButtonWithMargin("清除输入", func() {
			homeInputBox.SetText("")
		}),
	)

}

func BtnPanelLine2() *fyne.Container {
	leftSpacer := canvas.NewRectangle(theme.BackgroundColor())
	leftSpacer.SetMinSize(fyne.NewSize(3, 0)) // 设置固定宽度为20，高度为0表示不限制

	return container.NewHBox(
		leftSpacer,
		newEzeButtonWithMargin("JSON 解析", func() {
			jsonFormatFunc()
		}),

		newEzeButtonWithMargin("JSON 压缩", func() {
			jsonCompressFunc()
		}),

		newEzeButtonWithMargin("XML 解析", func() {
			xmlFormatFunc()
		}),

		newEzeButtonWithMargin("XML 压缩", func() {
			xmlCompressFunc()
		}),

		widget.NewCheck("自动换行", func(b bool) {
			homeInputBox.ToggleWrap()
		}),
	)
}

// --------------------------- 自定义按钮，可以设置最小尺寸

type EzeButton struct {
	widget.Button
	minSize fyne.Size
}

// MinSize returns the size that this widget should not shrink below
func (b *EzeButton) MinSize() fyne.Size {
	return b.minSize
}

func newEzeButton(label string, onClick func()) *EzeButton {

	btn := &EzeButton{}
	btn.ExtendBaseWidget(btn)
	btn.Text = label
	btn.OnTapped = onClick

	btn.minSize = fyne.Size{
		Width:  90,
		Height: 35,
	}

	return btn
}

func newEzeButtonWithMargin(label string, onClick func()) *fyne.Container {
	btn := newEzeButton(label, onClick)

	var margin float32 = 5

	//// 创建透明的矩形作为间隔
	//leftSpacer := canvas.NewRectangle(theme.BackgroundColor())
	//leftSpacer.SetMinSize(fyne.NewSize(0, 0)) // 设置固定宽度为20，高度为0表示不限制

	rightSpacer := canvas.NewRectangle(theme.BackgroundColor())
	rightSpacer.SetMinSize(fyne.NewSize(margin, 0)) // 同上

	marginBox := container.NewHBox(
		//leftSpacer, // 左边的外边距
		btn,
		rightSpacer, // 右边的外边距
	)
	return marginBox
}

// --------------------------- 可定义尺寸的自定义多行文本输入框

type EzeMultiLineEntry struct {
	widget.Entry
	minSize fyne.Size
}

func (entry *EzeMultiLineEntry) ToggleWrap() {
	if entry.Wrapping == fyne.TextWrapBreak {
		entry.Wrapping = fyne.TextWrapOff // 切换到不换行
	} else {
		entry.Wrapping = fyne.TextWrapBreak // 切换到自动换行
	}
	entry.Refresh() // 刷新控件以应用更改
}

// NewEzeMultiLineEntry creates a new instance of EzeMultiLineEntry with a specified minimum size.
// It ensures the embedded Entry is properly initialized.
func NewEzeMultiLineEntry(minSize fyne.Size) *EzeMultiLineEntry {
	entry := &EzeMultiLineEntry{minSize: minSize}
	entry.ExtendBaseWidget(entry) // This is crucial for ensuring the widget is properly setup
	entry.MultiLine = true        // Enable multi-line support
	return entry
}

// MinSize returns the custom minimum size of the entry.
func (e *EzeMultiLineEntry) MinSize() fyne.Size {
	originalMinSize := e.Entry.MinSize()
	if e.minSize.Width > originalMinSize.Width {
		originalMinSize.Width = e.minSize.Width
	}
	if e.minSize.Height > originalMinSize.Height {
		originalMinSize.Height = e.minSize.Height
	}
	return originalMinSize
}

// ---------------------------------------------------------------------
// ui 与数据绑定

func smartFormatFunc() {
	result, err := format.SmartFormat(homeInputBox.Text)
	setNoteMsg("智能格式化失败: ", err)

	SetFormatResult(result)
}

// smartCompressFunc 智能格式化
func smartCompressFunc() {
	result, err := format.SmartCompress(homeInputBox.Text)
	setNoteMsg("智能压缩失败: ", err)

	SetFormatResult(result)
}

// timeNowFunc 展示当前时间戳
func timeNowFunc() {
	result, err := format.FormatType(fmt.Sprint(time.Now().UnixMilli()), format.TypeTimeStampMills)
	setNoteMsg("展示当前时间戳失败: ", err)

	SetFormatResult(result)
}

func copyTextFunc() {
	xclip.SetClipboard(homeInputBox.Text)
}

// jsonFormatFunc 格式化
func jsonFormatFunc() {
	result, err := format.FormatType(homeInputBox.Text, format.TypeJson)
	setNoteMsg("JSON 格式化失败: ", err)

	SetFormatResult(result)
}

// jsonCompressFunc 压缩
func jsonCompressFunc() {
	result, err := format.CompressType(homeInputBox.Text, format.TypeJson)
	setNoteMsg("JSON 压缩失败: ", err)

	SetFormatResult(result)
}

// xmlFormatFunc 格式化
func xmlFormatFunc() {
	result, err := format.FormatType(homeInputBox.Text, format.TypeXml)
	setNoteMsg("Xml 格式化失败: ", err)

	SetFormatResult(result)
}

// xmlCompressFunc 压缩
func xmlCompressFunc() {
	result, err := format.CompressType(homeInputBox.Text, format.TypeXml)
	setNoteMsg("Xml 压缩失败: ", err)

	SetFormatResult(result)
}

func setNoteMsg(prefix string, err error) {
	if err != nil {
		homeNoteLabel.SetText(" " + prefix + err.Error())
	} else {
		homeNoteLabel.SetText("")
	}
}

// --------------------------------------

func getSelectTextAndSmartFormat() {
	selectText := xclip.GetSelection()

	selectText = strutils.StringTrim(selectText)

	// 聚焦
	//mainWindowsFocus()

	SetFormatResult(selectText)

	homeWindow.RequestFocus()

	smartFormatFunc()
}

// -1 代表展示的是最后一个历史记录
var historyShowIndex = 0

// 记录展示的是第几个历史记录, 比如 (10 / 10) 代表展示的是第 10 个历史记录
//var historyIndexText string

func saveHistory(result string) {
	// 如果 historyShowIndex = -1 证明此时并没有展示历史记录, 可以直接将 text 保存到最后

	// 如果 historyShowIndex != -1 证明此时展示的是历史记录
	// 需要将 text 保存到 history 的最末尾, 并且把 historyShowIndex 也变回来 -1
	// 然后 historyIndexText 也要更新

	historyList = append(historyList, result)
	historyTextLabel.SetText(fmt.Sprintf("%d/%d", len(historyList), len(historyList)))
	historyShowIndex = len(historyList)
}

// changePreviousHistory
// 展示上一个历史记录
func changePreviousHistory() {
	if historyShowIndex <= 1 {
		return
	}

	historyShowIndex = historyShowIndex - 1

	historyTextLabel.SetText(fmt.Sprintf("%d/%d", historyShowIndex, len(historyList)))

	arrIndex := historyShowIndex - 1
	homeInputBox.SetText(historyList[arrIndex])
	setNoteMsg("", nil)
}

// changeNextHistory
// 展示下一个历史记录
func changeNextHistory() {
	if historyShowIndex >= len(historyList) {
		return
	}

	historyShowIndex = historyShowIndex + 1

	historyTextLabel.SetText(fmt.Sprintf("%d/%d", historyShowIndex, len(historyList)))

	arrIndex := historyShowIndex - 1
	homeInputBox.SetText(historyList[arrIndex])
	setNoteMsg("", nil)
}

func SetFormatResult(result string) {
	if strutils.StringTrim(result) == strutils.StringTrim(homeInputBox.Text) {
		return
	}

	homeInputBox.SetText(result)

	saveHistory(result)
}
