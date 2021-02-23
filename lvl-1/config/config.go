package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
)

type Config struct {
	Port        string `json:"port" yaml:"port"`
	DbURL       string `json:"db_url" yaml:"db_url"`
	JaegerURL   string `json:"jaeger_url" yaml:"jaeger_url"`
	SentryURL   string `json:"sentry_url" yaml:"sentry_url"`
	KafkaBroker string `json:"kafka_broker" yaml:"kafka_broker"`
	SomeAppID   string `json:"some_app_id" yaml:"some_app_id"`
	SomeAppKey  string `json:"some_app_key" yaml:"some_app_key"`
}

func FromJSON(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	config := Config{}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		//log.Printf("Не могу декодировать json-файл в структуру: %v", err)
		return nil, err
	}
	if err = validateConfiguration(config); err != nil {
		return nil, err
	}
	return &config, nil
}

func FromYAML(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	config := Config{}

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		//log.Printf("Не могу декодировать yaml-файл в структуру: %v", err)
		return nil, err
	}
	if err = validateConfiguration(config); err != nil {
		return nil, err
	}

	return &config, nil
}

func FromEnvironment() (*Config, error) {
	config := Config{}
	config.Port = os.Getenv("PORT")
	config.DbURL = os.Getenv("DB_URL")
	config.JaegerURL = os.Getenv("JAEGER_URL")
	config.SentryURL = os.Getenv("SENTRY_URL")
	config.KafkaBroker = os.Getenv("KAFKA_BROKER")
	config.SomeAppID = os.Getenv("SOME_APP_ID")
	config.SomeAppKey = os.Getenv("SOME_APP_KEY")
	if err := validateConfiguration(config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfiguration(config Config) error {
	if config.DbURL == "" || config.SomeAppID == "" || config.SomeAppKey == "" {
		return errors.New("обязательное поле не задано")
	}

	if !IsUrl(config.DbURL) {
		return errors.New("DbURL должно быть ссылкой")
	}

	if config.SentryURL != "" && !IsUrl(config.SentryURL) {
		return errors.New("SentryURL должно быть ссылкой")
	}

	if config.JaegerURL != "" && !IsUrl(config.JaegerURL) {
		return errors.New("SentryURL должно быть ссылкой")
	}
	return nil
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
