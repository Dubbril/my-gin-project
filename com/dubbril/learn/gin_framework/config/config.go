package config

type DbConfig struct {
	Postgres struct {
		Connection string `mapstructure:"connection"`
	} `mapstructure:"postgres"`
}
