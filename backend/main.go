package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	//First of all, remember to clear tmp files before running the app. If not, configuring line 17 to the path is necessary
	err := godotenv.Load()
	if err != nil {
		slog.Info("NO .env file found, relying on system env variable.")
	}

	server := app.New()
	defer server.CloseDB()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	server.Run(port)
}
