package main

import "fmt"

func main() {
	conferenceName := "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Data Types of conferenceName is: %T , conferenceTickets is: %T and remainingTickets is: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v booking Application\n", conferenceName)
	fmt.Printf("We have total numbers of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket Now")

	var userName string
	var userTickets int

	// Ask User Input
	fmt.Scan(&userName)
	// Here in Golang variables are stored in memory and we can access them using their address vai  kind of table

	userTickets = 5
	fmt.Printf("Hello %v, you have booked %v tickets\n", userName, userTickets)
}
