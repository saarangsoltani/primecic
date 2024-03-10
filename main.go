// hello to you
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"gocandoit.xyz/cli"
	"gocandoit.xyz/fillers"
	api "gocandoit.xyz/http"
	"gocandoit.xyz/util"
)

func main() {
	readEnv()
	util.ClearScreen()

	applicantID := promptForApplicantID()
	applicant, _ := api.GetApplicant(applicantID)
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

func promptForApplicantID() string {
	fmt.Println("Welcome to the PRIMECIC CLI. This command-line application allows you complete CIC immigration forms in a quick and efficient manner.")

	inputValidated := false
	var id string
	fmt.Print("\nTo start, enter the applicant ID: ")
	retryMessage := "Invalid input. Please enter a valid applicant ID: "
	for !inputValidated {
		_, err := fmt.Scanf("%s", &id)
		if err != nil {
			fmt.Print(retryMessage)
		} else {
			_, err := strconv.Atoi(id)
			if err != nil {
				fmt.Print(retryMessage)
			} else {
				inputValidated = true
			}
		}
	}
	return id
}
