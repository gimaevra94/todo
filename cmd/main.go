package main

import (
	"log"

	todo "github.com/gimaevra94/todo"
	handler "github.com/gimaevra94/todo/pkg/handler"
	"github.com/gimaevra94/todo/pkg/repository"
	"github.com/gimaevra94/todo/pkg/service"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
