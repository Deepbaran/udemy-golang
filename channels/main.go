package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//Create a channel through which we want to communicate values of type string
	c := make(chan string) //make creates a value out of the given type (chan string <- A channel type whuich communicates only with string)

	//c is treated a normal value as any other.
	//So, if we want to access the channel inside any go routine, we need to pass the channel inside the routine as an argument

	for _, link := range links {
		// checkLink(link) //running it in the Main Routine
		go checkLink(link, c) //go keyword makes the function to run in a new Go routine
	}

	//Receive value from the channel
	//fmt.Println(<-c) //This is a Blocking call for the Main routine

	//As soon as the Main routine will get some value from the channel,
	//possibly from the go routine that is executing the google.com (first one to be fully executed by resolving the request),
	//Main routine will resume execution and exit the program without waiting for other routines to get executed.
	//So, only one go routine got executed and others did not get the chance to get executed.

	//As soon as a message is sent to the channel inside the child routine, the go runtime will search for any routine that is witing for a message from this channel
	//And the routines that are waiting will receive the message and resume execution.

	//Now if we were to send value to the channel before execution of the rest of the code in the child go routine,
	//then as soon as a value was put inside the channel, Main routine would have received it and resume execution and exit the program
	//without even waiting for the child routine that assigned value to the channel to complete execution of the rest of the code.
	//So, always put sending values to channels inside child go routines either at the end of the function or at the end of the error condition.

	//To prevent the Main routine from continuing and waiting for other child routine's execution, we can put more blocking statements.

	// fmt.Println(<-c) //Blocking code for google.com
	// fmt.Println(<-c) //Blocking code for facebook.com
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c) //This will make the Main routine stay blocked as there are no routines left for the channel to recive message from.

	// for i := 0; i < len(links); i++ {
	// 	//The for loop will be blocked till we receive message from the channel before continuing with the iteration
	// 	fmt.Println(<-c) //This line is essentially handling the flow whenever a request is getting compelted
	// }

	//Repeating Routines
	//Infinite for loop
	// for {
	// 	//This is a blocking statement. We will wait till we get a value (link) in the channel and immediately call the checkLink function in a new go routine
	// 	go checkLink(<-c, c) //The for loop will pause till we receive next message from the channel
	// }

	//Alternative syntax of the loop to make it more readable
	// for l := range c {
	// 	//Just like when we write a for loop for a slice (i, v := range slice), here it works similarly.
	// 	//What range and for loop does is, it takes out all the values in the slice, assigns them to variables and iterates over them.
	// 	//Here, we are saying, wait for the channel to return some value.
	// 	//After the channel returns some value, assign it to l and run the body of the for loop
	// 	go checkLink(l, c)
	// }

	//Function Literal -> Function literals are like lambda expression in Python or Anonymous functions in JavaScript
	//It is used to wrap some code for us to execute in the future.

	for l := range c {
		// fl := func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(l, c)
		// }
		// go fl()

		// go func() {
		// 	time.Sleep(5 * time.Second) //This will make sure that we only pause before the followup checkLink function execution and the Main routine stays unpaused
		// 	checkLink(l, c)
		// }() //The parenthesis at the end is to execute the function literal and as we are using go keyword, it will be executed inside a new go coroutine.

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)

	//As this is a Blocking call, once one Go routine hits it, and waits for response, it will let all the routines (inc. main routine) that is blocked here for sometime
	//Because of this, other routines will use the processor resources and start executing.
	//Main routine will hear about the blocking and continue with the for loop.
	//A new go routine will be created and it will start executing checkLink with
	_, err := http.Get(link) //Blocking call!
	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "Might be down I think" //send a message into the channel

		//pausing the coroutine for 5 seconds
		// time.Sleep(5 * time.Second)

		c <- link
		fmt.Println()
		return
	}

	fmt.Println(link, " is up!")
	//c <- "Yep its up" //send a message into the channel

	//pausing the coroutine for 5 seconds
	// time.Sleep(5 * time.Second)

	c <- link
}
