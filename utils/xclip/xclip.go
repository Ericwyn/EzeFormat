package xclip

import "github.com/Ericwyn/GoTools/shell"

// GetSelection 使用 xclip 获取选择的文字
func GetSelection() string {
	return shell.RunShellRes("xclip", "-out")
}
