package strings

import (
	"net/url"
	"regexp"
	"strings"
)

type Type func(s string) string

func TrimSpace() Type {
	return func(s string) string {
		return strings.TrimSpace(s)
	}
}

func TrimHtml() Type {
	return func(s string) string {
		// HTML标签全转换成小写
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		s = re.ReplaceAllStringFunc(s, strings.ToLower)

		// 去除STYLE
		re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
		s = re.ReplaceAllString(s, "")

		// 去除SCRIPT
		re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
		s = re.ReplaceAllString(s, "")

		// 去除所有尖括号内的HTML代码，并换成换行符
		re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		s = re.ReplaceAllString(s, "\n")

		// 去除连续的换行符
		re, _ = regexp.Compile("\\s{2,}")
		return re.ReplaceAllString(s, "\n")
	}
}

func Cut(n int) Type {

	return func(s string) string {
		r := []rune(s)
		if len(r) <= n {
			return string(r)
		}
		return string(r[:n])
	}
}

// Url编码
func UrlEncode() Type {
	return func(s string) string {
		return url.QueryEscape(s)
	}
}

// Url解码
func UrlDecode() Type {
	return func(s string) string {
		s, _ = url.QueryUnescape(s)
		return s
	}
}

func ToLower() Type {
	return func(s string) string {
		return strings.ToLower(s)
	}
}
