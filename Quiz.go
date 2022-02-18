package main

// TODO: Dictionary of questions to loop and then count the answers. Take all input from scan
import (
	"fmt"
	"strings"
)

func main()  {
	var playerName string
	
	fmt.Println("Welcome to the goQuiz!")

	fmt.Printf("Please enter your name: ")
	fmt.Scan(&playerName) //& addresses the ram location of what follows

	fmt.Printf("\nThanks, %v - welcome to the game!\n", playerName)

	fmt.Printf("Please enter age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("\nPerfect, you are old enough to play!")
	} else {
		fmt.Println("\nSorry, you need to be above 10 to play...")
		return // Will exit program in this case
	}

	//Score Tracker
	score := 0
	maxScore := 2

	fmt.Printf("Who killed Qui-Gon Jinn, Darth Maul or Darth Sidius? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2) //Scan seperates input by space
	/* 
	There is a better way to get the whole line (not shown in tutorial)
	Also - if 2 var is provided it will expect 2 input and will wait
	for 2nd if only 1 is provided
	*/
	if strings.ToUpper(answer + " " + answer2) == "DARTH MAUL" {
		fmt.Println("Correct!")
		score ++ // same as score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("\nHow many suns does Tattooine have? ")
	var suns uint
	fmt.Scan(&suns)
	if suns == 2 {
		fmt.Println("Correct!")
		score ++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("\nYou scored %v out of %v", score, maxScore)
	// two int cannot become float - so you need to convert
	percent := (float64(score) / float64(maxScore)) * 100
	fmt.Printf("\nThat equals to: %v%%", percent) //double %% as escape
	fmt.Print("\n\nThanks for playing!")
}