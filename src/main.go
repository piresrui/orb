package main

import (
	"context"
	orb2 "github.com/piresrui/orb/orb"
	"log"
	"os"
	"os/signal"
	"time"
)

func periodic(action orb2.PeriodicFunc, period time.Duration) {
	t := time.NewTicker(period * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			err := action()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func main() {
	orb := orb2.ProvideVirtualOrb()
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer cancel()

	go periodic(orb.Status, 1)
	go periodic(orb.Signup, 1)

	select {
	case <-c:
		cancel()
	case <-ctx.Done():
	}
}
