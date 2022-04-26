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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isvalidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isvalidEmail && isValidTicketNumber {

			if userTickets <= remainingTickets {
				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email in %v.\n", firstName, lastName, userTickets, email)
				fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
				firstNames := []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					var firstName = names[0]
					firstNames = append(firstNames, firstName)
				}
				fmt.Printf("The first names of bookings are %v\n", firstNames)

				if remainingTickets == 0 {
					// if there are no more tickets, end program
					fmt.Println("Our conference is booked out. Come back next year")
					break
				}
			}

		} else {

			if !isValidName {
				fmt.Println("first name or last name is too short.\n")

			}

			if !isvalidEmail {
				fmt.Printf("The email that you entered was invalid.\n")

			}

			if !isValidTicketNumber {
				fmt.Printf("Number that you entered was invalid.\n")
			}
			fmt.Printf("Your input data is invalid, try again\n")

		}
	}

}
