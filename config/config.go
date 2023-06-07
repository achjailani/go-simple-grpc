package config

import (
	"os"
	"strconv"
	"strings"
)

type DBConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBTimeZone string
	DBLog      bool
}

// RedisConfig is a struct which hold redis config fields
type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// RedisTestConfig is a struct which hold redis test config fields
type RedisTestConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// CacheConfig is a cache config for test
type CacheConfig struct {
	CacheHost     string
	CachePort     string
	CacheUsername string
	CachePassword string
	CacheDB       int
}

// TestCacheConfig is a cache config for test
type TestCacheConfig struct {
	CacheHost     string
	CachePort     string
	CacheUsername string
	CachePassword string
	CacheDB       int
}

// StorageConfig is a struct
type StorageConfig struct {
	StorageProjectID       string
	StorageEndpoint        string
	StorageHost            string
	StorageBucketName      string
	StorageAccessKeyID     string
	StorageAccessKeySecret string
}

type Config struct {
	AppName     string
	AppPort     int
	AppEnv      string
	AppLang     string
	AppTimeZone string
	GRPCPort    int
	TestMode    bool
	DBConfig
	DBTestConfig DBConfig
	RedisConfig
	RedisTestConfig
	CacheConfig
	TestCacheConfig
	StorageConfig
}

func New() *Config {
	return &Config{
		AppName:     getEnv("APP_NAME", ""),
		AppPort:     getEnvAsInt("APP_PORT", 8080),
		AppEnv:      getEnv("APP_ENV", "development"),
		AppLang:     getEnv("APP_LANG", "en"),
		AppTimeZone: getEnv("APP_TIMEZONE", ""),
		GRPCPort:    getEnvAsInt("GRPC_PORT", 9000),
		TestMode:    getEnvAsBool("TEST_MODE", false),
		DBConfig: DBConfig{
			DBDriver:   getEnv("DB_DRIVER", "postgres"),
			DBHost:     getEnv("DB_HOSt", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBName:     getEnv("DB_NAME", "db"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASS", ""),
			DBTimeZone: getEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:      getEnvAsBool("ENABLE_LOGGER", true),
		},
		DBTestConfig: DBConfig{
			DBDriver:   getEnv("TEST_DB_DRIVER", "postgres"),
			DBHost:     getEnv("TEST_DB_HOSt", "localhost"),
			DBPort:     getEnv("TEST_DB_PORT", "5432"),
			DBName:     getEnv("TEST_DB_NAME", "db"),
			DBUser:     getEnv("TEST_DB_USER", "postgres"),
			DBPassword: getEnv("TEST_DB_PASS", ""),
			DBTimeZone: getEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:      getEnvAsBool("ENABLE_LOGGER", false),
		},
		RedisConfig: RedisConfig{
			RedisHost:     getEnv("REDIS_HOST", "localhost"),
			RedisPort:     getEnv("REDIS_PORT", "6379"),
			RedisPassword: getEnv("REDIS_PASSWORD", ""),
			RedisDB:       getEnvAsInt("REDIS_DB", 0),
		},
		RedisTestConfig: RedisTestConfig{
			RedisHost:     getEnv("TEST_REDIS_HOST", "localhost"),
			RedisPort:     getEnv("TEST_REDIS_PORT", "6379"),
			RedisPassword: getEnv("TEST_REDIS_PASSWORD", ""),
			RedisDB:       getEnvAsInt("TEST_REDIS_HOST", 0),
		},
		CacheConfig: CacheConfig{
			CacheHost:     getEnv("CACHE_HOST", "127.0.0.1"),
			CachePort:     getEnv("CACHE_PORT", "6379"),
			CacheUsername: getEnv("CACHE_USERNAME", ""),
			CachePassword: getEnv("CACHE_PASSWORD", ""),
			CacheDB:       getEnvAsInt("CACHE_DB", 0),
		},
		TestCacheConfig: TestCacheConfig{
			CacheHost:     getEnv("TEST_CACHE_HOST", "127.0.0.1"),
			CachePort:     getEnv("TEST_CACHE_PORT", "6379"),
			CacheUsername: getEnv("TEST_CACHE_USERNAME", ""),
			CachePassword: getEnv("TEST_CACHE_PASSWORD", ""),
			CacheDB:       getEnvAsInt("TEST_CACHE_DB", 0),
		},
		StorageConfig: StorageConfig{
			StorageProjectID:       getEnv("STORAGE_PROJECT_ID", ""),
			StorageEndpoint:        getEnv("STORAGE_ENDPOINT", "127.0.0.1:9000"),
			StorageHost:            getEnv("STORAGE_HOST", "127.0.0.1:9000"),
			StorageBucketName:      getEnv("STORAGE_BUCKET_NAME", ""),
			StorageAccessKeyID:     getEnv("STORAGE_ACCESS_KEY_ID", ""),
			StorageAccessKeySecret: getEnv("STORAGE_ACCESS_KEY_SECRET", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultVal
	}

	value := strings.Split(valueStr, sep)

	return value
}
