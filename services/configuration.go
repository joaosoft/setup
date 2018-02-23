package gomock

// App ...
type App struct {
	Config Configurations `json:"configurations"`
}

// Configurations ...
type Configurations struct {
	NSQ   ConfigNSQ   `json:"nsq"`
	SQL   ConfigSQL   `json:"sql"`
	Redis ConfigRedis `json:"redis"`
}

// ConfigNSQ ...
type ConfigNSQ struct {
	Lookupd      string `json:"lookupd"`
	RequeueDelay int64  `json:"requeue_delay"`
	MaxInFlight  int    `json:"max_in_flight"`
	MaxAttempts  uint16 `json:"max_attempts"`
}

// ConfigSQL ...
type ConfigSQL struct {
	Driver     string `json:"driver"`
	DataSource string `json:"datasource"`
}

// ConfigRedis ...
type ConfigRedis struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Size     int    `json:"size"`
}
