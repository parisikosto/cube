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

// Command prints a cyan-colored message, used for shell command labels.
func Command(text string) {
	fmt.Println(colorCyan + text + colorReset)
}

// SubCommand prints a yellow-colored message, used for step descriptions.
func SubCommand(text string) {
	fmt.Println(colorYellow + text + colorReset)
}

// Success prints a green-colored message for successful operations.
func Success(text string) {
	fmt.Println(colorGreen + text + colorReset)
}

// Error prints a red-colored message for errors.
func Error(text string) {
	fmt.Println(colorRed + text + colorReset)
}

// Warning prints an orange-colored message for warnings.
func Warning(text string) {
	fmt.Println(colorOrange + text + colorReset)
}

// Suggestion prints a purple-colored message for suggestions.
func Suggestion(text string) {
	fmt.Println(colorPurple + text + colorReset)
}

// Instruction prints a white-colored message for instructions.
func Instruction(text string) {
	fmt.Println(colorWhite + text + colorReset)
}

// Styled returns a styled string without printing it.
// Useful when embedding a colored value inside another message.
func Styled(color, text string) string {
	return color + text + colorReset
}
