package main

import (
	"log"
	"os"
	"time"

	"import-file/module"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	envPath := ".env"
	return godotenv.Load(envPath)
}

func main() {
	err := LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := module.ExampleConfig{
		DatabaseUser:     os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
		DatabaseHost:     os.Getenv("POSTGRES_HOST"),
		DatabasePort:     60010,
	}
	connStr := module.ConnDatabase(dbConfig)
	log.Println(connStr)

	testServerConfig := module.ExampleServerConfig{
		Host: "localhost",
		Port: 8080,
	}
	testServerConnStr := module.ExampleServer(testServerConfig)
	log.Println(testServerConnStr)

	// Write the server and database connection strings to a file
	os.Create("example_server_connection.txt")
	os.WriteFile(
		"example_server_connection.txt", []byte(
			time.Now().String()+"\n"+testServerConnStr+"\n"+connStr), 0644,
	)
}
