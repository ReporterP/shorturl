package config

type Config struct {
    ServerAddress string `env:"SERVER_ADDRESS"`
    BaseURL string `env:"BASE_URL"`
    EnvLogLevel string `env:"LOG_LEVEL"`
}