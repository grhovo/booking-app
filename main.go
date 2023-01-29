package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []UserData
var maxTickets uint
var lastKnownWinner = ""

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	printWelcome()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidUserTickets {
			bookTicket(userTickets,firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets,firstName, lastName, email)
			
			maxTickets,lastKnownWinner = getMostTicketOwner(userTickets, maxTickets, firstName, lastKnownWinner)
			printInfo(remainingTickets)
			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				// end program
				wg.Wait()
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

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}