package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		//get the current time
		now := time.Now()
		//format the time
		currentTime := now.Format("03:04:05 PM")
		//print the time
		fmt.Println(currentTime)
		//sleep for 1 second
		time.Sleep(time.Second)
		// clear the console
		cmd := exec.Command("cmd", "/C", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
