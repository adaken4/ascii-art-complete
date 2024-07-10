package color

import (
	"errors"
	"regexp"
)

// Regular expressions for matching color formats
var (
	hex = regexp.MustCompile(`^#([a-f0-9]{3}|[a-f0-9]{6})$`)
	rgb = regexp.MustCompile(`^rgb\(\s*(\d+),\s*(\d+),\s*(\d+)\)$`)
)

// SetColor function converts color names, hexadecimal codes and RGB values to ANSI escape codes for terminal coloring
func SetColor(colorNameOrValue string) (string, error) {
	// Check if the given color name exists in the map
	code, ok := NamedColorCode(colorNameOrValue)

	// If not found, check for RGB or HEX values
	if !ok {
		if colorNameOrValue == "" {
			return "", nil

			// If Hex format, extract Hex values and return ANSI escape code
		} else if hex.MatchString(colorNameOrValue) {
			return HexColorCode(colorNameOrValue), nil

			// If RGB format, extract RGB values and return ANSI escape code
		} else if rgb.MatchString(colorNameOrValue) {
			RGBAnsiString, err := RgbColorCode(colorNameOrValue)
			if err != nil {
				return "", err
			}
			return RGBAnsiString, nil
		} else {
			return "", errors.New("invalid color or color code")
		}
	}
	// Return the ANSI code for provided named color
	return code, nil
}
