package main

// YouTube Part: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=13294s
/*
AGENDA: (control flow)
	Defer (function that can be invoked but delay its execution)
	Panic (enter a state where application can no longer run)
	Recover (if application can be saved, how we tell that to the entire application)
*/

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	fmt.Println("Defer, Panic & Recover:")
	// Defer is executed: Last in first out. Defer is often used to close out resources.
	// Resource closing: Let's say you open a DB connection and will use it throughout your function. if you put defer db.close() at the next line - you have remembered to close it - but it will run at the end so the rest of function wont fail.
	fmt.Println("\nstart")
	defer fmt.Println("middle") // Defer executes AFTER the main function is done but BEFORE it returns any results to the calling function. Not at the end of the function as you would think.
	defer fmt.Println("end")

	// defer is taken from the time it is called: see below tha ta is not replaced in the print even though it is executed after.
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// panicFunc()
	// panicHttp()
	recoverFunc()
}

func panicFunc() {
	// Similar to an exception - but different in go. Opening a file that is not there will give an error is logical and therefore not exceptional behavior. 
	// If the code gets into a situation where it doesnt know what to do it panics because it is exceptional behavior.
	fmt.Println("\nPanic")

	// Below will break code and produce a similar result to a Panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func panicHttp() {
	/*
	It is up to the developer to determine when an error is critical: see below. If listenandserve fails all fails therefor the panic is there.
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
// panic is called last normal defer called before you panic
func recoverFunc(){
	fmt.Println("Start")
	panicker() // Will exit the function after the panic but continue the code below - like a try-catch (panic-recover)
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() { // anonymous function. Defined at one point and can only be called 1 time
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err) // rethrow in order to get the FULL stack (if you need to debug)
		}
	}()
	panic("something went wrong...")
	fmt.Println("done panicking")
}