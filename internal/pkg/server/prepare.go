package server

import (
	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
	"github.com/bjerkio/trippl/internal/pkg/toggl"
	"github.com/bjerkio/trippl/internal/pkg/tripletex"
)

// PrepareConfig returns ServerArguments for use with Server.Run
func PrepareConfig(cnf *config.TripplConfig) (*Arguments, error) {
	if cnf == nil {
		res, err := config.GetConfig()
		if err != nil {
			return nil, err
		}

		cnf = res
	}

	db, err := db.EmbeddedKeyValueStore(cnf.DatabasePath)
	if err != nil {
		return nil, err
	}

	tripletexClient, tripletexAuth, err := tripletex.GetClient(*cnf, db)
	if err != nil {
		return nil, err
	}

	togglClient, err := toggl.GetClient(*cnf)
	if err != nil {
		return nil, err
	}

	return &Arguments{
		Config:          *cnf,
		DB:              db,
		TripletexClient: *tripletexClient,
		TripletexAuth:   tripletexAuth,
		TogglClient:     *togglClient,
	}, nil
}
