// In this example, we are going to use the helloGreetings package
package main

import(

	"fmt"
	"~/go/src/gofullstack/hellogreetings"


)

func main() {

	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	println("The value of the Magic Number is: ", greetingspackage.MagicNumber())
}