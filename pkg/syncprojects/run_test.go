package syncprojects

import (
	"testing"

	"github.com/bjerkio/trippl/pkg/server"
)

func TestSyncProject(t *testing.T) {
	config, err := server.PrepareConfig(nil)
	if err != nil {
		t.Fatal(err)
	}

	err = Run(*config)
	if err != nil {
		t.Fatal(err)
	}
}
