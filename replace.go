package main

import "strings"

func Replace(s string, that string, by string) string {
	if len(s) != 0 {
		var result string
		for _, v := range strings.TrimSpace(s) {
			if string(v) == that {
				result += by
				continue
			}
			result += string(v)
		}
		return result
	}
	return ""
}
