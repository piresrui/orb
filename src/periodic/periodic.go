package periodic

import (
	"log"
	"time"
)

type PeriodicFunc func() error

func Periodic(action PeriodicFunc, period time.Duration) {
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
