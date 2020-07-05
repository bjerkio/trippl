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
