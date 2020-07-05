package toggl

import (
	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/dougEfresh/gtoggl"
)

// GetClient
func GetClient(config config.TripplConfig) (*gtoggl.TogglClient, error) {
	return gtoggl.NewClient(config.TogglAPIToken)
}
