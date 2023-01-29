package main
import "fmt"

func printInfo(remainingTickets uint) {
	fmt.Printf("Here is the list of people who have acquired ticket - %v\n", getFirstNames())
	fmt.Printf("Remains - %v\n", remainingTickets)
}

func printWelcome() {
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v\n", conferenceTickets)
}

func printSummary(lastKnownWinner string, maxTickets uint)  {
	fmt.Println("We have sold out all tickets for conference. Here is the list of participants")
	for _, booking := range bookings {
		fmt.Println(booking.firstName + " " + booking.lastName)
	}

	fmt.Printf("Most number of tickets is acquired by %v and count is %v\n", lastKnownWinner, maxTickets)
}