package configs

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func IsProduction() bool {
	isProd := os.Getenv("IS_PROD")
	if len(isProd) == 0 {
		return false
	}
	isLocalBoolean, _ := strconv.ParseBool(isProd)
	return isLocalBoolean
}
func IsLocal() bool {
	isLocal := os.Getenv("IS_LOCAL")
	if len(isLocal) == 0 {
		return true
	}
	isLocalBoolean, _ := strconv.ParseBool(isLocal)
	return isLocalBoolean
}

func getEnvOrDefault(key, pattern string) string {
	value := os.Getenv(key) //always return string
	if len(value) == 0 {
		return viper.GetString(pattern)
	}
	return value
}

func GetDBName() string {
	return getEnvOrDefault("MONGODB__DATABASE", "mongodb.database")
}

func GetMongoHost() string {
	return getEnvOrDefault("MONGODB__HOST", "mongodb.host")
}

func GetMongoUsername() string {
	return getEnvOrDefault("MONGODB__USER", "mongodb.username")
}

func GetMongoPassword() string {
	return getEnvOrDefault("MONGODB__PASS", "mongodb.password")
}

func GetRedisAddr() string {
	return getEnvOrDefault("REDIS__HOST", "redis.host")
}

func GetRedisUserName() string {
	return getEnvOrDefault("REDIS__USER", "redis.username")
}

func GetRedisPassword() string {
	if IsProduction() {
		return ""
	}

	return getEnvOrDefault("REDIS__PASS", "redis.pass")
}

func GetRabbitMQInstance() string {
	return getEnvOrDefault("RABBIT__AMQP", "rabbit_instance.uri")
}
