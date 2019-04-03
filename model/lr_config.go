package model

type LrConfig struct {
	LRServerHost         string `json:"lr_server_host"`
	LRServerPort         int    `json:"lr_server_port"`
	LRServerUserNumber   string `json:"lr_server_user_number"`
	LRServerUserPassword string `json:"lr_server_user_password"`
	DebugMode            bool   `json:"debug_mode"`
	TimerPingInterval    int    `json:"timer_ping_interval"`
}
