package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	DbUser     string
	DbPassword string
	Dbdriver   string
	DbPort     string
	DbHost     string
	DbName     string
}

func LoadParameters() (Enviroment, error) {
	var err error
	env := Enviroment{}
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
		env.DbUser = os.Getenv("DB_USER")
		env.DbPassword = os.Getenv("DB_PASSWORD")
		env.Dbdriver = os.Getenv("DB_DRIVER")
		env.DbHost = os.Getenv("DB_HOST")
		env.DbPort = os.Getenv("DB_PORT")
		env.DbName = os.Getenv("DB_NAME")
	}
	return env, err
}
