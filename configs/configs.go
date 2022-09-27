package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var SchemaConfig Schema

func InitConfig() {

	err := loadConfigByViper("config", &SchemaConfig)
	if err != nil {
		panic(err)
	}
}

func loadConfigByViper(name string, config interface{}) (err error) {
	viper.SetConfigName(name)
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")      // Optionally look for config in the working directory.
	viper.AddConfigPath("../config/")   // Look for config needed for tests.
	viper.AddConfigPath("../")          // Look for config needed for tests.
	viper.AddConfigPath("../../config") // Look for config needed for tests.
	viper.AddConfigPath("../../../config")
	viper.AddConfigPath("configs/") // Optionally look for config in the working directory.

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if config != nil { //parse data
		err = viper.Unmarshal(config)
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	return
}
