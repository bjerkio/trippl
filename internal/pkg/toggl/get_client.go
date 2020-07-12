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
	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/dougEfresh/gtoggl"
)

// GetClient
func GetClient(config config.TripplConfig) (*gtoggl.TogglClient, error) {
	return gtoggl.NewClient(config.TogglAPIToken)
}
