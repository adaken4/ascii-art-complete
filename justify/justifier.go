package justify

import (
	"fmt"
	"strings"
)

func ArtAligner(position, artText string) (string, error) {
	if position == "left" {
		return artText, nil
	}

	var result strings.Builder

	terminalWidth, err := getTerminalWidth()
	if err != nil {
		return "", err
	}
	// Split artText into lines
	lines := strings.Split(artText, "\n")
	lines = lines[:len(lines)-1]

	allSpaceIndexes := [][]int{}
	for i := 0; i < 8; i++ {
		if strings.TrimSpace(lines[i]) == "" {
			continue
		}
		allSpaceIndexes = append(allSpaceIndexes, SpaceIndexes(lines[i], "        "))
	}
	spaceIndexes := findFurthestIndexes(allSpaceIndexes)
	fmt.Println(spaceIndexes)
	fmt.Println(allSpaceIndexes)

	// Define the padding
	var paddedLines []string

	// todo
	// calculate num spACES

	// Pad each line individually
	for _, line := range lines {
		textWidth := len(line)
		padding := 0
		switch position {
		case "center":
			padding = (terminalWidth - textWidth) / 2
			paddedLine := strings.Repeat(" ", padding) + line
			paddedLines = append(paddedLines, paddedLine)
		case "right":
			padding = (terminalWidth - textWidth)
			paddedLine := strings.Repeat(" ", padding) + line
			paddedLines = append(paddedLines, paddedLine)
		case "justify":
			// spaceIndexes := SpaceIndexes(line, "      ") // Assuming 6 spaces for ASCII art
			numSpaces := len(spaceIndexes)
			// fmt.Println(numSpaces)

			if numSpaces > 0 {
				totalSpaces := terminalWidth - textWidth
				spacesPerGap := totalSpaces / numSpaces
				extraSpaces := totalSpaces % numSpaces

				var paddedLine strings.Builder
				lastIndex := 0

				for i, index := range spaceIndexes {
					paddedLine.WriteString(line[lastIndex:index])
					numSpacesToAdd := spacesPerGap
					if i < extraSpaces {
						numSpacesToAdd++
					}
					paddedLine.WriteString(strings.Repeat(" ", numSpacesToAdd))
					lastIndex = index
				}

				// Append the remaining part of the line
				paddedLine.WriteString(line[lastIndex:])

				paddedLines = append(paddedLines, paddedLine.String())
			} else {
				// If no spaces found, left-align the line
				paddedLines = append(paddedLines, line)
			}

		}
	}

	// Write each padded line and ensure the prompt is not padded
	for i, paddedLine := range paddedLines {
		if i == len(paddedLines)-1 {
			result.WriteString("")
		} else {
			result.WriteString(paddedLine + "\n")
		}
	}
	return result.String(), nil
}

// findFurthestIndexes finds the furthest indexes in the provided 2D slice
func findFurthestIndexes(indexes [][]int) []int {
	if len(indexes) == 0 {
		return nil
	}

	// Step 1: Determine the smallest length
	smallestLength := len(indexes[0])
	for _, slice := range indexes {
		length := len(slice)
		if length < smallestLength {
			smallestLength = length
		}
	}

	// Step 2: Filter slices by the smallest length
	filteredSlices := [][]int{}
	for _, slice := range indexes {
		if len(slice) == smallestLength {
			filteredSlices = append(filteredSlices, slice)
		}
	}

	// Step 3: Find the furthest indexes from the filtered slices
	if len(filteredSlices) == 0 {
		return nil
	}

	maxIndexes := make([]int, len(filteredSlices[0]))
	for _, indexList := range filteredSlices {
		for i, index := range indexList {
			if index > maxIndexes[i] {
				maxIndexes[i] = index
			}
		}
	}
	return maxIndexes
}

func SpaceIndexes(line, spaces string) []int {
	var indexes []int
	for i := 0; i < len(line); {
		index := strings.Index(line[i:], spaces)
		if index == -1 {
			break
		}
		indexes = append(indexes, i+index)
		i += index + 6
	}
	return indexes
}
