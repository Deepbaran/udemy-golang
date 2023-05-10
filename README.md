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

## Questions

### 1. helloworld
- How do we run the go code in our project?
  - The 'go' command gives us the ability to compile and execute our projects.
  - Command to write the project: 
    - > go run main.go"
    - run multiple files:
      - > go run main.go filename.go
  - Go CLI tools
    - go build -> Compiles a bunch of go source code files (Creates an executable file)
    - go run -> Compiles and executes one or two files
    - go fmt -> Formats all the code in each file in the current directory
    - go install -> Compile and "installs" a package
    - go get -> Downloads the raw source code of someone else's package
    - go test -> Runs any tests associated with the current project
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
  - Iterating throug slices:
    - > for i, card := range cards {fmt.Println(i, card)}
    - i -> index of the element inside the slice
    - card -> Current card we're iterating over
    - range cards -> Take the slice of 'cards' and loop over it
    - We are using the := operator to declare and intialize i and card, as we are throwing out i and card after each iteration. So, it is not reassigning, rather it is a completely new declaration and initialization.
    - Now, as we know each variable tat is declared in our code must be used, otherwise code will throw error, so if we do not wish to use the index of the slices while iterating over it, we can use '_' instead of 'i'.
      - > for _, card := range cards {fmt.Println(card)}
  - More on Slices:
    - Slices are 0-indexed
    - range in slice:
      - > cards[startIndexIncluding : upToNotIncluding]
      - > cards[ : upToNotIncluding] // [0 : upToNotIncluding]
      - > cards[startIndexIncluding : ] // [startIndexIncluding : len(cards)]
      - This much like list slicing in python.
      - Slicing slices does not modify the real slice.
- Custom Types in Go:****
  - We can "extend" a base type (string, integer, float, array, map) and add some extra functionality to it.
  - Tell Go we want to create an array of strings and add a bunch of functions specifically made to work with it:
    - > type deck []string{}
      - deck will extend every type of behavior of a slice of string.
    - Functions, with 'deck' as a 'receiver' -> A function with a receiver is like a "method" - afunction that belongs to an "instance" So, these custom functions will only work with this 'deck' type.
    - Now in the codebase we can replace all slice of string with deck to create new instances of deck type if we wished to.
    - > cards := deck{"Ace of Diamonds", newCard()}
    - Syntax to add functionality/method with the receiver deck:
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
    