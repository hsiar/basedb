package gormdao

import "time"

type Config struct {
	Dsn         string
	TablePrefix string

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
}

func (this *Config) HasDsn() bool {
	return this.Dsn != ""
}

func DefaultConfig() *Config {
	return &Config{
		TablePrefix:     "z_",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: time.Hour,
	}
}
