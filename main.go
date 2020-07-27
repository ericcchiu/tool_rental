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
	var toolDataStore tools.ToolDataStore
	// ------- Load environmental variables-----------
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	// ------- Credentials ---------------------------
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	// fmt.Printf("WHAT IS THIS user: %s \n", host)
	// fmt.Printf("WHAT IS THIS port: %s \n", port)
	// fmt.Printf("WHAT IS THIS user: %s \n", user)
	// fmt.Printf("WHAT IS THIS password: %s \n", password)
	// fmt.Printf("WHAT IS THIS dbname: %s \n", dbname)
	// fmt.Printf("WHAT IS THIS sslmode: %s \n", sslmode)

	// ------ Create Connection -----------------------
	fmt.Printf("Establishing connection")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	db, err := psql.NewPostgresConnection(connection)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	toolDataStore = psql.NewPostgresToolDataStore(db)

	toolService := tools.NewToolService(toolDataStore)

	/*
		newTool := tools.Tool{
			Name:        "Shelf",
			Description: "Oak hardwood shelf",
			Price:       300,
			Quantity:    20,
		}

		err = toolService.CreateTool(&newTool)
		if err != nil {
			fmt.Println("Error while creating:", err)
		}
	*/
	/*
		foundTool, err := toolService.FindToolByID("fec9e1dc-02d2-4382-ad4b-e4ba07a53eab")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(foundTool)
	*/
	/*
		allTools, err := toolService.FindAllTools()
		if err != nil {
			fmt.Println(err)
		}
		for _, tool := range allTools {
			fmt.Println(tool)
		}
	*/

	updateTool := tools.Tool{
		ID:          "fec9e1dc-02d2-4382-ad4b-e4ba07a53eab",
		Name:        "Hamertech Hammer",
		Description: "Leading Hammertech recursive hammer technology that strikes itself",
		Price:       2000,
		Quantity:    2,
	}

	err = toolService.Update(&updateTool)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(updateTool)

	defer db.Close()

}
