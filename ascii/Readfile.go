package ascii

import (
	"bufio"
	"errors"
	"os"
)

// Function to read the ASCII Style files line by line
func ReadStyleFile(style string) ([]string, error) {
	if style == "standard" || style == "shadow" || style == "thinkertoy" {
		readFile, err := os.Open("banners/" + style + ".txt")

		if err != nil {
			return nil, err
		}

		// Scan the lines that was Read into slice of strings
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		styleSlice := []string{}
		for fileScanner.Scan() {
			styleSlice = append(styleSlice, fileScanner.Text())
		}

		readFile.Close()
		return styleSlice, nil
	} else {
		return nil, errors.New("bad Request, banner type is not valid")
	}
}
