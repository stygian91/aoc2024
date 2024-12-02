package q01

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/demo.txt
var dataDemo string

//go:embed inputs/main.txt
var dataMain string

func splitLine(line string) (int, int, error) {
	firstSpace := strings.Index(line, " ")
	lastSpace := strings.LastIndex(line, " ")
	if firstSpace == -1 || lastSpace == -1 {
		return 0, 0, fmt.Errorf("Couldn't find a space: '%s'", line)
	}

	l, err := strconv.ParseInt(line[0:firstSpace], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Error while parsing left int: %s", line[0:firstSpace])
	}

	r, err := strconv.ParseInt(line[lastSpace+1:], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Error while parsing right int: %s", line[lastSpace+1:])
	}

	return int(l), int(r), nil
}
