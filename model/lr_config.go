package model

type LrConfig struct {
	LRServerUrl          string `json:"lr_server_url"`
	LRServerUserNumber   string `json:"lr_server_user_number"`
	LRServerUserPassword string `json:"lr_server_user_password"`
	DebugMode            bool   `json:"debug_mode"`
	TimerPingInterval    int    `json:"timer_ping_interval"`
}
