package server

import (
	"time"

	"github.com/bjerkio/trippl/internal/pkg/server"
	"github.com/bjerkio/trippl/pkg/syncprojects"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Run(opts server.Arguments) error {

	interval := time.Duration(opts.Config.Interval) * time.Second
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			err := syncprojects.Run(opts)
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}
