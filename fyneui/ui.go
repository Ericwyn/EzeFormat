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
	"github.com/Ericwyn/EzeFormat/utils/format"
	"github.com/Ericwyn/EzeFormat/utils/strutils"
	"github.com/Ericwyn/EzeFormat/utils/xclip"
	"image/color"
	"time"
)

var version = "V1.0.4"

var mainApp fyne.App

var homeNoteLabel *widget.Label

var homeInputBox *EzeMultiLineEntry

func StartApp(useXclipData bool) {
	if useXclipData {
		go func() {
			// sleep 0.5s
			time.Sleep(time.Millisecond * 300)
			// 获取滑词然后格式化翻译
			getSelectTextAndSmartFormat()
		}()
	}

	ShowMainUi()

}

func ShowMainUi() {
	mainApp = app.New()
	mainApp.SetIcon(resource.ResourceIcon())
	mainApp.Settings().SetTheme(&resource.CustomerTheme{})

	homeWindow := mainApp.NewWindow("EzeFormat")

	homeWindow.Resize(fyne.Size{
		Width: 700,
		//Height: 600,
	})
	homeWindow.CenterOnScreen()

	// --------------------
	homeInputBox = NewEzeMultiLineEntry(fyne.Size{Height: 550})
	//homeInputBox.Resize()

	homeNoteLabel = widget.NewLabel("")

	authorNote := widget.NewLabel(" Source: https://github.com/Ericwyn/EzeFormat 【 " + version + "】")

	//hello := widget.NewLabel("Hello Fyne!")
	homeWindow.SetContent(container.NewVBox(
		homeInputBox,
		homeNoteLabel,
		BtnPanelLine1(),
		BtnPanelLine2(),
		authorNote,
	))

	homeWindow.ShowAndRun()
}

// BtnPanelLine1
// 智能解析, 智能压缩,当前时间, 换行
func BtnPanelLine1() *fyne.Container {
	return container.NewHBox(
		newEzeButton("智能解析", func() {
			smartFormatFunc()
		}),

		newEzeButton("智能压缩", func() {
			smartCompressFunc()
		}),

		newEzeButton("当前时间", func() {
			timeNowFunc()
		}),

		newEzeButton("换行切换", func() {
			homeInputBox.ToggleWrap()
		}),
	)

}

var closeColor *canvas.Rectangle
var openColor *canvas.Rectangle

func GetCloseColor() *canvas.Rectangle {
	if closeColor == nil {
		closeColor = canvas.NewRectangle(color.NRGBA{R: 242, G: 242, B: 242, A: 255})
	}
	return closeColor
}

func GetOpenColor() *canvas.Rectangle {
	if openColor == nil {
		openColor = canvas.NewRectangle(color.NRGBA{R: 204, G: 255, B: 204, A: 255})
	}
	return openColor
}

func BtnPanelLine2() *fyne.Container {

	btn := widget.NewButton("测试", func() {
		// hello.SetText("Welcome :)")
	})

	btn.Resize(fyne.NewSize(260, 160))
	btn.Move(fyne.NewPos(30-theme.Padding(), 10))

	// closeColor = canvas.NewRectangle(color.NRGBA{R: 242, G: 242, B: 242, A: 255})
	// // closeColor.MinSize()

	// openClolor = canvas.NewRectangle(color.NRGBA{R: 204, G: 255, B: 204, A: 255})
	// // closeColor.MinSize()

	//maxLayout := container.New(layout.NewMaxLayout(), GetCloseColor(), GetOpenColor(), btn)

	return container.NewHBox(
		//maxLayout,
		newEzeButton("JSON 解析 ", func() {
			jsonFormatFunc()
		}),

		newEzeButton("JSON 压缩 ", func() {
			jsonCompressFunc()
		}),

		newEzeButton("XML 解析  ", func() {
			xmlFormatFunc()
		}),

		newEzeButton("XML 压缩  ", func() {
			xmlCompressFunc()
		}),
	)
}

// 自定义按钮，可以设置最小尺寸
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
		Width:  100,
		Height: 35,
	}
	return btn
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

	homeInputBox.SetText(result)
}

// smartCompressFunc 智能格式化
func smartCompressFunc() {
	result, err := format.SmartCompress(homeInputBox.Text)
	setNoteMsg("智能压缩失败: ", err)

	homeInputBox.SetText(result)
}

// timeNowFunc 展示当前时间戳
func timeNowFunc() {
	result, err := format.FormatType(fmt.Sprint(time.Now().UnixMilli()), format.TypeTimeStampMills)
	setNoteMsg("展示当前时间戳失败: ", err)

	homeInputBox.SetText(result)
}

// jsonFormatFunc 格式化
func jsonFormatFunc() {
	result, err := format.FormatType(homeInputBox.Text, format.TypeJson)
	setNoteMsg("JSON 格式化失败: ", err)

	homeInputBox.SetText(result)
}

// jsonCompressFunc 压缩
func jsonCompressFunc() {
	result, err := format.CompressType(homeInputBox.Text, format.TypeJson)
	setNoteMsg("JSON 压缩失败: ", err)

	homeInputBox.SetText(result)
}

// xmlFormatFunc 格式化
func xmlFormatFunc() {
	result, err := format.FormatType(homeInputBox.Text, format.TypeXml)
	setNoteMsg("Xml 格式化失败: ", err)

	homeInputBox.SetText(result)
}

// xmlCompressFunc 压缩
func xmlCompressFunc() {
	result, err := format.CompressType(homeInputBox.Text, format.TypeXml)
	setNoteMsg("Xml 压缩失败: ", err)

	homeInputBox.SetText(result)
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

	homeInputBox.SetText(selectText)

	smartFormatFunc()
}
