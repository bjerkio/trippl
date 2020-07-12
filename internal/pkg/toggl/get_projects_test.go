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

package toggl

import (
	"testing"

	"github.com/bjerkio/trippl/internal/pkg/config"
)

func TestGetProjects(t *testing.T) {
	cnf, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	client, err := GetClient(*cnf)
	if err != nil {
		t.Fatal(err)
	}
	res, err := GetProjects(client, 3523086)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("Could not get projects")
	}
}
