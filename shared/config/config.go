package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AllConfig struct {
	TimeoutConfig *TimeoutConfig
	PGSQLConfig   *PGSQLConfig
	JWTConfig     *JWTConfig
}

type TimeoutConfig struct {
	TIMEOUT_SUPER_SHORT int
	TIMEOUT_SHORT       int
	TIMEOUT_MEDIUM      int
	TIMEOUT_LONG        int
	TIMEOUT_SUPER_LONG  int
}

type JWTConfig struct {
	JWT_SECRET                         string
	JWT_EXPIRE_AT                      int
	JWT_ISSUER                         string
	JWT_AUDIENCE                       []string
	JWT_REFRESH_EXPIRE_AT              int
	JWT_REFRESH_ALLOW_BEFORE_EXPIRE_AT float64
}

type PGSQLConfig struct {
	PG_HOST            string
	PG_PORT            string
	PG_DBNAME          string
	PG_USER            string
	PG_PASSWORD        string
	PG_SSLMODE         string
	PG_MaxIdleConn     int
	PG_MaxOpenConn     int
	PG_ConnMaxLifetime int
	PG_ConnMaxIdleTime int
}

func InitConfig() *AllConfig {
	if err := godotenv.Load(".env."); err != nil {
		log.Fatal("Error loading .env file")
	}

	timeoutConfig := &TimeoutConfig{
		TIMEOUT_SUPER_SHORT: getIntEnv("TIMEOUT_SUPER_SHORT"),
		TIMEOUT_SHORT:       getIntEnv("TIMEOUT_SHORT"),
		TIMEOUT_MEDIUM:      getIntEnv("TIMEOUT_MEDIUM"),
		TIMEOUT_LONG:        getIntEnv("TIMEOUT_LONG"),
		TIMEOUT_SUPER_LONG:  getIntEnv("TIMEOUT_SUPER_LONG"),
	}

	jwtConfig := &JWTConfig{
		JWT_SECRET:                         getEnv("JWT_SECRET"),
		JWT_EXPIRE_AT:                      getIntEnv("JWT_EXPIRE_AT"),
		JWT_ISSUER:                         getEnv("JWT_ISSUER"),
		JWT_AUDIENCE:                       getSliceEnv("JWT_AUDIENCE"),
		JWT_REFRESH_EXPIRE_AT:              getIntEnv("JWT_REFRESH_EXPIRE_AT"),
		JWT_REFRESH_ALLOW_BEFORE_EXPIRE_AT: getFloat64Env("JWT_REFRESH_ALLOW_BEFORE_EXPIRE_AT"),
	}

	password := getEnv("PGSQL_PASSWORD")
	//	var passwordBytes = []byte(password)

	// Enkripsi ke Base64
	//pgsqlPass := base64.StdEncoding.EncodeToString(passwordBytes)

	pgsqlConfig := &PGSQLConfig{
		PG_HOST:            getEnv("PGSQL_HOST"),
		PG_PORT:            getEnv("PGSQL_PORT"),
		PG_DBNAME:          getEnv("PGSQL_DBNAME"),
		PG_USER:            getEnv("PGSQL_USER"),
		PG_PASSWORD:        password,
		PG_SSLMODE:         getEnv("PGSQL_SSLMODE"),
		PG_MaxIdleConn:     getIntEnv("PGSQL_MaxIdleConn"),
		PG_MaxOpenConn:     getIntEnv("PGSQL_MaxOpenConn"),
		PG_ConnMaxLifetime: getIntEnv("PGSQL_ConnMaxLifetime"),
		PG_ConnMaxIdleTime: getIntEnv("PGSQL_ConnMaxIdleTime"),
	}

	return &AllConfig{
		TimeoutConfig: timeoutConfig,
		PGSQLConfig:   pgsqlConfig,
		JWTConfig:     jwtConfig,
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env variable %s not set", key)
	}
	return value
}

func getIntEnv(key string) int {
	valueStr := getEnv(key)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Error parsing int from env variable %s: %s", key, err)
	}
	return valueInt
}

func getFloat64Env(key string) float64 {
	valueStr := getEnv(key)
	valueFloat, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatalf("Error parsing float64 from env variable %s: %s", key, err)
	}
	return valueFloat
}

func getSliceEnv(key string) []string {
	value := getEnv(key)
	return strings.Split(value, ",")
}
