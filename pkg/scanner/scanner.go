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
	if !s.scn.Scan() {
		return "", false
	}
	return s.scn.Text(), true
}
