package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v. You booked %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year. ")

			}
		} else {

			if !isValidName {
				fmt.Println("The first or last name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("The email address you entered does not contain the '@' sign.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}

			{
				fmt.Println("Your data input is invalid, try again.")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Printf("We have a total of %v tickets and there are %v still available!\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}
