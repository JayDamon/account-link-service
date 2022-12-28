package config

type Config struct {
	HostPort string
}

func GetConfig() *Config {
	return &Config{
		HostPort: "3000",
	}
}
