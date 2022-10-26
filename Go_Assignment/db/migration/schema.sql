CREATE TABLE "flights" (
  "id" BIGSERIAL PRIMARY KEY,
  "origin_airport_code" varchar(3) NOT NULL,
  "destination_airport_code" varchar(3) NOT NULL,
  "departure_date_time" timestamp NOT NULL,
  "booking_class" varchar(16) NOT NULL,
  "flight_number" varchar(8) NOT NULL,
  "flight_duration" int NOT NULL,
  "enabled" boolean NOT NULL DEFAULT true,
  "seat_available" int NOT NULL DEFAULT 10,
  "currency_code" varchar(3) NOT NULL,
  "price" float NOT NULL DEFAULT 1000000,
  "create_by" varchar(64),
  "create_date" timestamp DEFAULT (now()),
  "modify_by" varchar(64),
  "modify_date" timestamp
);

CREATE TABLE "booking_transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "reservation_code" varchar(8),
  "status" varchar(32),
  "error_code" int,
  "error_message" varchar(256),
  "create_date" timestamp DEFAULT (now()),
  "modify_date" timestamp
);

CREATE TABLE "booking_flights" (
  "id" BIGSERIAL PRIMARY KEY,
  "booking_transaction_id" bigint,
  "sequence" int NOT NULL,
  "departure_airport_code" varchar(3) NOT NULL,
  "arrival_airport_code" varchar(3) NOT NULL,
  "departure_date_time" timestamp NOT NULL,
  "arrival_date_time" timestamp NOT NULL,
  "booking_class" varchar(16) NOT NULL,
  "currency_code" varchar(3) NOT NULL,
  "price" float NOT NULL
);

CREATE TABLE "booking_passengers" (
  "id" BIGSERIAL PRIMARY KEY,
  "booking_transaction_id" bigint,
  "title" varchar(8),
  "last_name" varchar(25),
  "first_name" varchar(25),
  "ffp_number" varchar(16),
  "ticket_number" varchar(16)
);

CREATE TABLE "users" (
	"id" BIGSERIAL PRIMARY KEY,
	"username" varchar(50),
	"password" varchar(50),
	"first_name" varchar(50),
	"last_name" varchar(50),    
	"address" varchar(200),
	"email" varchar(100)  NOT NULL,
	"phone_number" varchar(12),
	"active" boolean,
	"create_at" timestamp,
	"update_at" timestamp
);

ALTER TABLE "booking_flights" ADD FOREIGN KEY ("booking_transaction_id") REFERENCES "booking_transactions" ("id");

ALTER TABLE "booking_passengers" ADD FOREIGN KEY ("booking_transaction_id") REFERENCES "booking_transactions" ("id");
