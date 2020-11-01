package text

import "regexp"

func IsPhone(phone string) bool {
	re, err := regexp.Compile(`^1\d{10}$`)
	_ = err
	return re.Match([]byte(phone))
}

func IsEmail(mail string) bool {
	re, _ := regexp.Compile(`^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,})$`)
	return re.Match([]byte(mail))
}
