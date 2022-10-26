package main

import (
	"fmt"
	flight_api_handler "khoihm1/flight-booking-assignment/api/flight/handlers"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadApiConfig("./api/flight/")
	if err != nil {
		fmt.Println("Cannot load config")
		log.Fatal("Cannot load config with path:", err)
	}
	grpcClient, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Grpc.FlightGrpc.Host, config.Grpc.FlightGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcFlightClient := pb.NewFlightGrpcClient(grpcClient)
	flightApiHandler := flight_api_handler.InitFlightHandler(grpcFlightClient)

	engine := gin.Default()
	routerGroup := engine.Group("/api/flight")
	routerGroup.POST("/create", flightApiHandler.CreateFlight)
	routerGroup.PUT("/update", flightApiHandler.UpdateFlight)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
