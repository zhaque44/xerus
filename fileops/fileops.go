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
