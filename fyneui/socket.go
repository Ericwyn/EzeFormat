package fyneui

import (
	"github.com/Ericwyn/EzeFormat/ipc"
	"github.com/Ericwyn/EzeFormat/log"
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

	trySendMessage(ipc.IpcMessageSmartFormat)
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
		case ipc.IpcMessageSmartFormat:
			getSelectTextAndSmartFormat()
			break
		}
	})
}
