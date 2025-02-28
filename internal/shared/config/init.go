package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetupBackendConfig(name1, name2, name3 string) error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName(name1)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading first config file: %s", err)
	}

	viper.SetConfigName(name2)
	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("error reading second config file: %s", err)
	}

	viper.SetConfigName(name3)
	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("error reading third config file: %s", err)
	}
	return nil
}

func SetupPingerConfig(name1, name2 string) error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName(name1)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading first config file: %s", err)
	}

	viper.SetConfigName(name2)
	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("error reading second config file: %s", err)
	}
	return nil
}
