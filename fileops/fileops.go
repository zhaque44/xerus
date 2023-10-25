package fileops

import (
	"bufio"
	"errors"
	"net"
	"os"
	"path/filepath"
	"xerus/datapreparation"
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
	s, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	m := s.Mode()

	isFIFO := (m & os.ModeNamedPipe) != 0
	isCharDevice := (m & os.ModeCharDevice) == 0

	return isCharDevice || isFIFO, nil
}

func FindFilesByPattern(pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, errors.New("no files found matching the pattern")
	}

	return matches, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func CollectCIDRs(inputData string, separator string, recursionLimit int, fileDelimiter string) ([]string, error) {
	if recursionLimit < 0 {
		return nil, errors.New("recursionLimit must be >= 0")
	}

	items, err := datapreparation.SplitAndTrimWhitespace(inputData, separator)
	if err != nil {
		return nil, err
	}

	var parsedCIDRs []string

	for _, item := range items {
		if ip := net.ParseIP(item); ip != nil {
			parsedCIDRs = append(parsedCIDRs, item)
		} else if _, _, err := net.ParseCIDR(item); err == nil {
			parsedCIDRs = append(parsedCIDRs, item)
		} else if FileExists(item) {
			filedata, err := os.ReadFile(item)
			if err != nil {
				return nil, err
			}

			if len(filedata) > 0 && recursionLimit > 0 {
				recursedList, err := CollectCIDRs(string(filedata), separator, recursionLimit-1, fileDelimiter)
				if err != nil {
					return nil, err
				}
				parsedCIDRs = append(parsedCIDRs, recursedList...)
			}
		}
	}

	return parsedCIDRs, nil
}
