package main

import "fmt"

func main() {
    a := 10
    for i:=0; i<a; i++ {
        printAsterisk(i)
        fmt.Print("\n")
    }
}

func printAsterisk(numOfAsterisk int) {
    for i:=0; i<numOfAsterisk; i++ {
        fmt.Print("*")
    }
}
