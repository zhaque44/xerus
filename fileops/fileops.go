package fileops

import (
	"bufio"
	"os"
)

func LoadFile(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func IsPipedInput() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	mode := stat.Mode()

	isFIFO := (mode & os.ModeNamedPipe) != 0
	isCharDevice := (mode & os.ModeCharDevice) == 0

	return isCharDevice || isFIFO, nil
}
