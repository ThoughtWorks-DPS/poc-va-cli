package vconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadViperConfig() {
	viper.AddConfigPath(".")
	viper.SetDefault("api_service_base_url", "http://api.devportal.name")
	viper.AutomaticEnv()
	viper.SafeWriteConfigAs("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}