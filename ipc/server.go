package ipc

import "github.com/Ericwyn/EzeFormat/log"

type IpcMessage string

const IpcMessageSmartFormat IpcMessage = "SMART_FORMAT\n"
const IpcMessagePing IpcMessage = "PING\n"

type MessageHandler func(message IpcMessage)

const UnixSocketAddress = "/tmp/eze_json.socket"

var PONG = "PONG\n"

func StartUnixSocketListener(messageHandler MessageHandler) {
	log.D("开始监听 IPC 消息")
	us := NewUnixSocket(UnixSocketAddress)
	us.SetContextHandler(func(message IpcMessage) string {
		messageHandler(message)
		return PONG
	})

	us.StartServer()
}
