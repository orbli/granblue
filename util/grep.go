package util

import (
	"bufio"
	"strings"
)

func Grep(scanner *bufio.Scanner, s string) string {
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return scanner.Text()
		}
	}
	return ""
}
