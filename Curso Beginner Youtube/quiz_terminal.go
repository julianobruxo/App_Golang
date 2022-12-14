package main

import (
	"fmt"
	"os"
	"time"
)

func welcome() {
	fmt.Println("Welcome to my Quiz Game!\n Let's Start!\n What's your first name?")
}
func getUserData() {
	var firstName string
	var lastName string
	var age uint

	fmt.Scan(&firstName)
	time.Sleep(time.Second * 1)
	fmt.Println("Whats your last name?")
	fmt.Scan(&lastName)
	time.Sleep(time.Second * 1)
	fmt.Printf("What's your age, %v ?\n", firstName)
	fmt.Scan(&age)

	if age >= 18 {
		time.Sleep(time.Second * 1)
		fmt.Printf("Ok %v, your age is correct.\nThx for your confirmation.\nLet's start your Quiz :)\n", firstName)
	} else {
		fmt.Printf("Sorry %v, you're a minor and can't participate.", firstName)
		os.Exit(0)
	}
}
func questions() {
	score := 0
	num_of_questions := 5

	fmt.Printf("Which one is better, pizza(1) or sushi(2)?\n")
	var answer int
	fmt.Scan(&answer)

	if answer == 1 {
		fmt.Println("Great choice!")
		score++
	} else {
		fmt.Println("Hmm...I have some doubts about you...")
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("What's the Capital of Canada, Otawa(1) or Toronto(2)?\n")

	fmt.Scan(&answer)

	if answer == 1 {
		fmt.Println("Great choice!")
		score++
	} else {
		fmt.Println("You gotta study more")
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("Which one is bigger, Russia(1) or Canada(2)?\n")

	fmt.Scan(&answer)

	if answer == 1 {
		fmt.Println("Great choice!")
		score++
	} else {
		fmt.Println("You gotta study more")
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("Which one's the Capital of Brazil, São Paulo(1) or Brasilia(2)?\n")

	fmt.Scan(&answer)

	if answer == 2 {
		fmt.Println("Great choice!")
		score++
	} else {
		fmt.Println("You gotta study more")
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("Which one is bigger, Rio(1) or São Paulo(2)?\n")

	fmt.Scan(&answer)

	if answer == 2 {
		fmt.Println("Great choice!")
		score++
	} else {
		fmt.Println("You gotta study more")
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("Your score is %v out of %v.", score, num_of_questions)

	if score >= 3 {
		time.Sleep(time.Second * 1)
		fmt.Println("Congratulations for your score!")
	} else {
		fmt.Println("Sorry, good luck next time.")
	}
}

func main() {
	welcome()
	getUserData()
	questions()

}
