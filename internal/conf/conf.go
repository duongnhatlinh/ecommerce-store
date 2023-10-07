package conf

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	hostKey       = "HOST"
	portKey       = "PORT"
	dbHostKey     = "DB_HOST"
	dbPortKey     = "DB_PORT"
	dbNameKey     = "DB_NAME"
	dbUserKey     = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	jwtSecretKey  = "JWT_SECRET"
)

type Config struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
}

func NewConfig() Config {
	err := godotenv.Load("E:\\project\\fullstack website ecommerce store\\backend\\.env")
	if err != nil {
		panic(err)
	}
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}
	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		logAndPanic(portKey)
	}
	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logAndPanic(dbHostKey)
	}
	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		logAndPanic(dbPortKey)
	}
	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logAndPanic(dbNameKey)
	}
	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logAndPanic(dbUserKey)
	}
	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logAndPanic(dbPasswordKey)
	}
	jwtSecret, ok := os.LookupEnv(jwtSecretKey)
	if !ok || jwtSecret == "" {
		logAndPanic(jwtSecretKey)
	}

	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
	}
}

func NewTestConfig() Config {
	testConfig := NewConfig()
	testConfig.DbName = testConfig.DbName + "_test"
	return testConfig
}

func logAndPanic(envVar string) {
	log.Panic().Str("envVar", envVar).Msg("ENV variable not set or value not valid")
}
