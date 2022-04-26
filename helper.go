package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isvalidEmail, isValidTicketNumber
}
