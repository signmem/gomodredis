package g

type GlobalConfig struct {
	Debug         bool              `json:"debug"`
	LogFile 		string			`json:"logfile"`
	Http 			*HttpConfig 	`json:"http"`
	Redis 			*RedisConfig 	`json:"redis"`
	Conncurrency	int				`json:"conncurrency"`
}

type HttpConfig struct {
	Enabled 		bool 		`json:"enabled"`
	Listen 			string 		`json:"listen"`
}

type RedisConfig struct {
	Server 			string 		`json:"server"`
	MaxIdle 		int 		`json:"maxidle"`
	MaxActive 		int 		`json:"maxactive"`
	IdleTimeout 	int 		`json:"idletimeout"`
	MaxConnLifetime int 		`json:"maxconnlifetime"`
}