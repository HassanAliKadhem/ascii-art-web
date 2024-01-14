package ascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ConsoleWidth() int {
	cmd := exec.Command("stty", "size")
	defer exec.Command("clear")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}
