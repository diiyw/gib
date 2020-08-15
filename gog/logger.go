package gog

import (
	"encoding/json"
	"github.com/diiyw/gib/gos"
	"log"
	"os"
	"strconv"
	"time"
)

type Logger interface {
	Log() error
}

type CommonLogger struct {
	Message interface{} `json:"message"`
}

func (logger *CommonLogger) Log() error {
	log.Println(logger.Message)
	return nil
}

func Stdout(v interface{}, options ...Option) {
	logger := new(CommonLogger)
	logger.Message = v
	for _, op := range options {
		_ = op(logger)
	}
	_ = logger.Log()
}

func Fatal(v interface{}, options ...Option) {
	logger := new(CommonLogger)
	logger.Message = v
	for _, op := range options {
		_ = op(logger)
	}
	_ = logger.Log()
	os.Exit(1)
}

type FileLogger struct {
	CommonLogger

	Size     int64  `json:"-"`
	Dir      string `json:"-"`
	Filename string `json:"-"`
	Datetime string `json:"datetime"`
	Time     int64  `json:"time"`
	withRaw  bool   `json:"-"`
}

func (logger *FileLogger) Log() (err error) {
	f, err := os.OpenFile(logger.Dir+"/"+logger.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	var b []byte
	if !logger.withRaw {
		b, err = json.Marshal(logger)
		if err != nil {
			Stdout(err)
		}
	} else {
		if m, ok := logger.Message.(string); ok {
			b = []byte(m)
		}
	}
	_, err = f.Write(b)
	_, err = f.Write([]byte("\r\n"))
	return err
}

func File(v interface{}, options ...Option) {
	logger := new(FileLogger)
	logger.Size = 1024 * 1024 * 100
	logger.Dir = "logs"
	logger.Filename = gos.Date()
	logger.Datetime = gos.DateTime()
	logger.Message = v
	logger.Time = time.Now().UnixNano()
	for _, op := range options {
		if err := op(logger); err != nil {
			Stdout(err)
		}
	}
	var i int
	for {
		logger.Filename += "_" + strconv.Itoa(i) + ".log"
		fi, err := os.Stat(logger.Dir + "/" + logger.Filename)
		if err != nil {
			break
		}
		if fi.Size() < logger.Size {
			break
		}
		i++
	}
	_ = logger.Log()
}
