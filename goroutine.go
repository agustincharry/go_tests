package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go slow("Hola", 1000)
	go slow("abcd", 1500)
	fmt.Println("Waiting...")

	var wait string
	fmt.Scan(&wait)
}

func slow(name string, milliseconds int32) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		fmt.Println(letter)
	}
}
