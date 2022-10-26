package flight_api_handler

import (
	"fmt"
	api_flight_payload "khoihm1/flight-booking-assignment/api/flight/payload"
	core_api "khoihm1/flight-booking-assignment/core"
	"khoihm1/flight-booking-assignment/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	GrpcClient pb.FlightGrpcClient
}

func InitFlightHandler(grpcClient pb.FlightGrpcClient) FlightHandler {
	return FlightHandler{
		GrpcClient: grpcClient,
	}
}

func (f *FlightHandler) CreateFlight(c *gin.Context) {
	request := CreateFlightRequest{}

	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}

	grpcCreateFlightRequest := &pb.GrpcCreateFlightRequest{
		OriginAirportCode:      request.OriginAirportCode,
		DestinationAirportCode: request.DestinationAirportCode,
		DepartureDateTime:      timestamppb.New(time.Time(request.DepartureDateTime)),
		BookingClass:           request.BookingClass,
		FlightNumber:           request.FlightNumber,
		FlightDuration:         request.FlightDuration,
		SeatRemaining:          request.SeatRemaining,
		CurrencyCode:           request.CurrencyCode,
		Price:                  request.Price,
	}

	grpcCreateFlightResponse, err := f.GrpcClient.CreateFlight(c, grpcCreateFlightRequest)
	if err != nil {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}
	response := api_flight_payload.CreateFlightResponse{
		Id:       grpcCreateFlightResponse.Id,
		CreateBy: grpcCreateFlightResponse.CreateBy,
		CreateAt: grpcCreateFlightResponse.CreateAt.AsTime(),
	}

	fmt.Println("Create flight API - DONE")
	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}
