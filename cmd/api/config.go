package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

type InfluxDBConfig struct {
	URL    string
	Port   int
	Org    string
	Token  string
	Bucket string
}

type MySQLConfig struct {
	DSN        string
	MaxRetries int
	RetryDelay int
}

type DBConfig struct {
	InfluxDB InfluxDBConfig
	MySQL    MySQLConfig
}

type LimiterConfig struct {
	Enabled bool
	RPS     float64
	Burst   int
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}

type Config struct {
	Port    int
	Env     string
	DB      DBConfig
	Limiter LimiterConfig
	SMTP    SMTPConfig
}

func NewConfig() *Config {

	loadEnv := flag.Bool("loadEnv", false, "Set to true to load env var from a .env file")
	flag.Parse()
	if *loadEnv {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err.Error())
		}
	}

	return &Config{
		Port: getEnvStrToI("BACKEND_PORT", 4000),
		Env:  getEnvStr("BACKEND_ENVIRONMENT", "developement"),
		DB: DBConfig{
			InfluxDB: InfluxDBConfig{
				URL:    getEnvStr("INFLUXDB_URL", "", true),
				Port:   getEnvStrToI("INFLUXDB_PORT", 0, true),
				Org:    getEnvStr("INFLUXDB_ORG", "", true),
				Token:  getEnvStr("INFLUXDB_TOKEN", "", true),
				Bucket: getEnvStr("INFLUXDB_BUCKET", "", true),
			},
			MySQL: MySQLConfig{
				DSN:        getEnvStr("MYSQL_DSN", "", true),
				MaxRetries: getEnvStrToI("MYSQL_MAX_RETRIES", 3),
				RetryDelay: getEnvStrToI("MYSQL_RETRY_DELAY", 15),
			},
		},
		Limiter: LimiterConfig{
			Enabled: getEnvBool("LIMITER_ENABLED", true),
			RPS:     float64(getEnvStrToI("LIMITER_RPS", 2)),
			Burst:   getEnvStrToI("LIMITER_BURST", 4),
		},
		SMTP: SMTPConfig{
			Host:     getEnvStr("SMTP_HOST", "", true),
			Port:     getEnvStrToI("SMTP_PORT", 0, true),
			Username: getEnvStr("SMTP_USERNAME", "", true),
			Password: getEnvStr("SMTP_PASSWORD", "", true),
			Sender:   getEnvStr("SMTP_SENDER", "", true),
		},
	}
}
