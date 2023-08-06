package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidaeTickets := userTickets > 0 && userTickets <= remainingTickets

	// In golang u can return multiple values from function
	return isValidName, isValidEmail, isValidaeTickets
}