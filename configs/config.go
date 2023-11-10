package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	//DBHost        string `mapstructure:"DB_HOST"`
	//DBPort        string `mapstructure:"DB_PORT"`
	//DBUser        string `mapstructure:"DB_USER"`
	//DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort int    `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET" `
	JwtExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	AppName       string `mapstructure:"APP_NAME"`
}

var Conf *conf

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.AppName == "" {
		cfg.AppName = "EstudoGo"
	}

	Conf = cfg
	return cfg, err
}
