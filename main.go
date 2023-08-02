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

	greetUsers(conferenceName)

	fmt.Printf("Data Types of conferenceName is: %T , conferenceTickets is: %T and remainingTickets is: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v booking Application\n", conferenceName)
	fmt.Printf("We have total numbers of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket Now")

	for remainingTickets > 0 && len(bookings) < 50 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidaeTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidaeTickets {
			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			// Here in Golang variables are stored in memory and we can access them using their memory address vai adding "&" pointers which acts special var to store memory store-> Hash-table
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
		} else {
			//Enter Valid Details
			if !isValidName {
				fmt.Printf("Your Name is Invalid. Please Try Again\n")
			}
			if !isValidEmail {
				fmt.Printf("Your Email is Invalid. Please Try Again\n")
			}
			if !isValidaeTickets {
				fmt.Printf("Your Tickets are Invalid. Please Try Again\n")
			}
			fmt.Printf("Your Booking is Invalid. Please Try Again\n")
		}
	}
	// Switch Case
	// city := "London"
	// switch city {
	// 	case "London":
	// 		//code for booking London conference tickets
	// 		fmt.Println("Welcome to London")
	// 	case "New York":
	// 		//code for booking New York conference tickets
	// 		fmt.Println("Welcome to New York")
	// 	case "Sydney", "Melbourne":
	// 		//code for booking Sydney conference tickets
	// 		fmt.Println("Welcome to Sydney or Melbourne")
	// 	case "Paris":
	// 		//code for booking Paris conference tickets
	// 		fmt.Println("Welcome to Paris")
	// 	case "Tokyo" , "Singapore":
	// 		//code for booking Tokyo conference tickets
	// 		fmt.Println("Welcome to Tokyo or Singapore ")
	// 	default:
	// 		//code for No vaild Input display
	// 		fmt.Println("Welcome to Default")
	// }

}

func greetUsers(confeName string) {
	fmt.Printf("Welcome to our %v Booking Application\n", confeName)
}
