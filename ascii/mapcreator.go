package ascii

import (
	"bufio"
	"os"
	"strings"
)

func RuneAsciiArtMapCreator(bannerPath string) (map[rune]string, error) {
	banner, err := os.Open(bannerPath)
	if err != nil {
		return nil, err
	}
	defer banner.Close()
	scanner := bufio.NewScanner(banner)
	var art string
	printableRune := ' '
	runeArtMap := make(map[rune]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if art != "" {
				runeArtMap[printableRune] = art
				art = ""
				printableRune++
			}
		} else {
			if strings.ContainsRune(line, '\r') {
				line = strings.TrimSuffix(line, "\r") + "\n"
				art += line
			} else {
				art += line + "\n"
			}
		}
	}
	if art != "" {
		runeArtMap[printableRune] = art
	}
	return runeArtMap, nil
}
