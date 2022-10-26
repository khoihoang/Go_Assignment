package user_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	db "khoihm1/flight-booking-assignment/db/sqlc"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) CreateUser(c context.Context, in *pb.GrpcCreateUserRequest) (*pb.GrpcCreateUserResponse, error) {
	//validation data
	fmt.Println("====================0000===")
	if utils.IsEmptyString(in.Email) {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Username is mandantory field"),
		)
	}
	fmt.Println("====================111===")
	//check existed User
	existedUser, err := h.store.Queries.GetUserInfoByEmail(c, in.GetEmail())
	if nil != err {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Cannot look up existed User| %v", err))
	}
	fmt.Println("====================222===")
	if len(existedUser) > 0 {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Username is existed"),
		)
	}
	fmt.Println("====================333===")
	//Insert new User
	isActive := true
	createdAt := time.Now()
	createUserParam := db.CreateUserParams{
		Username:    sql.NullString{String: in.Username, Valid: true},
		Password:    sql.NullString{String: in.Password, Valid: true},
		FirstName:   sql.NullString{String: in.FirstName, Valid: true},
		LastName:    sql.NullString{String: in.LastName, Valid: true},
		Email:       in.Email,
		PhoneNumber: sql.NullString{String: in.PhoneNumber, Valid: true},
		Address:     sql.NullString{String: in.Address, Valid: true},
		Active:      sql.NullBool{Bool: isActive, Valid: true},
		CreateAt:    sql.NullTime{Time: createdAt, Valid: true},
	}
	fmt.Println("====================444===")
	response, err := h.store.Queries.CreateUser(c, createUserParam)
	fmt.Println("====================555===")
	return &pb.GrpcCreateUserResponse{
		Status:   "OK",
		Email:    response.Email,
		Active:   response.Active.Bool,
		CreateAt: timestamppb.New(createdAt),
	}, nil
}
