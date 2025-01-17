package main

import "fmt"

func main() {

    // conferenceName := "Go Conference" // Shorthand declaration
	const conferenceName = "Go Conference"
    const conferenceTickets uint = 50
	var remainingTickts uint = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickts)
	fmt.Println("Get your tickets here to attend")

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
	remainingTickts = remainingTickts - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at\n %v", firstName, lastName,userTickets, email)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickts)

}
