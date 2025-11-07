package main

import (
    "fmt"
    "example.com/mymodule/calc"
    "example.com/mymodule/stringutils"
)

func main() {
    str := "GoLang"
    num := 5

    fmt.Println("Теріс айналған жол:", stringutils.Reverse(str))
    fmt.Println("Санның квадраты:", calc.Square(num))
}
