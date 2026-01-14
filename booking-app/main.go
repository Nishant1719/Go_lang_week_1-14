package main

import (
	"fmt"
	"strings"
)

// func buy(total int, userInput int) {
// 	total = total - userInput
// 	fmt.Println("In func total")
// }

func main() {

	// fmt.Print("Hello, World!")
	// fmt.Println("Hello, State!") // Always and endline after printing
	// fmt.Println("Hello, City!") // Meaning this will print in a new line

	// fmt.Println("Welcome to our conf. booking app")
	// fmt.Println("Get your tickets")

	// variables: var,const,
	var conferenceName = "Go conference" // var can be updated
	// const conferenceName = "Const Go conference" // constants cannot be updated
	conferenceName = "New Go conference" // Auto imply the datatype by the value assigned if the value is defined.
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Println(conferenceName)
	fmt.Println("Printing dynamic string", conferenceName, "like this maybe")
	// fmt.Println("We have", conferenceTickets, "we have left with", remainingTickets)
	fmt.Printf("We have %v we have left with, %v\n", conferenceTickets, remainingTickets)
	fmt.Println("New line here!")
	// buy(conferenceTickets, 10)

	// Data Types: User defined values
	// Placeholders : %v - values , %T - Dataype of variables

	// This is how you define : Scope Name Datatype
	var userName string // Where values are not assigned we need to define a type explicitly.
	// Ask user  for their name
	var userTickets uint

	// userName = "Tom"
	// userTickets = 2

	// Pointers to update the values stored in the Variables
	println("Enter the name")
	fmt.Scan(&userName)

	println("Enter the tickets")
	fmt.Scan(&userTickets)
	// Validate the inputs
	if userTickets > remainingTickets {
		fmt.Println("Invalid Input")
	}
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("User %v booked this many tickets %v\n", userName, userTickets)
	fmt.Printf("Type of conferencename : %T, and Type of conferenceTickets %T\n", conferenceName, conferenceTickets)
	fmt.Println("Remaining Tickets", remainingTickets)
	// uint make its a positive integer
	// var num uint = -10 // this gives errors
	// var num uint = 10
	// var num1 float32 = 12.123
	// fmt.Println(num)
	// fmt.Println(num1)

	// Go specific variable declaration:
	// exampleNum := "Example String"
	// Not for contants
	// exampleNum = "Updated String"
	// print(exampleNum)
	// ----------------------------------------------------------------------
	// Array And slices
	// var bookings [50]string // Declaration
	// bookings = append(bookings, "Nishant")  // Wrong
	// bookings[0] = userName
	// bookings[1] = "waghade"
	// fmt.Println("Booking done by all", bookings)
	// fmt.Println("Booking done by", bookings[0])

	// Problem with arrays : if the coder doesnt know the size of an array.
	// Slices comes in the game for dynamic in size
	// Slices
	var bookings []string
	bookings = append(bookings, "Nishant")
	bookings = append(bookings, "Waghade")
	bookings = append(bookings, userName)
	// fmt.Printf("Booking type %T and Booking : %v\n", bookings, bookings)
	//Alternative syntax for creating slices
	// bookingt := []int{}
	// var bookingt = []int{}
	// print(bookingt)
	// To print everything just print the array or the slice
	// fmt.Printf("Printing everything %v", bookings)

	// Loops
	// for {
	// For infinie loops - press ctrl + c
	// 	fmt.Println("Enter again")
	// 	fmt.Scan(&userName)
	// 	bookings = append(bookings, userName)
	// 	fmt.Println(bookings)
	// }

	// for index, ele := range bookings {
	// 	fmt.Println(index)
	// 	fmt.Println(ele)
	// }

	for _, ele := range bookings { // underscore is a ignored variable
		preNames := "Sir"
		bookings = append(bookings, ele+" "+preNames)
		// fmt.Println(bookings)
		// Output
		// [Nishant Waghade nish Nishant Sir]
		// [Nishant Waghade nish Nishant Sir Waghade Sir]
		// [Nishant Waghade nish Nishant Sir Waghade Sir nish Sir]
	}
	fmt.Println(bookings)
	onlyNames := []string{}
	for _, ele := range bookings {
		// Expected output remove Sir from everything as just 3 elements at the end have combined names with space.
		removeNames := strings.Fields(ele)
		onlyNames = append(onlyNames, removeNames[0])
	}

	fmt.Printf("No sir in the list %v\n", onlyNames)

	// Conditionals
	if len(onlyNames) > 10 {
		fmt.Println("Yes it is greater than 10")
	} else if len(onlyNames) < 10 {
		fmt.Println("Yes it is lower than 10")
	} else {
		fmt.Println("Invalid")
	}
	addCheezWithNames(bookings)
}

func addCheezWithNames(bookingSlice []string) {
	bookingSlice = append(bookingSlice, "Cheeze")
	fmt.Println(bookingSlice)
}
