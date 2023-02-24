package periodic

import (
	"context"
	"log"
	"time"
)

type PeriodicFunc func() error

func Periodic(ctx context.Context, action PeriodicFunc, period time.Duration) {
	t := time.NewTicker(period * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			err := action()
			if err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			log.Println("terminating periodic func...")
			return
		}
	}
}
