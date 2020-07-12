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

package config

// TripplConfig configuration variables for Trippl
type TripplConfig struct {
	EmployeeToken    string  `required:"true" split_words:"true"`
	ConsumerToken    string  `required:"true" split_words:"true"`
	TogglAPIToken    string  `required:"true" split_words:"true"`
	TogglWorkspaceID uint64  `required:"true" split_words:"true"`
	DatabasePath     *string `split_words:"true"`
	Interval         int32   `default:"60"` // Interval in seconds
}
