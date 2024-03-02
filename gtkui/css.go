package gtkui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const cssData = `
#appContainer {
	padding: 10px;
}

#inputScrollWin {
	background-color: white;
	padding-top: 10px;
	padding-left: 10px;
	padding-right: 10px;
	padding-bottom: 0px;
	margin-bottom: 10px;
}

#inputTextView {

}

textview text {
  background-color: white;
}

#SmartFormatBtn, #CompressFormatBtn, #TimeNowBtn,
#FormatJsonBtn, #CompressJsonBtn, 
#FormatXmlBtn, #CompressXmlBtn,
#cleanTextViewBtn {
	margin-right: 10px;
	margin-top: 5px;
}

#WrapChangeBtn {
	margin-left: 20px;
}

#authorLine {
	margin-top: 10px;
	margin-bottom: 5px;
	color: #999;
	font-size: 14px;
}

`

// css implementation
func runCss() error {
	cssProv, err := gtk.CssProviderNew()
	if err == nil {
		if err = cssProv.LoadFromData(cssData); err == nil {
			screen, err := gdk.ScreenGetDefault()
			if err != nil {
				return err
			}
			gtk.AddProviderForScreen(screen, cssProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
		}
	}
	return err
}
