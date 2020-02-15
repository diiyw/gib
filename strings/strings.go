package strings

import "strings"

func Format(s string, types ...Type) string {
	for _, t := range types {
		s = t(s)
	}
	return s
}

// 字符串包含
func Has(s, sub string) bool {
	return strings.Contains(s, sub)
}
