package strings

import "regexp"

func MatchPhone(phone string) bool {
	re, _ := regexp.Compile(`1\d{10}`)
	return re.Match([]byte(phone))
}

func MatchEmail(mail string) bool {
	re, _ := regexp.Compile(`^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,})$`)
	return re.Match([]byte(mail))
}
