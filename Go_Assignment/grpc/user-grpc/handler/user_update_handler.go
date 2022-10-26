package user_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	db "khoihm1/flight-booking-assignment/db/sqlc"
	"khoihm1/flight-booking-assignment/pb"
	"khoihm1/flight-booking-assignment/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *UserGrpcHandler) UpdateUser(c context.Context, in *pb.GrpcUpdateUserRequest) (*pb.GrpcUpdateUserResponse, error) {
	//validation data
	if utils.IsEmptyString(in.Email) {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Email is mandantory field"),
		)
	}
	//check existed User
	existedUser, err := h.store.Queries.GetUserInfoByEmail(c, in.GetEmail())
	if nil != err {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Cannot look up existed User| %v", err))
	}

	if !(len(existedUser) > 0) {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("User not found"),
		)
	}

	//Update new User
	existedUserItem := existedUser[0]

	updateUserParam := db.UpdateUserParams{
		ID:          existedUserItem.ID,
		Username:    sql.NullString{String: in.Username, Valid: true},
		FirstName:   sql.NullString{String: in.FirstName, Valid: true},
		LastName:    sql.NullString{String: in.LastName, Valid: true},
		Email:       existedUserItem.Email,
		PhoneNumber: sql.NullString{String: in.PhoneNumber, Valid: true},
		Address:     sql.NullString{String: in.Address, Valid: true},
	}

	response, err := h.store.Queries.UpdateUser(c, updateUserParam)
	fmt.Printf("Updated for User %v", response.Username)

	return &pb.GrpcUpdateUserResponse{
		Status: "OK",
	}, nil
}
