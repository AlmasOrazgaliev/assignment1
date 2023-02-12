package apiserver

type Config struct {
	Port        string
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		Port:        ":8080",
		DatabaseURL: "host=localhost port=5432 user=postgres password=alma45884 dbname=godb sslmode=disable",
	}
}
