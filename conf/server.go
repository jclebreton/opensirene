package conf

// Server is the structure that holds the server configuration
type Server struct {
	Host  string `yaml:"host" env:"SERVER_HOST" default:"127.0.0.1"`
	Port  int    `yaml:"port" env:"SERVER_PORT" default:"8080"`
	Debug bool   `yaml:"debug" env:"SERVER_DEBUG"`
}
