package main

import (
	"fmt"
)

func main() {
	var conferanceName string = "Go Conferance"
	const conferanceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v bookings app\n", conferanceName)
	fmt.Printf("We have only tickets %v and only remain %v Book you ticket don't miss your chance\n", conferanceTickets, remainingTickets)
	fmt.Println("Book the tickets here")

	// bookings slice var
	var bookings []string

	for {

		//variables
		var firstName string
		var lastName string
		var userTicket uint
		var email string

		// asking for input
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Print("Enter you email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want to buy:")
		fmt.Scan(&userTicket)

		// book ticket in system
		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for bookings the %v tickets for %v, You will recieve a email at %v", firstName, lastName, userTicket, conferanceName, email)
		fmt.Printf("RemainingTickets %v \n", remainingTickets)

		fmt.Printf("The Whole slice %v", bookings)
		fmt.Printf("Length of slice %v\n", len(bookings))
	}

}
