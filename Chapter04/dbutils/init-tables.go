package dbutils

import "log"
import "database/sql"

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}
	statement, _ = dbDriver.Prepare(station)
	statement.Exec()
	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()
	log.Println("All tables created/initialized successfully!")
}
