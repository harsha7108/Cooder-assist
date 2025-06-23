package scanner

import (
	"bufio"
	"os"
)

type Scanner struct {
	scn *bufio.Scanner
}

func New() Scanner {
	return Scanner{scn: bufio.NewScanner(os.Stdin)}
}

func (s Scanner) GetUserMessage() (string, bool) {
	var message string
	for s.scn.Scan() {
		line := s.scn.Text()
		if line == "" {
			return message, true
		}
		message += line + "\n"
	}
	return message, false
}
