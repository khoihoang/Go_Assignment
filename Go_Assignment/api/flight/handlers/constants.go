package flight_api_handler

import core_api "khoihm1/flight-booking-assignment/core"

var (
	BAD_REQUEST_CODE      = core_api.InitApiError(400, "Bad request")
	INTERNAL_SERVER_ERROR = core_api.InitApiError(500, "Server error")
)
