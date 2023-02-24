package main

import (
	"context"
	orb2 "github.com/piresrui/orb/orb"
	"os"
	"os/signal"
)

func main() {
	orb := orb2.ProvideVirtualOrb()
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer cancel()

	orb.Run()

	select {
	case <-c:
		cancel()
	case <-ctx.Done():
	}
}
