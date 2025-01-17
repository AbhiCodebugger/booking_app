package main

import (
	"booking_app/booking_app/helper"
	"fmt"
	"strconv"
)

// conferenceName := "Go Conference" // Shorthand declaration
const conferenceName = "Go Conference"
const conferenceTickets uint = 50

var remainingTickets uint = 50

// var bookings = [50]string{} // Array of strings
// var bookings []string // Slice of strings
// bookings := []string{} // shorthand declaration for Slice of strings
var bookings = make([]map[string]string, 0) // List of maps

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicket {

			bookTickets(userTickets, firstName, lastName, email)

			// Print first name of the user who have booked tickets
			firstNames := getFirstNames()
			fmt.Printf("Users who have booked tickets: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are booked. Thank you for booking.")
				// end the program
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Please enter valid first name or last name.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address.")
			}
			if !isValidUserTicket {
				fmt.Println("Please enter valid number of tickets.")
			}

		}
	}

}
func greetUser() {

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {		
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to book:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickets)

}
