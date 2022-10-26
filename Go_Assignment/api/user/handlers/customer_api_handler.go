package user_api_handler

import (
	"fmt"
	user_api_payload "khoihm1/flight-booking-assignment/api/user/payload"
	core_api "khoihm1/flight-booking-assignment/core"
	"khoihm1/flight-booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApiHandler struct {
	GrpcClient pb.UserGrpcClient
}

func InitUserApiHandler(grpcClient pb.UserGrpcClient) UserApiHandler {
	return UserApiHandler{
		GrpcClient: grpcClient,
	}
}

// new User
func (h *UserApiHandler) CreateUser(c *gin.Context) {
	request := user_api_payload.CreateUserRequest{}
	fmt.Println(" ====== call api create user", request)
	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}
	grpcCreateUserReq := &pb.GrpcCreateUserRequest{
		Username:    request.Username,
		Password:    request.Password,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	grpcCreateUserRes, err := h.GrpcClient.CreateUser(c, grpcCreateUserReq)
	if err != nil {

		c.JSON(http.StatusOK, core_api.CreateUserApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}

	response := user_api_payload.CreateUserResponse{
		Status:    grpcCreateUserRes.Status,
		Username:  grpcCreateUserRes.Email,
		CreatedAt: grpcCreateUserRes.CreateAt.String(),
	}
	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}

//Update user
func (h *UserApiHandler) UpdateUser(c *gin.Context) {
	request := user_api_payload.UpdateUserRequest{}
	fmt.Println(" ====== call api update user", request)
	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}

	grpcUpdateUserRequest := &pb.GrpcUpdateUserRequest{
		Username:    request.Username,
		Password:    "",
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	grpcUpdateUserRes, err := h.GrpcClient.UpdateUser(c, grpcUpdateUserRequest)

	if err != nil {
		c.JSON(http.StatusOK, core_api.CreateUserApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}
	fmt.Println(" ====== call api update user 222", request)
	response := user_api_payload.UpdateUserResponse{
		Status: grpcUpdateUserRes.Status,
	}

	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}
