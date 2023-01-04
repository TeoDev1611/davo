package core

import "strings"

func IndexOf(element string, data []interface{}) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
