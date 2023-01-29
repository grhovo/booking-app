package main

import (
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []map[string]string
var maxTickets uint
var lastKnownWinner = ""



func main() {

	printWelcome()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidUserTickets {
			bookTicket(userTickets,firstName, lastName, email )
			
			maxTickets,lastKnownWinner = getMostTicketOwner(userTickets, maxTickets, firstName, lastKnownWinner)

			printInfo(remainingTickets)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				// end program
				printSummary(lastKnownWinner,maxTickets)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("lastName or firstName is incorrect")
			}
			if !isValidEmail {
				fmt.Println("Your Email is incorrect")
			}
			if !isValidUserTickets {
				fmt.Println("You've entered wrong number of UserTickets")
			}
			fmt.Println("Your input is invalid, please try again")
		}
	}

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
}

func printInfo(remainingTickets uint) {
	fmt.Printf("Here is the list of people who have acquired ticket - %v\n", getFirstNames())
	fmt.Printf("Remains - %v\n", remainingTickets)
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func printWelcome() {
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v\n", conferenceTickets)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getMostTicketOwner(userTickets uint, maxTickets uint, firstName string, lastKnownWinner string) (uint, string) {
	if userTickets > maxTickets {
		lastKnownWinner = firstName
		return userTickets, lastKnownWinner
	} else {
		return maxTickets, lastKnownWinner
	}
}

func printSummary(lastKnownWinner string, maxTickets uint)  {
	fmt.Println("We have sold out all tickets for conference. Here is the list of participants")
	for _, booking := range bookings {
		fmt.Println(booking["firstName"] + " " + booking["lastName"])
	}

	fmt.Printf("Most number of tickets is acquired by %v and count is %v\n", lastKnownWinner, maxTickets)
}