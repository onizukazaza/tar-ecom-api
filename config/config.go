package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
	"sync"
	"time"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}


	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    int           `mapstructure:"bodylimit" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		JWTSecretKey string        `mapstructure:"jwt_secret_key" validate:"required"`
	}

	// OAuth2 struct {
	// }

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

)

// start app set  1 load > "sync" callback function synchroton
var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config { //undo function
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./") // ./ -> ./bin (deploy)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // . -> _

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
		validating := validator.New()
		if err := validating.Struct(configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}
