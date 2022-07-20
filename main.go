package main

import (
	"fmt"
	"strings"
)

var x int
var y int
var asdm string

func main() {
	fmt.Println("Hello! Welcome to Flash's Calculator")
	fmt.Println("What Would your first number be?")	
        fmt.Scan(&x)		
        fmt.Println("ok now what is your second number")
        fmt.Scan(&y)
        fmt.Println("ok now are you going yo add subtract multiply or divide?")
	fmt.Scan(&asdm)	
        if strings.EqualFold(asdm, "add") {
	fmt.Println("ok your answer will be:", x + y)
		}
  if strings.EqualFold(asdm, "subtract") {
    fmt.Println("ok your answer will be:", x - y)
  
}
  if strings.EqualFold(asdm, "divide") {
    fmt.Println("ok your answer will be:", x / y)
  }
  if strings.EqualFold(asdm, "multiply") {
  fmt.Println("ok your answer will be:", x * y)  
  }
}
