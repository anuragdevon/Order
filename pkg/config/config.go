package config

import "github.com/spf13/viper"

type Config struct {
	Port            string `mapstructure:"PORT"`
	DBUrl           string `mapstructure:"DB_URL"`
	InventorySvcUrl string `mapstructure:"INVENTORY_SVC_URL"`
	DbUser          string `mapstructure:"DB_USER"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbName          string `mapstructure:"DB_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("/home/anurag/docs/swiggy/order/pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
