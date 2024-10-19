package reverse



// ExtractAsciiArt extracts a specific section of ASCII art from the 2D board starting
// from a given startIndex and spans across a specified width.
// Returns the concatenated string representing the extracted ASCII art.
func ExtractAsciiArtChar(board [][]string, startIndex, width int) string {
	// Get the number of rows and columns in the board
	rows := len(board)
	cols := len(board[0])
	result := ""
	// Iterate through each row of the board
	for i := 0; i < rows; i++ {
		// For each row, iterate through the characters in the specified width from the starting index of column
		for j := startIndex; j < startIndex+width; j++ {
			// Ensure the index doesn't exceed the number of columns
			if j < cols {
				// Append the character to the result string
				result += board[i][j]
			}
		}
	}
	// Return the concatenated result as the extracted ASCII art to be checked in the Universal Map
	return result
}
