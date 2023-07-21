package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	newUUID := uuid.New()

	// Print the generated UUID
	fmt.Println(newUUID)
}
