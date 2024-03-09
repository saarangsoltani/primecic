package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gocandoit.xyz/cli"
	"gocandoit.xyz/fillers"
	api "gocandoit.xyz/http"
	"gocandoit.xyz/util"
)

func main() {
	readEnv()
	util.ClearScreen()

	applicant, _ := api.GetApplicant("2")
	choice := cli.Start()
	if choice == "IMM5257" {
		fillers.FillForm(applicant)
	} else {
		fmt.Println("Comming soon!")
	}
}

func readEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
