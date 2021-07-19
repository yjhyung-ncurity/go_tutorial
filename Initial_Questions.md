## 1. How do we run code in our project?
#   - 1) Navigate to the project directory 
#   - 2) Use Go command (used to build and compile the code written in Go)
#   - 3) Ex) go run [filename.go] , go build [filename.go]
#   - 4) Go CLI commands : build, run, fmt, install, get, test
#   - 5) Difference between 'go run' and 'go build':
#       - Go Run Command: immediately compile and execute the program file
#       - Go Build Command: Only compiles the program file 
#

## 2. What does Package Main mean? 
#   - 1) Package is a collection of common source code files 
#   - 2) Packages can have many files inside of it with an ending extensions of .go
#       - (main.go , packages.go, test1.go etc)
#   - 3) Requirement: every file inside of a package must declare the package that it belongs to at the top of the file
#   - 4) Difference between main.go and other files.go:
#       - Go has 2 Types of Packages : Executable , Resusable 
#           Executables : Generates a runnable file
#           Resuable    : files with reusable logics or functions 
#   - 5) ** Name of the packages used determines executable or reusable files. (Executable Files needs to be declared under main package)

## 3. What does import fmt mean?
#   - 1) import statement is used to give package the access to some code that is written inside of another package
#   - 2) by importing FMT gives package an access to all of the code and functionality written in a packaged called FMT
#   - 3) FMT is the (defualt) standard library package provided by Go Language (Format for short)
#   - 4) Other Standard Library may include : debug, math, encoding, crypto, io, fmt etc. 
#   - 5) Can also import reusable package created by other engineers or self created
#   - REFERENCE: golang.org/pkg

## 4. What's that 'func' thing?
#   - 1) func is short for function 
#   - 2) func main(){ }
#           func: declaring function
#           main: name of the function
#           ()  : list of arguments to pass- parameters
#           {}  : function body to run the code 

## 5. How is the main.go file organized?
#   - 1) declare package 
#           ex) package main
#   - 2) import other packages needed / used
#           ex) import ("fmt" "math" "io")
#   - 3) declare functions 
#           ex) func main(){ fmt.Println("HELLO WORLD")}
