package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	tm "github.com/buger/goterm"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //darwin example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	tm.Clear() // Clear current screen
	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)
		tm.Println("Current Time:", time.Now().Format(time.RFC1123))
		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second)
	}
}

func main1() {
	fmt.Println("I will clean the screen in 2 seconds!")
	time.Sleep(2 * time.Second)
	CallClear()
	fmt.Println("I'm alone...")
}
