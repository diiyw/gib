package log

import (
	"log"
	"os"
)

type Type func() string

func Stdout(v string) Type {
	return func() string {
		log.Println(v)
		return v
	}
}

func File(name, v string) Type {
	return func() string {
		f, err := os.OpenFile("logs/"+name+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Println(err)
			return v
		}
		defer f.Close()
		_, err = f.Write([]byte(v))
		if err != nil {
			log.Println(err)
		}
		return v
	}
}
