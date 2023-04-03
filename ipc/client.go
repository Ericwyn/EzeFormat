package ipc

import (
	"github.com/Ericwyn/EzeFormat/log"
)

func SendMessage(message IpcMessage) error {
	us := NewUnixSocket(UnixSocketAddress)
	res, err := us.ClientSendContext(message)
	if err != nil {
		return err
	}
	log.D("ipc response:", res)
	return nil
}
