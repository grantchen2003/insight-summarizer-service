package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grantchen2003/insight/summarizer/internal/config"
	databasePackage "github.com/grantchen2003/insight/summarizer/internal/database"
	serverPackage "github.com/grantchen2003/insight/summarizer/internal/server"
)

func main() {
	env := os.Getenv("ENV")
	log.Printf("ENV=%s", env)
	if err := config.LoadEnvVars(env); err != nil {
		log.Fatalf("failed to load env vars")
	}

	database := databasePackage.GetSingletonInstance()
	database.Connect()
	defer database.Close()

	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	server := serverPackage.NewServer()
	if err := server.Start(address); err != nil {
		log.Fatalf("failed to start server")
	}
}
