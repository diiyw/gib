package gos

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulClose(f func()) {
	go func() {
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-s
			f()
			os.Exit(0)
		}()
	}()
}
