package main

import (
	"github.com/joelseq/surreal-search/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load env vars from .env, fail silently in prod
	godotenv.Load(".env")

	cmd.Execute()
}
