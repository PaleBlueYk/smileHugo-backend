package config

import (
	"github.com/PaleBlueYk/smileHugo-backend/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Application string `yaml:"application"`
	Repository  string `yaml:"repository"`
	Port        string `yaml:"port"`
	RunMode     string `yaml:"runMode"`

	Logger   Logger   `yaml:"logger"`
	Database Database `yaml:"database"`
}

type Logger struct {
	Level         string `yaml:"level"`
	FilePath      string `yaml:"filePath"`
	ErrDetail     string `yaml:"errDetail"`
	ErrInResponse string `yaml:"errInResponse"`
}

type Database struct {
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	ShowSql  bool   `yaml:"showSql"`
}

var AppConfig Config

func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err)
		return err
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		logger.Error(err)
	}
	logger.Info(AppConfig)
	return nil
}
