package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	// expose this variable to global scope
	Env Configuration
)

const (
	projectDirName = "myshopify"
	ModeDev        = "debug"
)

func init() {
	currentWorkDirectory, _ := os.Getwd()
	confFile := fmt.Sprintf("%s/conf.yml", currentWorkDirectory)
	err := cleanenv.ReadConfig(confFile, &Env)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Load %s failed, exit", confFile))
	}
}

type Configuration struct {
	Mode string `env:"MODE" env-required:"true" env-default:"debug"`

	Log struct {
		File  string `env:"LOG_FILE_NAME" env-description:"Service log file name" env-required:"true"`
		Level string `env:"LOG_LEVEL" env-description:"Service log level" env-default:"info"`
	} `yaml:"log"`

	Shopify struct {
		APIKey       string `yaml:"api_key" env:"SHOPIFY_API_KEY"`
		SharedSecret string `yaml:"shared_secret" env:"SHOPIFY_SHARED_SECRET"`
		RedirectURI  string `yaml:"redirect_uri" env:"SHOPIFY_REDIRECT_URL"`
		Scopes       string `yaml:"scopes" env:"SHOPIFY_SCOPES"`
	} `yaml:"shopify"`

	Database struct {
		Host     string `yaml:"host" env:"DB_HOST" env-description:"Database host address" env-default:"localhost"`
		Port     string `yaml:"port" env:"DB_PORT" env-description:"Database host port" env-default:"5432"`
		Name     string `yaml:"name" env:"DB_NAME" env-description:"Database name" env-default:"reserveit"`
		Username string `yaml:"username" env:"DB_USERNAME" env-description:"Database username" env-default:"qe"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-description:"Database password"`
	} `yaml:"database"`

	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST" env-description:"Redis host address" env-default:"127.0.0.1"`
		Port     int    `yaml:"port" env:"REDIS_PORT" env-description:"Redis service port" env-default:"6379"`
		Db       int    `yaml:"db" env:"REDIS_DB" env-description:"Redis db index" env-default:"0"`
		PoolSize int    `yaml:"pool_size" env:"REDIS_POOL_SIZE" env-description:"Redis connection pool size" env-default:"30"`
	} `yaml:"redis"`

	Webhooks struct {
		RouterPrefix string `yaml:"router_prefix" env:"ROUTER_PREFIX" env-description:"Router prefix" env-default:"/webhooks"`
	} `yaml:"webhooks"`
}
