package text

import (
	"net/url"
	"regexp"
	"strings"
)

func TrimHtml(s string) string {
	// HTML to lower
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	s = re.ReplaceAllStringFunc(s, strings.ToLower)
	// remove <style>
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	s = re.ReplaceAllString(s, "")
	// remove <script>
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	s = re.ReplaceAllString(s, "")
	// remove tags
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	s = re.ReplaceAllString(s, "\n")
	// remove space
	re, _ = regexp.Compile("\\s{2,}")
	return re.ReplaceAllString(s, "\n")
}

func Split(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return string(r)
	}
	return string(r[:n])
}

// UrlEncode encode url
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// UrlDecode decode url
func UrlDecode(s string) string {
	s, _ = url.QueryUnescape(s)
	return s
}
