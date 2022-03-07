package config

type Config struct {
	Services ServicesConfigs `toml:"services"`
}

type ServicesConfigs struct {
	Server ServerConfig   `toml:"server"`
	Db     DatabaseConfig `toml:"db"`
}

type ServerConfig struct {
	Host string `toml:"HOST"`
	Port string `toml:"PORT"`
}

type DatabaseConfig struct {
	Url string `toml:"URL"`
}
