package config

import (
    "log"

    "github.com/kelseyhightower/envconfig"
)

type Configuration struct {
    Address        string `default:":8080"`
    TokenSecretKey string `default:"secret" split_words:"true"`
}

var Config Configuration

func init() {
    err := envconfig.Process("app", &Config)
    if err != nil {
        log.Fatal(err.Error())
    }
}
