package main

import (
	"flag"
	"fmt"
	"geekbrains-go/lvl-1/config"
	"os"
)

/*func main() {
	fmt.Println("Введите номер числа Фибоначчи, которое хотите получить")
	n := scan.Int64(os.Stdout, os.Stdin)
	fmt.Println("Loop result: ", fibonacci.Loop(n))
	//fmt.Println("Recursive result: ", Fibonacci.Recursive(n))
	fmt.Println("RecursiveWithMap result: ", fibonacci.RecursiveWithMap(n))
}*/

var (
	filename   = flag.String("filename", "config"+string(os.PathSeparator)+"valid.json", "Путь до конфигурационного файла")
	configType = flag.String("config_type", "json", "Тип конфигурационного файла. Допустимые значения: json, yaml или environment")
)

func main() {
	flag.Parse()
	var err error
	var conf *config.Config
	switch *configType {
	case "environment":
		os.Setenv("PORT", "8080")
		os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
		os.Setenv("JAEGER_URL", "http://jaeger:16686")
		os.Setenv("SENTRY_URL", "http://sentry:9000")
		os.Setenv("KAFKA_BROKER", "kafka:9092")
		os.Setenv("SOME_APP_ID", "testid")
		os.Setenv("SOME_APP_KEY", "testkey")
		conf, err = config.FromEnvironment()
	case "json":
		conf, err = config.FromJson(*filename)
	case "yaml":
		conf, err = config.FromYaml(*filename)
	default:
		fmt.Printf("Неизвестный тип конфигурационного файла %s\n", *configType)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Не удалось прочитать конфигурацию: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Загружена конфигурация:\n%+v\n", *conf)
}
