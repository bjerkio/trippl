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
	"encoding/json"
	"fmt"

	"github.com/dougEfresh/gtoggl"
	"github.com/dougEfresh/gtoggl-api/gtproject"
)

func GetProjects(client *gtoggl.TogglClient, workspaceID uint64) (*gtproject.Projects, error) {
	body, err := client.TogglHttpClient.GetRequest(fmt.Sprintf("https://www.toggl.com/api/v8/workspaces/%d/projects", workspaceID))
	if err != nil {
		return nil, err
	}

	var projects gtproject.Projects

	err = json.Unmarshal(*body, &projects)
	if err != nil {
		return nil, err
	}
	return &projects, nil
}
