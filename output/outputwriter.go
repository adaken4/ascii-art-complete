package output

import (
	"errors"
	"os"
	"strings"
)

// OutputWriter writes the art representation to a .txt file provided in the output flag
func OutputWriter(outputFile, artText string) error {
	if !strings.HasSuffix(outputFile, ".txt") {
		return errors.New("invalid file extension: please provide a .txt file")
	} else if strings.HasPrefix(outputFile, "./banners/") || strings.HasPrefix(outputFile, "banners/") {
		return errors.New("writing to banners directory is forbidden to prevent modifying art banner files")
	}
	err := os.WriteFile(outputFile, []byte(artText), 0o644)
	if err != nil {
		return err
	}
	return nil
}
