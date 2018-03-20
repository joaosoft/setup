package gomock

// App ...
type App struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

// Configurations ...
type Configurations struct {
	Connections Connections `json:"connections"`
}

// Connections ...
type Connections struct {
	NsqConfig   NsqConfig   `json:"nsq"`
	SqlConfig   SqlConfig   `json:"sql"`
	RedisConfig RedisConfig `json:"redis"`
}

// NsqConfig ...
type NsqConfig struct {
	Lookupd      string `json:"lookupd"`
	RequeueDelay int64  `json:"requeue_delay"`
	MaxInFlight  int    `json:"max_in_flight"`
	MaxAttempts  uint16 `json:"max_attempts"`
}

// SqlConfig ...
type SqlConfig struct {
	Driver     string `json:"driver"`
	DataSource string `json:"datasource"`
}

// RedisConfig ...
type RedisConfig struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Size     int    `json:"size"`
}
