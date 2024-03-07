package ftchttp

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	username string
	authKey  string
)

// init retrieves the username and authentication key used for authentication on HTTP requests sent
// to the FTC server API endpoint.
func init() {
	godotenv.Load()

	username = os.Getenv("USERNAME")
	authKey = os.Getenv("AUTHORIZATION_KEY")
}
