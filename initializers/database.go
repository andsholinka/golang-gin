package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// CoonetToDB is creating a new connection to db
func CoonetToDB() {
	var err error

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed")
	// }

	// dbUser := os.Getenv("MYSQL_USER")
	// dbPass := os.Getenv("MYSQL_PASSWORD")
	// dbHost := os.Getenv("MYSQL_HOST")
	// dbName := os.Getenv("MYSQL_DBNAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dsn := "root:@tcp(127.0.0.1:3306)/golang_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connet to database")
	}
	// return DB
}

//CloseDatabaseConnection method is closing a connection between your app and your db
// func CloseDBConnection(db *gorm.DB) {
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed to close connection from database")
// 	}
// 	dbSQL.Close()
// }
