package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

type (
	Config struct {
		AccessKeyVK  string `env:"VK_ACCESS_TOKEN"`
		AccessKeyTG  string `env:"TG_TOKEN"`
		AccessKeyOCD string `env:"OCD_TOKEN"`
		AuthURL      string `env:"AUTH_URL"`
	}

	ServerConfig struct {
		AuthServerURL string `env:"AUTH"`
		ClientID      string `env:"CLIENT_ID"`
		ClientSecret  string `env:"CLIENT_SECRET"`
		Redirect      string `env:"REDIRECT"`
		Port string `env:"PORT"`
	}
)

func Init() (*Config, *ServerConfig, error) {
	var (
		cfg       Config
		cfgServer ServerConfig
	)

	re := regexp.MustCompile(`^(.*` + "FerrisWheel" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		err.Error()
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, nil, err
	}

	err = cleanenv.ReadEnv(&cfgServer)
	if err != nil {
		return nil, nil, err
	}

	return &cfg, &cfgServer, nil
}
