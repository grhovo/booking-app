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
