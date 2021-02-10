package config

import (
	"os"
	"reflect"
	"testing"
)

var valid = Config{
	Port:        "8080",
	DbUrl:       "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
	JaegerUrl:   "http://jaeger:16686",
	SentryUrl:   "http://sentry:9000",
	KafkaBroker: "kafka:9092",
	SomeAppId:   "testid",
	SomeAppKey:  "testkey",
}

func TestFromYaml(t *testing.T) {
	config, err := FromYaml("valid.yaml")
	if err != nil {
		t.Errorf("Неверный результат выполнения на valid.yaml: %v", err)
	}
	if reflect.DeepEqual(valid, config) {
		t.Errorf("Неверный результат выполнения на valid.yaml: %v", config)
	}
	config, err = FromYaml("invalid.yaml")
	if err == nil {
		t.Errorf("Неверный результат выполнения на invalid.yaml: %v", config)
	}
	config, err = FromYaml("incomplete.yaml")
	if err == nil {
		t.Errorf("Неверный результат выполнения на incomplete.yaml: %v", config)
	}
	config, err = FromYaml("invalid_url.yaml")
	if err == nil {
		t.Errorf("Неверный результат выполнения на incomplete.yaml: %v", config)
	}
	config, err = FromYaml("nonexistent.yaml")
	if err == nil {
		t.Errorf("Неверный результат выполнения на nonexistent.yaml: %v", config)
	}
}

func TestFromJson(t *testing.T) {
	config, err := FromJson("valid.json")
	if err != nil {
		t.Errorf("Неверный результат выполнения на valid.json: %v", err)
	}
	if reflect.DeepEqual(valid, config) {
		t.Errorf("Неверный результат выполнения на valid.json: %v", config)
	}
	config, err = FromJson("invalid.json")
	if err == nil {
		t.Errorf("Неверный результат выполнения на invalid.json: %v", config)
	}
	config, err = FromJson("incomplete.json")
	if err == nil {
		t.Errorf("Неверный результат выполнения на incomplete.json: %v", config)
	}
	config, err = FromJson("invalid_url.json")
	if err == nil {
		t.Errorf("Неверный результат выполнения на invalid_url.json: %v", config)
	}
	config, err = FromJson("nonexistent.json")
	if err == nil {
		t.Errorf("Неверный результат выполнения на nonexistent.json: %v", config)
	}
}

func TestFromEnvironment(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("JAEGER_URL", "http://jaeger:16686")
	os.Setenv("SENTRY_URL", "http://sentry:9000")
	os.Setenv("KAFKA_BROKER", "kafka:9092")
	config, err := FromEnvironment()
	if err == nil {
		t.Errorf("Неверный результат выполнения на неполных данных: %v", config)
	}

	os.Setenv("SOME_APP_ID", "testid")
	os.Setenv("SOME_APP_KEY", "testkey")
	config, err = FromEnvironment()
	if err != nil {
		t.Errorf("Неверный результат выполнения на полных данных: %v", err)
	}
	if reflect.DeepEqual(valid, config) {
		t.Errorf("Неверный результат выполнения на полных данных: %v", config)
	}
	os.Setenv("DB_URL", "db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	config, err = FromEnvironment()
	if err == nil {
		t.Errorf("Неверный результат выполнения на данных с некорректным db_url: %v", config)
	}
}
