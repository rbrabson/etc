package main

import (
	"github.com/joho/godotenv"
	"github.com/rbrabson/ftc/ftc"
	log "github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

// main is the entry point for the FTC OpenSkill calculations
func main() {
	log.SetLevel(log.DebugLevel)

	/*
		general, err := ftc.GetGeneral()
		if err != nil {
			log.Error(err)
		}
		fmt.Println("Body", general)
	*/

	_, err := ftc.GetTeams("2023")
	if err != nil {
		log.Error(err)
	}
	// fmt.Println("Body", leagues)

	/* TODO: uncomment this out when done
	// Wait for the user to interrupt the server and then exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Info("Press Ctrl+C to exit")
	<-sc
	*/
}
