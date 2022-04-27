package main

import (
	"fmt"
	"sync"
	"time"
)


const conferenceTickets uint = 50
const conferenceName string = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName,lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//call function getFirstNames
		firstNames := getFirstNames()
		fmt.Printf("Here are the first names of the people who have booked: %v\n", 		firstNames)

		if remainingTickets == 0 {
			//end the program
			fmt.Println("Sorry, the conference is booked out. Come back next year!")
			// break
		}
	} else {
		if !isValidName {
				fmt.Println("First or last name is invalid.")
		}
		if !isValidEmail {
					fmt.Println("Invalid email.")
		} 
		if !isValidTicketNumber {
		fmt.Println("Invalid number of tickets.")
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend...")
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

	//ask user for their first name, last name, and email
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	//ask user how many tickets they want
	fmt.Println("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				numberOfTickets: userTickets,
			}

			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, 	conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	//simulating sending ticket process
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v to eamil address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()

}

