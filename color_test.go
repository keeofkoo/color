package color

import (
	"fmt"
)

func ExampleColored() {
	s := Colored("Happy", Red, Underline) + " " + Colored("new", Gray, OnYellow, Blink) + " " + Colored("year", Green) + "!"
	fmt.Printf("%q\n", s)
	// Output:
	// "\x1b[4m\x1b[31mHappy\x1b[0m \x1b[5m\x1b[43m\x1b[30mnew\x1b[0m \x1b[32myear\x1b[0m!"
}
