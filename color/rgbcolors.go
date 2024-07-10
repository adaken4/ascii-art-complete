package color

import (
	"errors"
	"fmt"
	"strconv"
)

func RgbColorCode(rgbValue string) (string, error) {
	matches := rgb.FindStringSubmatch(rgbValue)
	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])
	if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
		return "", errors.New("invalid RGB color code")
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil
}
