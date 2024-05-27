package config

type Database struct {
	Port     string `yaml:"port"  env-default:"5432"`
	Host     string `yaml:"host"  env-default:"localhost"`
	Name     string `yaml:"name"  env-default:"postgres"`
	User     string `yaml:"user"  env-default:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	DB Database `yaml:"database"`
}
