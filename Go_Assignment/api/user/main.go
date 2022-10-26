package main

import (
	"fmt"
	user_api_handlers "khoihm1/flight-booking-assignment/api/user/handlers"
	"khoihm1/flight-booking-assignment/common"
	"khoihm1/flight-booking-assignment/pb"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	apiConfig, err := common.LoadApiConfig("./api/user/")
	if err != nil {
		log.Fatal("Can not load api config with path:", err)
	}
	userGrpcClientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", apiConfig.Grpc.UserGrpc.Host, apiConfig.Grpc.UserGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userGrpcClient := pb.NewUserGrpcClient(userGrpcClientConn)
	userApiHandler := user_api_handlers.InitUserApiHandler(userGrpcClient)

	engine := gin.Default()
	routerUserGroup := engine.Group("/api/user")

	routerUserGroup.POST("/create", userApiHandler.CreateUser)
	routerUserGroup.PUT("/update", userApiHandler.UpdateUser)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", apiConfig.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
