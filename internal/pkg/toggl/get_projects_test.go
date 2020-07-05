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
