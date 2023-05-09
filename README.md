# udemy-golang

My practice code with Udemy Golang course <https://www.udemy.com/go-the-complete-developers-guide/>

---

# Notes

## FAQ

- Golang is a Compiled language.
- To run it on our local computer we need the Go Runtime which allows us to build and execute Go programs.

## Questions

### 1. helloworld
- How do we run the go code in our project?
  - The 'go' command gives us the ability to compile and execute our projects.
  - Command to write the project: "go run main.go"
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
