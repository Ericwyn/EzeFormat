package xclip

import (
	"bytes"
	"github.com/Ericwyn/GoTools/shell"
	"os/exec"
)

// GetSelection 使用 xclip 获取选择的文字
func GetSelection() string {
	return shell.RunShellRes("xclip", "-out")
}

func SetClipboard(text string) {
	cmd := exec.Command("sh", "-c", "xclip -selection clipboard")
	cmd.Stdin = bytes.NewBufferString(text)

	cmd.Run()
}
