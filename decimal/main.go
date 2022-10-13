package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(decimal.NewFromInt(122).Shift(-2))
}
