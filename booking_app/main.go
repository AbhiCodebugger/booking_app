package main

import (
	"fmt"
	"strings"
)

func main() {

	// conferenceName := "Go Conference" // Shorthand declaration
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickts uint = 50
	// var bookings = [50]string{} // Array of strings
	var bookings []string // Slice of strings
	// bookings := []string{} // shorthand declaration for Slice of strings

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickts)
	fmt.Println("Get your tickets here to attend")

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidUserTicket :=  userTickets > 0 && userTickets <= remainingTickts

		if isValidName && isValidEmail && isValidUserTicket {
			remainingTickts = remainingTickts - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickts)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First name of the user who have booked tickets: %v\n", firstNames)

			if remainingTickts == 0 {
				fmt.Println("All tickets are booked. Thank you for booking.")
				// end the program
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Please enter valid first name or last name.")
			}
			if !isValidEmail {
				fmt.Printf("Sorry! We have only %v tickets available.\n", remainingTickts)
			}
			if !isValidUserTicket{
				fmt.Println("Please enter valid number of tickets.")
			}

		}
	}
}
