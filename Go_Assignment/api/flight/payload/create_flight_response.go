package api_flight_payload

import "time"

type CreateFlightResponse struct {
	Id       int64     `json:"id"`
	CreateBy string    `json:"createBy"`
	CreateAt time.Time `json:"createAt"`
}
