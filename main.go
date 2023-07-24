package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Using Slices instead of Arrays gives us more flexibility and we can add more items in it without defining the size of array
	// var bookings []string
	bookings := []string{}

	fmt.Printf("Data Types of conferenceName is: %T , conferenceTickets is: %T and remainingTickets is: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v booking Application\n", conferenceName)
	fmt.Printf("We have total numbers of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket Now")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Ask User Input
		fmt.Println("Please Enter Your First Name:  ")
		fmt.Scan(&firstName)

		fmt.Println("Please Enter Your Last Name:  ")
		fmt.Scan(&lastName)

		fmt.Println("Please Enter Your Email:  ")
		fmt.Scan(&email)

		fmt.Println("Enter Number of Tickets You Want to Book:  ")
		fmt.Scan(&userTickets)
        

		if userTickets > remainingTickets {
			fmt.Printf("Sorry! We only have %v tickets left, You cant book %v, tickets\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Whole Booking Array is: %v\n", bookings)
		fmt.Printf("First Booking is: %v\n", bookings[0])
		fmt.Printf("Slice type is: %T\n", bookings)
		fmt.Printf("Slice Length is: %v\n", len(bookings))

		// Here in Golang variables are stored in memory and we can access them using their memory address vai adding "&" pointers which acts special var to store memory store-> Hashtable

		fmt.Printf("Thank you %v %v for booking %v tickets for %v conference. Your tickets are booked and you will receive a confirmation email on %v\n", firstName, lastName, userTickets, conferenceName, email)
		fmt.Printf("Remaining Tickets are: %v for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		// Foreach loop which gives index and its corresponding value in each iteration
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The First Names of Bookings are: %v\n", firstNames)


		if remainingTickets == 0 {
			fmt.Println("Sorry! We are Sold Out")
			break
		}

		// If we want to break the loop we can use break keyword
	}
}
