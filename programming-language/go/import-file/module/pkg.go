package module

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnvironment() error {
	envPath := ".env"
	return godotenv.Load(envPath)
}

type ExampleConfig struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     int
}

func ConnDatabase(self ExampleConfig) string {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		self.DatabaseHost,
		self.DatabasePort,
		self.DatabaseUser,
		self.DatabasePassword,
		self.DatabaseName,
	)
	return connStr
}

type ExampleServerConfig struct {
	Host string
	Port int
}

func ExampleServer(self ExampleServerConfig) string {
	connStr := fmt.Sprintf(
		"host=%s port=%d",
		self.Host,
		self.Port,
	)
	return connStr
}
