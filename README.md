# udemy-golang

My practice code with Udemy Golang course <https://www.udemy.com/go-the-complete-developers-guide/>

---

# Notes

## FAQ

- Golang is a Compiled language.
- Golang is a statically typed language.
- Go is not a Object Oriented Programming Language.
- To run it on our local computer we need the Go Runtime which allows us to build and execute Go programs.
- Everything (almost) in go is a Type.
- Any variable that is declared in the code must be used. If we do not want to use it, just use '_' as the variable name.
- := can only be used when declaring and intializing a varaible for the first time and not for reassigning.
- There are not try-catch or Error Handling in Golang as of yet.

## Questions

### 1. helloworld
- How do we run the go code in our project?
  - The 'go' command gives us the ability to compile and execute our projects.
  - Command to write the project: 
    - > go run main.go"
    - run multiple files:
      - > go run main.go filename.go
  - Go CLI tools
    ```
    go build -> Compiles a bunch of go source code files (Creates an executable file)
    go run -> Compiles and executes one or two files
    go fmt -> Formats all the code in each file in the current directory
    go install -> Compile and "installs" a package
    go get -> Downloads the raw source code of someone else's package
    go test -> Runs any tests associated with the current project
    ```
- What does 'package main' mean?
  - A package in go is like a project or a workspace.
  - A package is a collection of common source code files.
  - If we were working on one single application, we would traditionally be working on one single package.
  - A package can have many related files inside of it. Each file ending with the extension of .go
  - The only requirement of every file inside a package, is that the first line should declare the package where the file belongs to.
  - So, each file belonging to package main, must have the "package main" written at the top.
  - Now the reason why we called the package as "main" inside helloworld folder, instead of using the folder name, is because:
    - Inside go, there are two different types of packages.
      - Executable -> Generates a file that we can run
      - Reusable -> Code used as 'helpers'. Good place to put reusable logic.
    - Name of the package determines if we are creating an Executable or Dependency type package.
    - So, Specifically the word "main" is used to create executable type package.
    - Any other package name other than "main" will create reusable package.
    - "go build" only creates executable file from Executable package. If we ran it for reusable package, nothing will happen.
    - Reusable packages are a good place to define dependencies (helper code)
    - Anyime we create an Executable package, it MUST have a func called "main".
- What does 'import "fmt"' mean?
  - This line simply means that give my current code, all of the code and functionality contained within the package called "fmt".
  - "fmt" is a standard library package that is included with go by default.
  - "fmt" is kind of a short for "format".
  - The "fmt" library is used to print out information to the terminal.
  - Some of the default standard library packages that come with go by default are:
    - fmt
    - debug
    - math
    - encoding
    - crypto
    - io
    - bufio
    - os
  - We can use the "import" statement to get fuctionalities from not only standard packages but from packages that are built by others as well.
  - In the below link we can get all the standard packages.
    - https://pkg.go.dev/std
  - So, to summarize, we use "import" to gain the functionalities defined in another package inside the package that we are authoring.
- What's that 'func' thing?
  - It is a keyword to declare a function.
- How is the main.go file organized?
  - At the very top, we always do package declaration
    > package main
  - Then we import other packages that we need (standard or other user created)
    > import "fmt"
  - Then at the end we Declare functions, tell Go to do things
    > func main() {
      > fmt.Println("hi there")
    > }
---
### 2. cards
- Functionalities in the Cards game:
  - newDeck -> Create a list of playing cards. Essentially an array of strings
  - print -> Log out the contents of a deck of cards.
  - shuffle -> Shuffles all the cards in a deck.
  - deal -> Create a 'hand' of cards.
  - saveToFile -> Save a list of cards to a file on the local machine.
  - newDeckFromFile -> Load a list of cards from the local machine.
- Variable Declaration in go:
  > var card string = "Ace of Spades"
  - var -> keyword that defines that we are about to create a new variable.
  - card -> The name of the variable.
  - string -> Type of the varaible / Type of value that this variable can store.
  -  "Ace of Spades" -> Value that is stored in the variable.
  -  If we have an unused varaible in our code, lexer will be throwing errors.
     -  Lexer is the one that checks the grammar of the code.
-  Basic Data Types in Go:
   -  bool -> true | false
   -  string -> "Hi!"
   -  int -> 0 | -10000 | 99999
   -  float64 -> 10.000001 | 0.00009 | -100.003
-  Other ways to declare variables in Go:
    - 2nd way:
      > var card = "Ace of Spades"
      - Here the go compiler will infer the type of variable form the value that is getting assigned to it.
      - In this case it is a string
    - 3rd way:
      > card := "Ace of Spades"
      - We can use the walrus operator.
      - Here also the go compiler helps to infer the type of the variable.
      - Walrus operator cannot be used to declare and intialize a variable in global scope.
    - One thing to remember, these two type of declaration does not make our variable dynamic.
    - We can declare and initialize a new variable in this way. But reassigning is not possible. We will need to use '='.
- If we only declare a variable, we can initialize it in a different line. But for the time, the variable is not initalized, go does not store garbage value in them, rather stores some proper values. Like 0 or int type, empty string for string type, etc.
- We can declare variables in global scope and later define them in a function.
- Functions in Go:
  > func newCard() string {}
  - func -> Keyword for function
  - newCard -> Name of the function
  - string -> return type of the function (If there are no return type, the function is not expected to return anything).
  - We can use functions, declared in a file, in a different file inside the same package, without importing the declaration file.
  - Return multiple values from a function in Go:
  ```
  func deal(d deck, handSize int) (deck, deck) { 
    //return multiple values
    //Different types can be returned together
    return d[:handSize], d[handSize:] 
  }
  ```
- Syntax to call and recive a function that returns multiple values:
  ```
  hand, remainingDeck := deal(cards, 5)
  //hand, remainingDeck := <deck>, <deck>
  //the first value returned will be stored in hand and the 2nd value returned will be stored in remainingDeck.
  ```
- Data Structure in Go to handle list of records:
  - There are two data structures in Go to handle them:
    - Array -> fixed length list of things
      - > cards := [4]string{}
      - > cards := [4]string{"1","2","3","4"}
      - One thing to note here is that even if no element is inserted in the array, this will reserve 4 spaces for the array and we can access them without them having any values assigned. This is not the case with slice.
    - Slice -> An array that can grow or shrink with more functionality
      - > cards := []string{}
      - > cards := []string{"1","2","3","4"}
  - Slices and arrays both must be defined with a data type. So, every element in array and slice must be of same type.
  - Append elements to a slice:
    - > cards = append(cards, "Six of Spades")
    - the append function takes all the elements from cards slice, appends "Six of Spades" and then returens a modified slice. 
    - So, append does not modify the real slice, it just returns an updated slice, which we then store in the original slice.
  - Iterating through slices:
    - > for i, card := range cards {fmt.Println(i, card)}
    - i -> index of the element inside the slice
    - card -> Current card we're iterating over
    - range cards -> Take the slice of 'cards' and loop over it
    - We are using the := operator to declare and intialize i and card, as we are throwing out i and card after each iteration. So, it is not reassigning, rather it is a completely new declaration and initialization.
    - Now, as we know each variable tat is declared in our code must be used, otherwise code will throw error, so if we do not wish to use the index of the slices while iterating over it, we can use '_' instead of 'i'.
      - > for _, card := range cards {fmt.Println(card)}
  - Also, we can only iterate over the elements over the index in a slice rather than the elements itself:
    - > for i := range cards {{fmt.Println(i)}}
  - More on Slices:
    - Slices are 0-indexed
    - range in slice:
      - > cards[startIndexIncluding : upToNotIncluding]
      - > cards[ : upToNotIncluding] // [0 : upToNotIncluding]
      - > cards[startIndexIncluding : ] // [startIndexIncluding : len(cards)]
      - This much like list slicing in python.
      - Slicing slices does not modify the real slice.
    - Length of a slice:
      - > len(cards)
- Custom Types in Go:****
  - We can "extend" a base type (string, integer, float, array, map) and add some extra functionality to it.
  - Tell Go we want to create an array of strings and add a bunch of functions specifically made to work with it:
    - > type deck []string{}
      - deck will extend every type of behavior of a slice of string or a string slice.
      - Think of deck as a thin layer or a kind of abstraction, where deck is essectially a string but now we can have particular functions tied to something specifically of type deck.
    - Functions, with 'deck' as a 'receiver' -> A function with a receiver is like a "method" - a function that belongs to an "instance". So, these custom functions will only work with this 'deck' type or are bound to variables of type "deck".
      - We should not make all functions bound to our custom type if they are working on them.
      - Create method (functions with receivers) for our custom type, only if they are acting on the varaible and modifying it.
    - Now in the codebase we can replace all slice of string or string slice with deck to create new instances of deck type if we wished to.
    - > cards := deck{"Ace of Diamonds", newCard()}
    - Syntax to add functionality/method/Receiver Function with the receiver deck:
      ```
      func (d deck) print() {
        for i, card := range d {
          fmt.Println(i, card)
        }
      }
      ```
        - d -> The actual copy of the deck we're working with is available in the function as a variable called 'd'. d refers to the actual value / actual instance and not a copy of it. [think of d as 'this' equivalent in oops]
        - Every variable of type 'deck' can call this function on itself
          - > cards.print()
        - By convention we name the receiver by 1 or 2 letters that matches the type of the receiver. Hence, d in the case of type deck. Go does not use words like "this" and "self".
      - Here deck is the receiver to the function print. And it can be defined as d inside the function.
      - As we can see inside the print function, range d, that d refers to the varaible of type "deck" through which this function is called.
      - So, receivers sets up methods on varaibles that we create. Basically associates the methods to varaibles of type "Deck"
- Write to a File in golang: https://pkg.go.dev/os#WriteFile
  > func WriteFile(name string, data []byte, perm FileMode) error
  ```
  package main

  import (
    "log"
    "os"
  )

  func main() {
    //0666 permission means that anyone can read or write in the file
    err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
    if err != nil {
      log.Fatal(err)
    }
  }
  ```
  - This os.WriteFile(), requires a byte slice as an argument. 
    - Think of of a string of characters evertime you encounter a byte slice
    - Each character in the string can be converted to their ascii value and get stored inside a byte slice. (https://asciitable.com)
      - "Hi there!" (string) -> [72 105 32 116 104 101 114 101 33] (byte slice)
    - Type Conversion from string to byte slice in Go:
      - > []byte("Hi there!")//TypeWeWant(ValueWeHave)
- Read from a File in golang: https://pkg.go.dev/os#ReadFile
  > func ReadFile(name string) ([]byte, error)
  ```
  package main

  import (
    "log"
    "os"
  )

  func main() {
    data, err := os.ReadFile("testdata/hello")
    if err != nil {
      log.Fatal(err)
    }
    os.Stdout.Write(data)

  }
  ```
  - err -> Value of type 'error'. If nothing went wrong, it will have a value 'nil' (null of golang).
- As Golang does not have proper error handling as of yet, we can do any one of the below things in case an error occurs.
  - Log the error and return a default value
  - Log the error and entirely quit the program
    > panic(err)
- Random in Go: https://pkg.go.dev/math/rand
  - > func Intn(n int) int
  - This randome number generator is a pseudo-random generator that depends on some seed value.
  - Think of the seed as some source of randomness inside the number generator.
  - We take the seed and pass it to the random generator and the generator gived random numbers or values.
  - The issue here is that the go random generator by default always uses the exact same seed.
  - So, every time we restart our program and run the random generator, same sequence of random values will get generated as we are using the same exact seed value.
  - In order to fix this, we need to generate some random seed value and then feed it to the random number generator.
  ```
  source := rand.NewSource(time.Now().UnixNano()) //time.Now().UnixNano() this seed value makes sure that the seed passed to the source for rand stays random
	r := rand.New(source)                           //Creating a new rand with our custom seed
  newPosition := r.Intn(len(d) - 1) //Our randomly seeded rand is used here
  ```
  - But as of Go 1.20, the math/rand package automatically seeds the global random number generator.
- Easy way to swap elements in Go:
  - > d[i], d[newPosition] = d[newPosition], d[i]
- Writing **Tests** in Go:
  - Go testing is not RSpec, mocha, jasmine, selenium, etc!
  - To make a test, create a new file ending in _test.go
    - deck_test.go
  - To run all tests in a package, run the command:
    - > go test
  - We can write Test functions for each individual functions inside a file or we can write a test function for all the closely related functions in the file.
  - Run all tests in go:
    > go test
  - Unlike other languages, Go does not know how many test cases were run inside a function. All Go knows is that we ran a function and either all test cases passed or some/all test cases failed.
  - Whenever we are testing with go, we need to make sure that we need to manually take care of any cleanup that is needed. Because the Go testing framework will not cleanup for us. 
    - For example if we open a file and do something, we need to manually cleanup for it ourselves after operations are done. 
    - The Go testing framework will not detect that we opened a file and did something and automatically clean it up for us.
    - We need to remember this as if for some reason our test crashes midway and the file is not cleaned up, then the next time we run the test case, we will get unwanted results.
    - So, we need to manually cleanup the file after the test runs or even if the test crashes mid way.
- Initaite go module:
  > go mod init cards
  - This will initiate a go module called cards and without go module you can not run tests.
- Formatting:
  - %T -> Type
  - %v -> Any value
- Printing in Go:
  ```
  fmt.println() // Adds new line at the end. Does not support formatting like %v and %T
  fmt.printf("\n") // Need to manually add the new line character. Supports %v and %T
  ```
---