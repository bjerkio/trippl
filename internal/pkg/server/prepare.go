// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
