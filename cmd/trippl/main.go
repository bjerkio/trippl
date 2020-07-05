package main

import (
	iserver "github.com/bjerkio/trippl/internal/pkg/server"
	"github.com/bjerkio/trippl/pkg/server"
)

func main() {
	config, err := iserver.PrepareConfig(nil)
	if err != nil {
		panic(err)
	}

	defer config.DB.Close()

	err = server.Run(*config)
	if err != nil {
		panic(err)
	}
}
