package ui

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorOrange = "\033[38;5;208m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func Command(text string) {
	fmt.Println(colorCyan + text + colorReset)
}

func SubCommand(text string) {
	fmt.Println(colorYellow + text + colorReset)
}

func Success(text string) {
	fmt.Println(colorGreen + text + colorReset)
}

func Error(text string) {
	fmt.Println(colorRed + text + colorReset)
}

func Warning(text string) {
	fmt.Println(colorOrange + text + colorReset)
}

func Suggestion(text string) {
	fmt.Println(colorPurple + text + colorReset)
}

func Instruction(text string) {
	fmt.Println(colorWhite + text + colorReset)
}

// Styled returns a styled string without printing it.
// Useful when you need to embed a styled value inside another message.
func Styled(color, text string) string {
	return color + text + colorReset
}
