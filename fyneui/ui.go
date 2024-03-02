package fyneui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Ericwyn/EzeFormat/fyneui/resource"
	"image/color"
)

var version = "V1.0.4"

var mainApp fyne.App

var homeNoteLabel *widget.Label

var homeInputBox *EzeMultiLineEntry

func StartApp(useXclipData bool) {
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

		}),

		newEzeButton("智能压缩", func() {

		}),

		newEzeButton("当前时间", func() {

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

		}),

		newEzeButton("JSON 压缩 ", func() {

		}),

		newEzeButton("XML 解析  ", func() {

		}),

		newEzeButton("XML 压缩  ", func() {

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
