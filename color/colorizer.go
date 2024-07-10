package color

// Colorize function applies color to the given message using ANSI escape codes
func Colorize(color, message string) string {
	return color + message + "\u001b[0m" // Apply and reset color
}
