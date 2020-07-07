package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ericcchiu/tool_rental/db/psql"
	"github.com/ericcchiu/tool_rental/tools"
	"github.com/joho/godotenv"
)

func main() {
	// Load environmental variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var toolDataStore tools.ToolDataStore
	// ------- Credentials -----------
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	fmt.Printf("WHAT IS THIS user: %s \n", user)
	// fmt.Printf(os.Getenv("DB_HOST"))

	// ------ Create Connection ------
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	db, err := psql.NewPostgresConnection(connection)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// toolDataStore = "firefly" //Data store constructor implemented here

	// toolService := tools.NewToolService(toolDataStore)

	fmt.Println(toolDataStore)
}
