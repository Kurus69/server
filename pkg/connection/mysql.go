package connection

import (
	"fmt"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error

	// connection := "root:@tcp(127.0.0.1:3306)/landtick?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	var DB_HOST = "containers-us-west-159.railway.app"
	var DB_USER = "postgres"
	var DB_PASSWORD = "SIKif6a2L2RKw4UWWPHb"
	var DB_NAME = "railway"
	var DB_PORT = "7643"

	fmt.Println(DB_HOST)
	fmt.Println(DB_USER)
	fmt.Println(DB_PASSWORD)
	fmt.Println(DB_PORT)
	fmt.Println(DB_NAME)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected!")
}
