package apiserver

type Config struct {
	BinAddr string
}

func NewConfig() *Config {
	return &Config{
		BinAddr: ":8080",
	}
}
