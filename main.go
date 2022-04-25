package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and a total of %v is still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend.")

	for {

		var firstName string
		var lastName string
		var email string

		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("slice type: %T\n", bookings)
		fmt.Printf("Length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email in %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are %v\n", firstName)
	}

}
