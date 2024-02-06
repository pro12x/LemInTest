package main

import (
	"fmt"
	"os"
	"strings"
)

func Read(filename string) ([]string, error) {
	file, err := os.ReadFile("./file/" + filename)
	var result []string
	if err != nil {
		return nil, fmt.Errorf("ERROR: cannot open <%v> file: no such or directory", filename)
	}
	if len(strings.TrimSpace(string(file))) == 0 {
		return nil, fmt.Errorf("ERROR: <%v> file is empty", filename)
	}
	result = strings.Split(string(file), "\n")
	return result, nil
}
