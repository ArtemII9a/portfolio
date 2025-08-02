package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ArtemII9a/portfolio/src/api"
	"github.com/joho/godotenv"
)

var (
	flagDebug = flag.Bool("debug", false, "used to print API debug info")
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Could not load the .env file.")
	}
	apiVersion := os.Getenv("API_VERSION")
	port := os.Getenv("API_PORT")

	apiServer := api.NewAPIServer(apiVersion, port)
	err = apiServer.Run()
	if err != nil {
		panic(fmt.Sprintf("Got an error running the server: %v", err))
	}
}
