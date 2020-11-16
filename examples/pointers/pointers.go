package main

import (
	"fmt"
	"github.com/zukeep/pointers/v1"
)

// requiresStrPointer is a function that receives a string pointer and checks if it's null
func requiresStrPointer(str *string) {
	if str == nil {
		fmt.Println("Input str is nil")
		return
	}

	fmt.Printf("Printing string from a pointer: %v\n", *str)
}

func main() {
	// Normally, in order to call requiresPointer you'll need a string variable as follows:
	strVar := "Variable String"
	// And you would call requiresStrPointer with StrVar
	// Calling it as follows requiresStrPointer(&"String") will not compile.
	requiresStrPointer(&strVar)

	// To keep your code neater you can use the pointers package and call requiresStrPointer as follows:
	requiresStrPointer(pointers.String("Wrapped by pointers package"))

	// And because requiresStrPointer can receive nil value, we can safely call it with nil
	requiresStrPointer(nil)

	// Pointers provide wrappers for all primitive types as can be seen in this anonymous struct example:
	s := struct {
		Exists         *bool
		Temperature    *float32
		NumberOfPeople *int64
		InitialLetter  *rune
	}{
		Exists:         pointers.Bool(true),
		Temperature:    pointers.Float32(42.645),
		NumberOfPeople: pointers.Int64(42),
		InitialLetter:  pointers.Rune('D'),
	}

	fmt.Println(s)
}
