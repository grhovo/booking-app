package main
import "fmt"

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
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
