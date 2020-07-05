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
