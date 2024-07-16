package main

import "fmt"

func main() {
	var conferanceName string = "Go Conferance"
	const conferanceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking app\n", conferanceName)
	fmt.Printf("We have only tickets %v and only remain %v Book you ticket don't miss your chance\n", conferanceTickets, remainingTickets)
	fmt.Println("Book the tickets here")

	var userName string
	var userTicket uint

	userName = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)

}
