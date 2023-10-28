package main

import (
	"errors"
	"gateway/pkg/client"
	"gateway/pkg/common"
	"gateway/pkg/delivery"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Println("error loading env", err)
	}

	server, err := InitializeApi(config)
	if err != nil {
		log.Println("error configuring server", err)
	}

	server.Run(config.Port)
}

func InitializeApi(cfg common.Config) (server *gin.Engine, err error) {
	server = gin.Default()
	client, err := client.InitUserClient(cfg)
	if err != nil {
		return nil, errors.New("Initializing client error" + err.Error())
	}
	handlers := delivery.NewUserHandlers(client)

	routes := delivery.NewUserRoutes(handlers)

	routes.StartUserRoutes(server)

	return server, nil
}
