package main

import "fmt"

func main() {
	fmt.Println("hey ,it is first go project")
	for i := 0; i < 2; i++ {
		say("love go")
	}

}
func say(message string) {
	fmt.Print("You Said : ", message)
}
