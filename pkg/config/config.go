package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotToken      string
	ConsumerKey   string
	AuthServerUrl string
	BotUrl        string `mapstructure:"bot_url"`
	DbPath        string `mapstructure:"db_path"`

	Messages Messages
}

type Messages struct {
	Responses Responses
	Errors    Errors
}

type Errors struct {
	InvalidLink  string `mapstructure:"invalid_link"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
	Default      string `mapstructure:"default"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("bot_token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}

	cfg.BotToken = viper.GetString("bot_token")
	cfg.ConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerUrl = viper.GetString("auth_server_url")

	return nil
}
