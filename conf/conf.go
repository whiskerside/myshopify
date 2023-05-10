package conf

import (
	"os"
	"regexp"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/whiskerside/myshopify/log"
)

var (
	// expose this variable to global scope
	Env Configuration
)

const (
	projectDirName = "myshopify"
	ModeDev        = "dev"
)

func init() {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	envFile := string(rootPath) + `/.env`
	err := cleanenv.ReadConfig(envFile, &Env)
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("Server init failed")
	}
}

type Configuration struct {
	Mode                string `env:"MODE" env-required:"true" env-default:"dev"`
	ShopifyAPIKey       string `env:"SHOPIFY_API_KEY" env-required:"true"`
	ShopifySharedSecret string `env:"SHOPIFY_SHARED_SECRET" env-required:"true"`
	ShopifyRedirectURL  string `env:"SHOPIFY_REDIRECT_URL" env-required:"true"`
	ShopifyScopes       string `env:"SHOPIFY_SCOPES" env-required:"true"`
	LogFileName         string `env:"LOG_FILE_NAME" env-description:"Service log file name" env-required:"true"`

	DbHost     string `env:"DB_HOST" env-description:"Database host address" env-default:"localhost"`
	DbPort     string `env:"DB_PORT" env-description:"Database host port" env-default:"5432"`
	DbName     string `env:"DB_NAME" env-description:"Database name" env-default:"reserveit"`
	DbUserName string `env:"DB_USERNAME" env-description:"Database username" env-default:"qe"`
	DbPassword string `env:"DB_PASSWORD" env-description:"Database password"`

	RedisHost     string `env:"REDIS_HOST" env-description:"Redis host address" env-default:"127.0.0.1"`
	RedisPort     int    `env:"REDIS_PORT" env-description:"Redis service port" env-default:"6379"`
	RedisDb       int    `env:"REDIS_DB" env-description:"Redis db index" env-default:"0"`
	RedisPoolSize int    `env:"REDIS_POOL_SIZE" env-description:"Redis connection pool size" env-default:"30"`

	APIPort int `env:"API_PORT" env-description:"API service port" env-default:"9000"`
}
