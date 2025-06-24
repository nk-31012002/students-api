package config

type HTTPServer struct {
	Addr string
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" `
	StoragePath string `yaml:"env" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
