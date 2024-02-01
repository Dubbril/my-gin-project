package config

type DatabaseConfig struct {
	Host       string
	User       string
	DBName     string
	SSLMode    string
	Password   string
	SearchPath string
}

// NewDatabaseConfig initializes a new DatabaseConfig with default values
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:       "localhost",
		User:       "postgres",
		DBName:     "postgres",
		SSLMode:    "disable",
		Password:   "bit@1234",
		SearchPath: "test",
	}
}
