package color

import (
	"fmt"
	"strconv"
	"strings"
)

func HexColorCode(hexValue string) string {
	hexDigit := strings.TrimPrefix(hexValue, "#")
	var r, g, b int64

	if len(hexDigit) == 3 {
		// 3-character hex code, convert to RGB
		r, _ = strconv.ParseInt(hexDigit[0:1]+hexDigit[0:1], 16, 64)
		g, _ = strconv.ParseInt(hexDigit[1:2]+hexDigit[1:2], 16, 64)
		b, _ = strconv.ParseInt(hexDigit[2:3]+hexDigit[2:3], 16, 64)
	} else {
		// 6-character hex code, convert to RGB
		r, _ = strconv.ParseInt(hexDigit[0:2], 16, 64)
		g, _ = strconv.ParseInt(hexDigit[2:4], 16, 64)
		b, _ = strconv.ParseInt(hexDigit[4:6], 16, 64)
	}
	// Return ANSI escape code for RGB
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}
