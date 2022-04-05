package util

type DatabaseConfig struct {
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBName     string `mapstructure:"db_name"`
	DBPassword string `mapstructure:"db_pass"`
	DBSSLMode  string `mapstructure:"db_ssl_mode"`
}

type BolorLoginConfig struct {
	Token   string
	Key     string
	AuthURL string `mapstructure:"auth_url"`
	Name    string
}

type Config struct {
	BolorLogin BolorLoginConfig `mapstructure:"bolorlogin"`
	DB         DatabaseConfig   `mapstructure:"database"`
	Debug      bool             `mapstructure:"debug"`
	Service    ServiceConfig    `mapstructure:"service"`
}

type ServiceConfig struct {
	Port   string
	URL    string
	WebURL string `mapstructure:"web_url"`
	ApiURL string `mapstructure:"api_url"`
}
