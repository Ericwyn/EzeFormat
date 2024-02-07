package gtkui

import (
	"fmt"
	"github.com/Ericwyn/EzeFormat/ipc"
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/EzeFormat/utils/xclip"
	"runtime"
	"time"
)

func trySendMessage(message ipc.IpcMessage) bool {
	if runtime.GOOS != "linux" {
		log.D("not linux, don't send socket msg")
		return false
	}
	err := ipc.SendMessage(message)
	if err == nil {
		log.D("已发送给其他翻译进程 : " + string(message))
		return true
	} else {
		log.D("IPC 消息发送失败")
		return false
	}
}

func sendSocketMessage(sleep bool) {

	if sleep {
		time.Sleep(time.Millisecond * 500)
	}

	//if xclip {
	//	trySendMessage(ipc.IpcMessageNewSelection)
	//} else if ocr {
	//	trySendMessage(ipc.IpcMessageOcr)
	//} else if jsonFormat {
	//	trySendMessage(ipc.IpcMessageJsonFormat)
	//}

	trySendMessage(ipc.IpcMessageJsonFormat)
}

func startUnixSocketServer() {
	if runtime.GOOS != "linux" {
		log.D("not linux, don't start socket server")
		return
	}
	go ipc.StartUnixSocketListener(func(message ipc.IpcMessage) {
		log.D("接收到 IPC 消息 : " + string(message))
		switch message {
		case ipc.IpcMessagePing:
			break
		case ipc.IpcMessageJsonFormat:
			setSelectTextToJsonFormatBox()
			break
		}
	})
}

func setSelectTextToJsonFormatBox() bool {
	//jsonWindow.RequestFocus()
	//inputBox := jsonEntryBox

	selectText := xclip.GetSelection()
	fmt.Println("获取的划词:", selectText)

	//inputBox.SetText(selectText)
	//
	//startJsonFormat()

	// 聚焦
	mainWindowsFocus()

	SetInputFunc(selectText)
	FormatSmartFunc()

	return true
}
