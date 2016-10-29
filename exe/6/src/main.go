package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    promptName(reader)
    specialCharCheckerHandler(reader)
    alphabeticStringChecker(reader)
    numericStringChecker(reader)
    lowercaseStringChecker(reader)
    uppercaseStringChecker(reader)
}

func promptName(reader *bufio.Reader) {
    fmt.Print("Please Enter your name: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}

func specialCharCheckerHandler(reader *bufio.Reader) {
    fmt.Println("<< Special Character Checker >>")
    // _, _ := reader.ReadString('\n')
    if true {
        fmt.Println("There is no special character inside.")
    } else {
        fmt.Println("There are special character inside.")
    }
}

func alphabeticStringChecker(reader *bufio.Reader) {
    fmt.Println("<< AlphabeticStringChecker >>")
    if true {
        fmt.Println("")
    } else {
        fmt.Println("")
    }
}

func numericStringChecker(reader *bufio.Reader) {
    fmt.Println("<< NumericStringChecker >>")
    if true {
        fmt.Println("")
    } else {
        fmt.Println("")
    }
}

func lowercaseStringChecker(reader *bufio.Reader) {
    fmt.Println("<< Lowercase String Checker >>")
    if true {
        fmt.Println("")
    } else {
        fmt.Println("")
    }
}

func uppercaseStringChecker(reader *bufio.Reader) {
    fmt.Println("<< Uppercase String Checker >>")
    if true {
        fmt.Println("")
    } else {
        fmt.Println("")
    }
}
