package main

import (
	"fmt"
	"strings"
)

// global vars
var conferanceName string = "Go Conferance"

const conferanceTickets int = 50

var remainingTickets uint = 50

func main() {

	greeUsers()
	// bookings slice var
	var bookings []string

	for {

		// asking for input
		firstName, lastName, email, userTicket := takeUserInfo()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for bookings the %v tickets for %v, You will recieve a email at %v", firstName, lastName, userTicket, conferanceName, email)
			fmt.Printf("RemainingTickets %v \n", remainingTickets)

			var firstNames = getsFirstNames(bookings)
			fmt.Printf("First Names of bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Unfortunately Our conferance is completely booked, Hope you see you at the next conferance")
				break
			}
		} else {
			fmt.Printf("Only tickets avaliable are %v please try again\n", remainingTickets)
			continue
		}
	}

}

func getsFirstNames(bookings []string) []string {
	var firstNames []string
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
func takeUserInfo() (string, string, string, uint) {

	//variables
	var firstName string
	var lastName string
	var userTicket uint
	var email string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Print("Enter you email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to buy:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}
