package conf

// Server is the structure that holds the server configuration
type Server struct {
	Host  string `yaml:"host" env:"SERVER_HOST" default:"127.0.0.1"`
	Port  int    `yaml:"port" env:"SERVER_PORT" default:"8080"`
	Debug bool   `yaml:"debug" env:"SERVER_DEBUG"`
	Cors  Cors   `yaml:"cors"`
}

// Cors is a simple structure holding the information about CORS configuration
type Cors struct {
	AllowOrigins   []string `yaml:"allow_origins"`
	Enabled        bool     `yaml:"enabled" env:"CORS_ENABLED"`
	PermissiveMode bool     `yaml:"permissive_mode" env:"CORS_PERMISSIVE"`
}

func (sc Server) DebugMode() bool {
	return sc.Debug
}

func (sc Server) GetPort() int {
	return sc.Port
}

func (sc Server) GetHost() string {
	return sc.Host
}
