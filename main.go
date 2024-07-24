package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference" // the same as var confN string = "GC"
	const conferenceTickets uint = 50 // setting a type is protecting against assigning wrong data type
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	// fmt.Printf("conferenceTickets is %T, ramainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // needs to be a pointer!

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// quit app
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}
}
