package strutils

import (
	"fmt"
	"testing"
)

func TestGuessType(t *testing.T) {
	//fmt.Println(GuessType("{'aa' : 'bb'}"))
	//fmt.Println(GuessType("<aa>bb</aa>"))
	//fmt.Println(GuessType("1234567890"))
	//fmt.Println(GuessType("1234567890123"))
	//fmt.Println(GuessType("2023-04-03"))
	//fmt.Println(GuessType("2023-04-03 00:00"))
	//fmt.Println(GuessType("2023-04-03 00:00:00"))
	//fmt.Println(GuessType("2023-04-03 00:00:00.000"))

	fmt.Println(GuessType("2023-04-03"))
}

func TestFormatSmart(t *testing.T) {
	fmt.Println(FormatSmart("2023-04-03 18:18"))
}
