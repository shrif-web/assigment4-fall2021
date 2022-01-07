package go_server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func NewKillNotifierContext() context.Context {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	ctx, canFu := context.WithCancel(context.Background())
	go func() {
		<-c
		canFu()
	}()
	return ctx
}
