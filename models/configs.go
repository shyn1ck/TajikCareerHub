package models

type AppConfig struct {
	AuthParams     AuthParams     `json:"auth"`
	LogParams      LogParams      `json:"log_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
}

type AuthParams struct {
	JwtTtlMinutes int `json:"jwt_ttl_minutes"`
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`
	LogInfo          string `json:"log_info"`
	LogError         string `json:"log_error"`
	LogWarn          string `json:"log_warn"`
	LogDebug         string `json:"log_debug"`
	MaxSizeMegabytes int    `json:"max_size_megabytes"`
	MaxBackups       int    `json:"max_backups"`
	MaxAge           int    `json:"max_age"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"local_time"`
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`
	PortRun    string `json:"port_run"`
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
}

type PostgresParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
}
