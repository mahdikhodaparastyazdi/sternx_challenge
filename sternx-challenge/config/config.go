package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	config Config
	logger = log.With().Str("task", "stern-x").Logger()
)

type (
	Config struct {
		Log    Log    `mapstructure:"log" validate:"required"`
		DB     DB     `mapstructure:"db" validate:"required"`
		Server Server `mapstructure:"server" validate:"required"`
	}
	DB struct {
		DriverName     string `mapstructure:"driverName" validate:"required"`
		DataSourceName string `mapstructure:"dataSourceName" validate:"required"`
	}
	Log struct {
		Level int `mapstructure:"level"`
	}
	Server struct {
		Addr string `mapstructure:"addr" validate:"required"`
	}
)

func Get() *Config {
	return &config
}

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

func Load(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		fmt.Print("asdada")
		return false
	}
	fmt.Println(file)
	viper.SetConfigFile(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	// viper.SetEnvPrefix(Namespace)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadConfig(bytes.NewReader([]byte(Default))); err != nil {
		logger.Error().Err(err).Msgf("error loading default configs")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info().Msgf("Config file changed %s", file)
		reload(e.Name)
	})

	return reload(file)
}

func reload(file string) bool {
	err := viper.MergeInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error().Err(err).Msgf("config file not found %s", file)
		} else {
			logger.Error().Err(err).Msgf("config file read failed %s", file)
		}
		return false
	}

	err = viper.GetViper().UnmarshalExact(&config)
	if err != nil {
		logger.Error().Err(err).Msgf("config file loaded failed %s", file)
		return false
	}

	if err = config.Validate(); err != nil {
		logger.Error().Err(err).Msgf("invalid configuration %s", file)
	}

	logger.Info().Msgf("Config file loaded %s", file)
	return true
}
