package ascii

import (
	"errors"
	"strings"
)

func GenerateAscii(txt string, asciiStyle string) (string, error) {
	// Splitting the string based on the \n separator
	fileContent, err := ReadStyleFile(asciiStyle)
	if err != nil {
		return "", err
	}
	result := ""
	index := 0
	var lines []string
	if strings.Contains(txt, "\n") {
		lines = strings.Split(txt, "\n")
	} else {
		lines = strings.Split(txt, "\\n")
	}
	for _, line := range lines {
		if line != "" {
			for i := 0; i < 8; i++ {
				currentLine := ""
				for _, letter := range line {

					if letter > 31 && letter < 128 {
						currentLine += strings.TrimSuffix(fileContent[getIndexRune(letter)+i], "\n")
					} else if letter != '\n' && letter != '\r' {
						return "", errors.New("contains non ascii compatible characters")
					}

					index++
				}
				index = index - len(line)
				result += currentLine + "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result, nil
}

// Return the Index of the ascii character
func getIndexRune(r rune) int {
	return ((int(r) - 32) * 9) + 1
}
