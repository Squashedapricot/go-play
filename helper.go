package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
