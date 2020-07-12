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

package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	_, b, _, _      = runtime.Caller(0)
	basePath        = filepath.Dir(b)
	projectBasePath = filepath.Join(basePath, "..", "..", "..", ".env")
)

// GetConfig returns config from environment variables
func GetConfig() (*TripplConfig, error) {
	err := godotenv.Load(projectBasePath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	var c TripplConfig
	err = envconfig.Process("trippl", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
