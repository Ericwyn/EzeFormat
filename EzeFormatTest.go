package main

import (
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/EzeFormat/utils/pathutils"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	//gtkui.OpenNewApp()

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.E("Unable to create window:", err)
	}
	win.SetTitle("hahahahah")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetSizeRequest(800, 600)
	win.SetResizable(false)
	win.SetPosition(gtk.WIN_POS_CENTER)

	//iconPath := pathutils.GetRunnerPath() + "/res-static/icon/json.ico"
	iconPath := "./res-static/icon/json.ico"
	log.I("runPath: ", pathutils.GetRunnerPath(), ", icon: "+iconPath)

	err = win.SetIconFromFile("/home/ericwyn/dev/go/EzeFormat/res-static/json.ico")

	//iconImg, err := gdk.PixbufNewFromFile(iconPath)
	if err != nil {
		log.E("set icon error ", err)
	}
	//win.SetIcon(iconImg)
	win.SetIconName("EzeFormat")
	win.SetName("EzeFormat")

	win.ShowAll()

	gtk.StatusbarNew()

	gtk.Main()

}
