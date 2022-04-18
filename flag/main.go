package main

import (
	"flag"
	"fmt"
)

var (
	intflag    *int
	boolflag   *bool
	stringflag *string
)

func init() {
	intflag = flag.Int("intflag", 0, "int flag value")
	boolflag = flag.Bool("boolflag", false, "bool flag value")
	stringflag = flag.String("stringflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("int flag:", *intflag)
	fmt.Println("bool flag:", *boolflag)
	fmt.Println("string flag:", *stringflag)
}
