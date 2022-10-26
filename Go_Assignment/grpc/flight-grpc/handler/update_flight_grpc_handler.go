package flight_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	db "khoihm1/flight-booking-assignment/db/sqlc"
	"khoihm1/flight-booking-assignment/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightGrpcHandler) UpdateFlight(ctx context.Context, in *pb.GrpcUpdateFlightRequest) (*pb.GrpcUpdateFlightResponse, error) {
	updateBy := "khoihm1"
	updateAt := time.Now()
	updateFlightParams := db.UpdateFlightParams{
		ID:                     in.Id,
		OriginAirportCode:      in.OriginAirportCode,
		DestinationAirportCode: in.DestinationAirportCode,
		DepartureDateTime:      time.Time(in.DepartureDateTime.AsTime()),
		BookingClass:           in.BookingClass,
		FlightNumber:           in.FlightNumber,
		FlightDuration:         in.FlightDuration,
		SeatAvailable:          in.SeatRemaining,
		CurrencyCode:           in.CurrencyCode,
		Price:                  in.Price,
		Enabled:                in.Enabled,
		ModifyBy:               sql.NullString{String: updateBy, Valid: true},
		ModifyDate:             sql.NullTime{Time: updateAt, Valid: true},
	}
	flight, err := h.store.Queries.UpdateFlight(ctx, updateFlightParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Could not update flight with id=%v|%v", in.Id, err))
	}

	out := &pb.GrpcUpdateFlightResponse{
		Id:       flight.ID,
		UpdateBy: updateBy,
		UpdateAt: timestamppb.New(updateAt),
	}

	return out, nil
}
