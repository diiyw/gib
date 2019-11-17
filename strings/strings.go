package strings

import (
	"net/url"
	"regexp"
	"strings"
)

func TrimHtml(src string) string {

	// HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// Url编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// Url解码
func UrlDecode(s string) string {
	s, _ = url.QueryUnescape(s)
	return s
}

func ToLower(s string) string {
	// 小写
	return strings.ToLower(s)
}

// 去除html并截取固定长度
func TrimHtmlSplit(src string, n int) string {
	r := []rune(TrimHtml(src))
	if len(r) <= n {
		return string(r)
	}
	return string(r[:n])
}

// 字符串包含
func Has(s, sub string) bool {
	return strings.Contains(s, sub)
}
