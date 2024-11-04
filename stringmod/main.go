// importing the module ie is created

package main

import (
	"fmt"

	"github.com/aditya/stringmod/odd_even"
	"github.com/aditya/stringmod/quotes"
)

func main() {
	o, e := odd_even.CountOddEven("1234566789")
	fmt.Println(o, e) //32

	fmt.Println(quotes.Getemoji("turtle"))
}
