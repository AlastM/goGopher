package greetingspackage

import "fmt"


var MagicNumber int

// This is an exported functions and can be called outside the package
func GopherGreetings() {

	fmt.Println("A very jolly hallo my fellow Gophers! I'm printing from GopherGreetings() function ;) ")
	// now we are calling an unexported package, because this function
	// is within the same package, we have acces to it all
	printGreetingsUnexported()
}

// example of a packages init() function

func init() {

	MagicNumber = 108
}