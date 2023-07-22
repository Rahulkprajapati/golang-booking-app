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

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Ask User Input
	fmt.Println("Please Enter Your First Name:  ")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter Your Last Name:  ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter Your Email:  ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets You Want to Book:  ")
	fmt.Scan(&userTickets)

	// Here in Golang variables are stored in memory and we can access them using their memory address vai adding "&" pointers which acts special var to store memory store-> Hashtable

	fmt.Printf("Thank you %v %v for booking %v tickets for %v conference. Your tickets are booked and you will receive a confirmation email on %v\n", firstName, lastName, userTickets, conferenceName, email)
}
