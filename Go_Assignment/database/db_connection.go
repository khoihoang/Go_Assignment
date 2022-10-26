package database

import (
	"database/sql"
	"fmt"

	//_constant "khoihm1/flight-booking-assignment/constant"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection(host string, port int, user string, password string, dbName string) *sql.DB {

	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	dbConn, err := sql.Open(string("_constant.POSTGRESQL"), connection)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return dbConn
}
