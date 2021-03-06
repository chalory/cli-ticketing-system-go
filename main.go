package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// for {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		go sendTicket(userTickets, firstName, lastName, email)
		wg.Add(1)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v\n", firstNames)

		if userTickets <= remainingTickets {

			if remainingTickets == 0 {
				// if there are no more tickets, end program
				fmt.Println("Our conference is booked out. Come back next year")
				// break
			}
		}

	} else {

		if !isValidName {
			fmt.Println("first name or last name is too short.\n")

		}

		if !isValidEmail {
			fmt.Printf("The email that you entered was invalid.\n")

		}

		if !isValidTicketNumber {
			fmt.Printf("Number that you entered was invalid.\n")
		}
		fmt.Printf("Your input data is invalid, try again\n")

	}
	wg.Wait()

	//}

}

func greetUsers() {
	fmt.Printf("Welcome to the %v!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and a total of %v is still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create map

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets}

	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email in %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

// simulated email service
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	println("################")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending a ticket: \n %v \nto email address %v\n", ticket, email)
	println("################")
	wg.Done()
}
