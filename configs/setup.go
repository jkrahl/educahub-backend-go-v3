package configs

import (
	"log"

	gin "github.com/gin-gonic/gin"
	viper "github.com/spf13/viper"
)

func init() {
	SetupViper()
}

func SetupViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	if gin.Mode() != gin.ReleaseMode {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	viper.AutomaticEnv()
}

func GetViperString(key string) string {
	return viper.GetString(key)
}
