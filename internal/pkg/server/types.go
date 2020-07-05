package server

import (
	txxclient "github.com/bjerkio/tripletex-go/client"
	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
	"github.com/dougEfresh/gtoggl"
	"github.com/go-openapi/runtime"
)

type Arguments struct {
	Config          config.TripplConfig
	DB              db.KeyValueStore
	TripletexClient txxclient.Tripletex
	TripletexAuth   runtime.ClientAuthInfoWriter
	TogglClient     gtoggl.TogglClient
}
