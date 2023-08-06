package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// We defining our of main func as well as global variables here but we have to expllicitly define the data type of variables
var conferenceName string = "Golang Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

//Here we are creating a empty list of maps needs to provide size of list intially -> make([]map[string]string, 0)
// var bookings = make([]map[string]string, 0)

var bookings = make([]UserData, 0)

// Mixed Data Types kinda of class in OOP
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for remainingTickets > 0 && len(bookings) < 50 {

		// Validate User Input
		firstName, lastName, email, userTickets := getUserInput()

		//helper -> helper.go -> ValidateUserInput
		isValidName, isValidEmail, isValidaeTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidaeTickets {
			bookTicket(userTickets, firstName, lastName, email)

			// Send Email "go" keyword is used to run the function in background and it will not wait for the function to complete
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			//first Name
			firstNames := getfirstNames()
			fmt.Printf("The First Names of Bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry! We are Sold Out, come back next year ✈️")
				// break
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
		wg.Wait()
	// }
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

func greetUsers() {
	fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
	fmt.Printf("Data Types of conferenceName is: %T , conferenceTickets is: %T and remainingTickets is: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("We have total numbers of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket Now")
}

func getfirstNames() []string {
	// when we returning a value from function needs to also define the return type of function []string
	firstNames := []string{}
	//
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Please Enter Your First Name:  ")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter Your Last Name:  ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter Your Email:  ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets You Want to Book:  ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// creating empty maps for userData -> map[key]value
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings: %v\n", bookings)

	// Here in Golang variables are stored in memory and we can access them using their memory address vai adding "&" pointers which acts special var to store memory store-> Hash-table
	fmt.Printf("Thank you %v %v for booking %v tickets for %v conference. Your tickets are booked and you will receive a confirmation email on %v\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("Remaining Tickets are: %v for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("**********Sending Email**********")
	fmt.Printf("Your Tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("**********Email Sent**********")
	wg.Done()
}
