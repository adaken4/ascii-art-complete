package color

func NamedColorCode(name string) (string, bool) {
	// Map of color names to ANSI escape codes
	colorCodes := map[string]string{
		"black":     "\033[38;5;16m",
		"red":       "\033[38;5;196m",
		"green":     "\033[38;5;22m",
		"blue":      "\033[38;5;21m",
		"yellow":    "\033[38;5;226m",
		"magenta":   "\u001b[35m",
		"cyan":      "\u001b[36m",
		"white":     "\033[38;5;231m",
		"orange":    "\u001b[38;5;208m",
		"purple":    "\033[38;5;55m",
		"teal":      "\033[38;5;23m",
		"silver":    "\033[38;5;145m",
		"gray":      "\033[38;5;240m",
		"brown":     "\033[38;5;94m",
		"pink":      "\u001b[38;5;207m",
		"olive":     "\u001b[38;5;58m",
		"navy":      "\u001b[38;5;18m",
		"turquoise": "\u001b[38;5;80m",
		"lime":      "\033[38;5;46m",
		"indigo":    "\u001b[38;5;54m",
		"lavender":  "\u001b[38;5;183m",
		"charteuse": "\033[33m\033[34m",
		"salmon":    "\033[38;5;209m",
		"peach":     "\033[33m\033[96m",
		"seafoam":   "\033[32m\033[96m",
		"fuchsia":   "\033[38;5;201m",
		"violet":    "\033[33m\033[95m",
		"aqua":      "\033[38;5;51m",
		"maroon":    "\033[38;5;52m",
	}

	// Retrieve the Ansi code and report
	colorCode, ok := colorCodes[name]
	return colorCode, ok
}
