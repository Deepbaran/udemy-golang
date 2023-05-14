package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		// fmt.Println("Error: ", err)
		// os.Exit(1)
		panic(err)
	}

	/*
		//resp is a structure.
		//resp.Body is an interface of type io.ReaderCloser
		//io.ReaderCloser has two interface types inside it, Reader and Closer.
		//Reader has a Read function declaration inside it.
		// fmt.Println(resp)
	*/

	/*
		//make function will take a type of a slice and initialize the slice with 99999 elements/empty spaces
		//In this way we are creating a slice with 99999 elements inside of it
		//The below line gives us a byte slice that has space for 99999 elements
		bs := make([]byte, 99999)

		//The reason why we created a byte slice with 99999 spaces at the start is because, the Read function is not designed to resize the slice if the slice is full.
		//Meaning if the Response body requires more space than the byte slice initially had, then the Read function will be unable to add more space to it.
		//So, the Read function will read data from Response body and store int he byte slice, till the byte slice is full.
		//So, if we passed an empty byte slice to Read function, then Read function will not store anything in the byte slice as it will appear as already full.
		//So, we are assuming that 99999 space should be enough to hold all the data that the Response Body has.

		resp.Body.Read(bs)      //The Response struct (resp) has a Body attribute that satisfies the ReadClose interface, which in turn satisifes the Reader interface, meaning it must have a Read function.
		fmt.Println(string(bs)) //We can think of byte slice as a string. This prints the whole source of "http://google.com" homepage.
	*/

	/*
		io.Copy(os.Stdout, resp.Body) //This prints the whole source of "http://google.com" homepage.
		//Stdout is of type File that implements the Writer interface.
	*/

	lw := logWriter{}
	io.Copy(lw, resp.Body) //This will work even with wrong Writer logic. As lw still passes as Writer interface type
}

func (logWriter) Write(bs []byte) (int, error) {
	//With this Write function being assocaiated with logWriter, now logWriter is implementing the Writer interface
	//Although logWriter is now implementing the Writer interface, we can write the code in such a way that the Writer interface is not meant to do.

	//Example of wrong logic
	// return 1, nil

	//Correct logic
	fmt.Println(string(bs))
	//Custom implementation of the Writer interface
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil //The integer should be the number if bytes written in the byte slice that was processed [0 <= n <= len(bs)]
}
