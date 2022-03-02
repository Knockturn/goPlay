package main

// Youtube Part Start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=10080s

/*
AGENDA (Control Flow):

    If Statement
        Operators
        if -else and if-else if statements
    Switch Statement
        Simple cases
        Cases with multiple tests
        Falling through
        Type Switches
*/

import (
	"fmt"
	"math/rand"
	"time"
)
 

func main() {
    if true { // simplest example
        fmt.Println("If and Switch Statements")
    }

    statePopulations := map[string]int{
        "California": 39250017,
        "Texas" : 27862596,
        "Florida": 20612439,
    }

 

    /*
    pop, ok := statePopulations["Texas"] is the initializer part.
    Below the pop then takes out in a variable the stuff that comes after. , ok .. and until ; is the taking out the Texas and putting "value"/isExisting in the ok - the part after ; is the logical test 

    Alternative:
    if pop, ok := statePopulations["Texas"]; ok {
        fmt.Println(pop)
    }
    */

    pop, ok := statePopulations["Texas"]
    if ok {
        fmt.Println(pop)
    }

    // Operators
    operators()

     /*
    if
    else if
    else

    first one will be used so even if if and the else if would be true - the if is always taken and the rest skipped.
    */
 
    switchStatements(1)
}

 

// NOTE: Be careful when evaluating floats - as they are aproximations of the value (around here: https://youtu.be/YS4e4q9oBaU?t=11064)

func operators(){
    rand.Seed(time.Now().Unix()) // Seed will determine what number is generated. Time is used to truly create a random pattern.
    // Setting the seed above w. a fixed number is great if you want to predictably test a random number.

    number := rand.Intn(10)
    var guess int  

    for guess != number {
        fmt.Print("\nGuess a number: ")
        fmt.Scan(&guess)    

        // the or test below: if the first part is true (short circuiting) then it will skip the other or. Note: you can have several or
        if guess < 1 || guess > 10{ // oppisite could be && operator, ex: guess > 1 && guess < 10. Remember also the ! operator which reverces the bool
            fmt.Println("Number is from 1-10")
        }

        if guess < number {
            fmt.Println("Too low")      
        }

        if guess > number {
            fmt.Println("Too high")
        }

        if guess == number{
            fmt.Println("CORRECT!")
        }
    }

    /*
    <=
    >=
    !=
    */
}

 

func returnTrue() bool {
    fmt.Println("Returning True")
    return true
}

// SWITCH
func switchStatements(test int) {

    fmt.Println("\nSWITCH STATEMENTS:")

    switch test {
    case 1:
        fmt.Println("one")
        // break -> the break is implied - so not needed
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("not one or two")
    }

    switch test {
    case 1, 5, 10:
        fmt.Println("one, five, ten")
    case 2, 4, 6:
        fmt.Println("two, four, six")
    default:
        fmt.Println("not one or two")
    }

     // tagless syntax. more powerful - but more verbose
    i := 10
    switch {
    case i <= 10:
        fmt.Println("Less than or equal to 10")
        fallthrough // this enables the switch to go to the next case. It is logicless. So even IF the 2nd case is NOT true -> it WILL execute
    case i <= 20:
        fmt.Println("Less than or equal to 20")
    // case i == "test": // invalid - has to be the same type
    default:
        fmt.Println("greater than twentry")
    }

    // with interface and pulling underlying type
    var iFace interface{} = 1
    switch iFace.(type) {
    case int:
        fmt.Println("iFace is an int")
    case float64:
        fmt.Println("iFace is a float64")
    case string:
        fmt.Println("iFace is a string")
    default:
        fmt.Println("iFace is not: int, float64 or string")
    }

    // How to leave a switch early: (use break keyword)
    var iBreak interface{} = 1.0
    switch iBreak.(type) {
    case int:
        fmt.Println("iFace is an int")
        // break
        fmt.Println("This will not print if break above is uncommented")
    case float64:
        fmt.Println("iFace is a float64")
    case string:
        fmt.Println("iFace is a string")
    default:
        fmt.Println("iFace is not: int, float64 or string")
    }
}