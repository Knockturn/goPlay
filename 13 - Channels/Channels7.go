package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
// Select statements
// Start: https://youtu.be/YS4e4q9oBaU?t=23238

/*
	Showcase for when you don't have an obvious way to close your channel.
	This is a basic logger

	A select statement allow goroutines to monitor several channels at once.
*/
const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Struct w. no fields require 0 memory allocation. This is a signal only channel.
// Note to above. You could be tempted to use a bool: var doneCh = make(chan bool). But that actually requires memory to be allocated. So it's more optimal (however a minor one) to use an empty struct

func main() {
	go logger() // Will monitor the logCh for entries coming into that channel (forever basically). In this example it will theoretically stay open until a signal is recieved in the doneCh
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond) // Gives it enough time to process the goroutine. Because it will forcebly close too soon without.
	// Note to above: The app closes as soon as the last statement of the main function has finished executing.
	doneCh <- struct{}{} // Without this it will not be a graceful shutdown. And sometimes this can be okay. But it is best practice to always close it gracefully. Worst case without it it can become a subtle memory leak.
	// Explanation to the initialization of the empty struct above: https://youtu.be/YS4e4q9oBaU?t=23584
}

func logger() {
	for {
		select { // A select with out a default is called a blocking select statement. https://youtu.be/YS4e4q9oBaU?t=23548
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh: // Will loop forever until it recieves an signal from the doneCh. Normally it would get sent in as a parameter in the logger function above.
			break
		// A default here would break it because it would then not be a blocking select statement. So if no entries come in for the above cases it will execute the default.
		}
	}
}