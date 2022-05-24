package main

import (
	"fmt"

	"github.com/cmilheim/greetings"
)

func main() {
	// get a greeting message and print it.
	message := greetings.Hello("Chris")
	fmt.Println(message)
}
