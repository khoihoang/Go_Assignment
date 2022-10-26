package flight_grpc_handler

import (
	db "khoihm1/flight-booking-assignment/db/sqlc"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"
)

type FlightGrpcHandler struct {
	pb.UnimplementedFlightGrpcServer
	config utils.GrpcConfig
	store  *db.Store
}

func InitFlightGrpcHandler(config utils.GrpcConfig, store *db.Store) (*FlightGrpcHandler, error) {
	return &FlightGrpcHandler{
		config: config,
		store:  store,
	}, nil
}
