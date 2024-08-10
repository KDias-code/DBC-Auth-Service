package configs

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDb       string `mapstructure:"REDIS_DB"`
	PostgresDb    string `mapstrcuture:"POSTGRES_DB"`
}

func LoadConfigs() (*Config, error) {
	conf := new(Config)

	v := viper.New()
	v.AutomaticEnv()

	err := v.BindEnv("PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_ADDRESS")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_PASSWORD")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_DB")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("POSTGRES_DB")
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}
