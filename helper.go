package main

import "fmt"

func greeUsers() {
	fmt.Printf("Welcome to the %v bookings app\n", conferanceName)
	fmt.Printf("We have only tickets %v and only remain %v Book you ticket don't miss your chance\n", conferanceTickets, remainingTickets)
	fmt.Println("Book the tickets here")
}
