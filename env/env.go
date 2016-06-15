package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	port        = "PORT"
	portDefault = "80"
)

func init() {
	godotenv.Load()
	if err := initVariables(); err != nil {
		log.Fatal(err.Error())
	}
	Logger = log.New(os.Stdout, "", log.LstdFlags)
}

var (
	//Port is the port used by the application (default is 80)
	Port int
	//Logger is the main stdout for all logs
	Logger *log.Logger
)

//initVariables initalizes all Envrionment Variables for access throughout the application
func initVariables() (err error) {
	tmpPort := os.Getenv(port)
	if tmpPort == "" {
		tmpPort = portDefault
	}
	Port, err = strconv.Atoi(tmpPort)
	if err != nil {
		return NewEnvironmentError(ErrEnvWrongType, port, "int")
	}
	return nil
}
