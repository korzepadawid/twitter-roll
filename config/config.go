package config

import "github.com/spf13/viper"

type Config struct {
	TwitterBearerToken string`mapstructure:"TWITTER_BEARER_TOKEN"`
	TwitterStreamRule string`mapstructure:"TWITTER_RULE"`
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}