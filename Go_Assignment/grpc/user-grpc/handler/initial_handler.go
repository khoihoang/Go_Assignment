package user_grpc_handler

import (
	db "khoihm1/flight-booking-assignment/db/sqlc"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"
)

type UserGrpcHandler struct {
	pb.UnimplementedUserGrpcServer
	config utils.GrpcConfig
	store  *db.Store
}

func InitUserGrpcHandler(config utils.GrpcConfig, store *db.Store) (*UserGrpcHandler, error) {
	return &UserGrpcHandler{
		config: config,
		store:  store,
	}, nil
}
