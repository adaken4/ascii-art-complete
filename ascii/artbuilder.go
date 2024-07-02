package ascii

import "strings"

func ArtStringBuilder(s string, path string) (string, error) {
	asciiMap, err := RuneAsciiArtMapCreator(path)
	if err != nil {
		return "", err
	}
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, v := range s {
			artLines := strings.Split(asciiMap[v], "\n")
			result.WriteString(artLines[i])
		}
		result.WriteString("\n")
	}
	return result.String(), nil
}
