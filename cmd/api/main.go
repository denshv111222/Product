package main

import (
	"flag"
	"log"

	"GitHab/Standart_Server_API/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {

	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("It works")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериалзиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
		log.Println("Э")
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
