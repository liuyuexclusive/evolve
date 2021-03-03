package main

import (
	messageHandler "github.com/liuyuexclusive/evolve/srv/basic/handler/message"
	roleHandler "github.com/liuyuexclusive/evolve/srv/basic/handler/role"
	userHandler "github.com/liuyuexclusive/evolve/srv/basic/handler/user"

	message "github.com/liuyuexclusive/evolve/srv/basic/proto/message"
	role "github.com/liuyuexclusive/evolve/srv/basic/proto/role"
	user "github.com/liuyuexclusive/evolve/srv/basic/proto/user"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("basic"),
		service.Version("latest"),
	)

	s := srv.Server()
	// Register handler
	role.RegisterRoleHandler(s, new(roleHandler.Handler))
	user.RegisterUserHandler(s, new(userHandler.Handler))
	message.RegisterMessageHandler(s, new(messageHandler.Handler))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
