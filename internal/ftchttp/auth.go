package ftchttp

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	username string
	authKey  string
)

// init retrieves the username and authentication key and creates a base64-encoded token for authentication
func init() {
	godotenv.Load()

	username = os.Getenv("USERNAME")
	authKey = os.Getenv("AUTHORIZATION_KEY")
}
