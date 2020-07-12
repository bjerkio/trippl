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
