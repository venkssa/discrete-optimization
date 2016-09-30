package common

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type LineNum uint32

func Parse(rc io.ReadCloser, fn func(LineNum, uint32, uint32)) error {
	defer rc.Close()

	scanner := bufio.NewScanner(rc)
	var lineNum LineNum

	for scanner.Scan() {
		lineNum++
		d1, d2, err := splitAndConvertToInt(scanner.Text())

		if err != nil {
			return err
		}

		fn(lineNum, d1, d2)
	}

	return scanner.Err()
}

func splitAndConvertToInt(s string) (uint32, uint32, error) {
	splitStrings := strings.Split(strings.TrimSpace(s), " ")
	if len(splitStrings) != 2 {
		return 0, 0, fmt.Errorf("Line should contain 2 numbers; but was instead %s", splitStrings)
	}
	first, err := strconv.ParseUint(splitStrings[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	second, err := strconv.ParseUint(splitStrings[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return uint32(first), uint32(second), nil
}
