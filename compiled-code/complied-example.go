package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I am a compiled computer programme\nI am written in Go but you don't have Go installed on this machine.\nInstead Go programmes are complied (built) into binary executables that target a specifc operating system and processor architecture")
	fmt.Println("This programme will self-destruct after 30 seconds")
	time.Sleep(30000 * time.Millisecond)
}
