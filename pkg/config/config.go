package config

type Config struct {
	HostPort string
	// DB *DBConfig
}

// type DBConfig struct {
// 	Host string
// }

func GetConfig() *Config {
	return &Config{
		HostPort: "3000",
		// DB: &DBConfig{
		// 	Host: "mysql",
		// },
	}
}
