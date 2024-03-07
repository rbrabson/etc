package ftc

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	server string
)

// init retrieves the username and authentication key and creates a base64-encoded token for authentication
func init() {
	godotenv.Load()

	server = os.Getenv("SERVER")
}
