package config

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    SendGridAPIKey      string
    ExchangeRateAPIKey  string
    DBUser             string
    DBPassword         string
    DBHost             string
    DBPort             string
    DBName             string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        SendGridAPIKey:     os.Getenv("SENDGRID_API_KEY"),
        ExchangeRateAPIKey: os.Getenv("EXCHANGERATE_API_KEY"),
        DBUser:            os.Getenv("DB_USER"),
        DBPassword:        os.Getenv("DB_PASSWORD"),
        DBHost:            os.Getenv("DB_HOST"),
        DBPort:            os.Getenv("DB_PORT"),
        DBName:            os.Getenv("DB_NAME"),
    }, nil
}

func (c *Config) GetDBConnString() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}