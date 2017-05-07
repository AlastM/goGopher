package main

import "fmt"

var lightSwitchIsOn bool = false


func main() {
	
	printlightSwitchState()
	togglelightSwitch()
	printlightSwitchState()
	togglelightSwitch()
	printlightSwitchState()
}

func printlightSwitchState() {
	fmt.Println("The light swift is off:", lightSwitchIsOn)

}

func togglelightSwitch() {

// ###togglelightSwitch, when the light is on the "!lightSwitchIsOn" <--"invert" same as off

	lightSwitchIsOn = !lightSwitchIsOn
}