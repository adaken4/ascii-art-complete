package justify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	widthBytes, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	width, err := strconv.Atoi(strings.TrimSpace(string(widthBytes)))
	if err != nil {
		return 0, err
	}

	return width, nil
}
