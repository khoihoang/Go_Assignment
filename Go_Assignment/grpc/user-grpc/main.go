package main

import (
	"database/sql"
	"fmt"
	db "khoihm1/flight-booking-assignment/db/sqlc"
	customer_grpc_handler "khoihm1/flight-booking-assignment/grpc/user-grpc/handler"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Start customer-service-grpc")
	config, err := utils.LoadGrpcConfig("./grpc/user-grpc/")

	if nil != err {
		log.Fatal("Cannot load config ", err)
		return
	}

	//init database connection
	conn, err := sql.Open(config.Database.DriverName, utils.GetDatabaseSourceNameForGrpc(config.Database))
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}
	store := db.InitalStore(conn)

	//init GRPC service
	grpcServer := grpc.NewServer()
	grpcHandler, err := customer_grpc_handler.InitUserGrpcHandler(config, store)
	if err != nil {
		log.Fatal("Can not initial grpc service:", err)
	}

	reflection.Register(grpcServer)
	pb.RegisterUserGrpcServer(grpcServer, grpcHandler)

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		panic(err)
	}
	log.Printf("Listening at port: %v\n", config.Server.Port)

	grpcServer.Serve(listen)
}
