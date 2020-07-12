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

package tripletex

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/bjerkio/tripletex-go/client"

	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
)

// GetClient returns a Tripletex Client with Auth
func GetClient(config config.TripplConfig, db db.KeyValueStore) (*apiclient.Tripletex, runtime.ClientAuthInfoWriter, error) {
	r, err := GetTransport(config, db)
	if err != nil {
		return nil, nil, err
	}
	return apiclient.New(r, strfmt.Default), r.DefaultAuthentication, nil
}
