package db

type DBConfig struct {
	DatabaseURI string
}

func NewDBConfig(url string) *DBConfig {
	return &DBConfig{
		DatabaseURI: url,
	}
}
