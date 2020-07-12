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

package project

import (
	"fmt"
	"testing"

	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/pkg/server"
)

func TestGetProjects(t *testing.T) {
	cnf, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	cnf.DatabasePath = nil

	config, err := server.PrepareConfig(cnf)
	if err != nil {
		t.Fatal(err)
	}

	projects, txxProjects, err := GetProjects(*config)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("local projects", projects)
	fmt.Println("tripletex projects", txxProjects)

	t.Fatal("Hello")

	if &projects == nil {
		t.Fatal("Not able to get projects from local storage")
	}

	if &txxProjects == nil {
		t.Fatal("Not able to get projects from Tripletex")
	}
}
