package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TwitterBearerToken string `mapstructure:"TWITTER_BEARER_TOKEN"`
	TwitterStreamRule string `mapstructure:"TWITTER_RULE"`
	RollCapcity int `mapstructure:"ROLL_CAPACITY"`
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		} else {
			viper.Set("TWITTER_BEARER_TOKEN", viper.GetString("TWITTER_BEARER_TOKEN"))
			viper.Set("TWITTER_RULE", viper.GetString("TWITTER_RULE"))
			viper.Set("ROLL_CAPACITY", viper.GetInt("ROLL_CAPACITY"))
		}
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}