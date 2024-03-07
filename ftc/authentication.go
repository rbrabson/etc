package ftc

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	username string
	authKey  string
	server   string
)

// init retrieves the username and authentication key and creates a base64-encoded token for authentication
func init() {
	godotenv.Load()

	username = os.Getenv("USERNAME")
	authKey = os.Getenv("AUTHORIZATION_KEY")
	server = os.Getenv("SERVER")
}
