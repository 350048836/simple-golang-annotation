package fileutils

import (
	"io/ioutil"
	"strings"
)

func ReadFileAsLines(filepath string) ([]string, error) {
	var result []string
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}
