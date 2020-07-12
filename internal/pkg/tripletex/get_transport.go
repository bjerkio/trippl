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

	apiclient "github.com/bjerkio/tripletex-go/client"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
)

func GetTransport(config config.TripplConfig, db db.KeyValueStore) (*httptransport.Runtime, error) {
	token, err := CreateToken(config.ConsumerToken, config.EmployeeToken, db)
	if err != nil {
		return nil, err
	}

	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth("0", *token)
	r.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	return r, nil
}
