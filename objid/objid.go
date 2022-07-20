package main

import (
	"fmt"

	"github.com/rs/xid"
)

// New returns a globally unique ID.
func New() string {
	return xid.New().String()
}

func main() {
	fmt.Println(New())
}
