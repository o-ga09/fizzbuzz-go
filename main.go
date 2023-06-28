package main

import (
	"fmt"
	"main/fizzbuzz"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("number : %s\n", fizzbuzz.Convert(i+1))
		time.Sleep(1 * time.Second)
	}
}