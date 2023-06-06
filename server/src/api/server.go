package api

import (
	"log"
	"pruebaTecnica/src/api/controllers"
	"pruebaTecnica/src/config"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var server = controllers.Server{}

	env, err := config.LoadParameters()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		server.Initialize(env.DbUser, env.DbPassword, env.DbPort, env.DbHost, env.DbName)
	}

	server.Run(":8080")

}
