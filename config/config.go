package config

import "os"

type Environment struct {
    DatabaseURL string
    APIKey      string
    GinMode       string
}

func LoadEnvironment() Environment {
    return Environment{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        APIKey:      os.Getenv("API_KEY"),
        GinMode:       os.Getenv("GIN_MODE"),
    }
}