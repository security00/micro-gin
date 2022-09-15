package global

import (
	"github.com/spf13/viper"
	"micro-gin/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var App = new(Application)