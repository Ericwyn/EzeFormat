package gotkutils

import (
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/gotk3/gotk3/gtk"
)

func NewBox(orient gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.E("Unable to create box:", err)
		return nil
	}
	return box
}

//func SetMargin(widget interface{}, marginTop int, marginBottom int, marginStart int, marginEnd int) {
//	refutils.ReflectCall(widget, "SetMarginTop", []any{marginTop})
//	refutils.ReflectCall(widget, "SetMarginBottom", []any{marginBottom})
//	refutils.ReflectCall(widget, "SetMarginStart", []any{marginStart})
//	refutils.ReflectCall(widget, "SetMarginEnd", []any{marginEnd})
//}

type BtnDefine struct {
	Name    string
	Label   string
	OnClick func()
	Width   int
	Height  int
}

func NewBtn(btnDefine BtnDefine) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(btnDefine.Label)
	if err != nil {
		log.E("Unable to create button:", err)
		return nil
	}
	btn.SetName(btnDefine.Name)
	btn.Connect("clicked", btnDefine.OnClick)
	btn.SetSizeRequest(btnDefine.Width, btnDefine.Height)
	return btn
}

type CheckBtnDefine struct {
	Name    string
	Label   string
	OnClick func(bool)
	Width   int
	Height  int
}

func NewCheckBtn(btnDefine CheckBtnDefine) *gtk.CheckButton {
	btn, err := gtk.CheckButtonNewWithLabel(btnDefine.Label)
	if err != nil {
		log.E("Unable to create button:", err)
		return nil
	}
	log.I("create check btn")
	btn.SetSizeRequest(btnDefine.Width, btnDefine.Height)
	btn.Connect("toggled", func() {
		btnDefine.OnClick(btn.GetActive())
	})
	btn.SetName(btnDefine.Name)
	return btn
}

//var textViewBuffCache = make(map[*gtk.TextView]*gtk.TextBuffer)

func SetTextViewInput(tv *gtk.TextView, text string) {
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.E("Unable to get buffer:", err)
		return
	}
	buffer.SetText(text)
}

func GetTextViewInput(tv *gtk.TextView) string {
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.E("Unable to get buffer:", err)
		return ""
	}
	start, end := buffer.GetBounds()

	text, err := buffer.GetText(start, end, true)
	if err != nil {
		log.E("Unable to get text:", err)
		return ""
	}
	return text
}
