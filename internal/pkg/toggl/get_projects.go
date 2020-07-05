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
