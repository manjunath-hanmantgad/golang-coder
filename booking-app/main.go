package main

import "fmt"

func main() {
	//fmt.Print("helo world!")
	//fmt.Println("Welcome to our coneference on building booking application.")
	//fmt.Println("Get your tickets here to attend.")
	// variables and constants
	var ConferenceName = "GoConference"
	const conferenceTickets = 50 // this is a constant , because it will not change throughtout our application.
	//fmt.Println(ConferenceName)
	// tracker for remainingtickets
	var remainingTickets = 50

	// using variable names in the function.

	fmt.Println("Welcome to", ConferenceName, "on building booking application.")
	fmt.Println("We have a total of", remainingTickets, "tickets remaining.")
	fmt.Println("Get your tickets now as there are only", conferenceTickets, "remaining")
	//fmt.Printf("Welcome to %v on building booking application.\n", ConferenceName)
	// fmt.Printf("We have a total of %v tickets remaining.\n", remainingTickets)
	// fmt.Printf("Get your tickets now as there are only %v remaining.\n", conferenceTickets)

	// DATA TYPES
	// idhar isko btana padta kya hai type string ya int ya array etc
	//var UserName string
	var FirstName string
	var LastName string
	var email string
	var UserTickets int

	//UserName = "Tomya"
	// what if we want to ask user his name ?
	// use .Scan() function
	// *************************************************

	//fmt.Println("enter your name here:")
	//fmt.Scan(&UserName) // &mm --> & matlab pointer.

	fmt.Println("enter your first name here:")
	fmt.Scan(&FirstName) // &mm --> & matlab pointer.

	fmt.Println("enter your last name here:")
	fmt.Scan(&LastName) // &mm --> & matlab pointer

	fmt.Println("enter your email here:")
	fmt.Scan(&email) // &mm --> & matlab pointer

	fmt.Println("enter how many tickets you want here:")
	fmt.Scan(&UserTickets) // &mm --> & matlab pointer

	//UserTickets = 11
	// fname, lname, email , number of tickets
	fmt.Printf("User %v %v booked %v tickets!. You will receive your tickets on the email %v given by you.\n", FirstName, LastName, UserTickets, email)

	// jana hai kya type hai var and const ka ?
	//fmt.Printf("ConferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", ConferenceName, conferenceTickets, remainingTickets)
}
