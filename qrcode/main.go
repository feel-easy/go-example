// main.go
package main

import (
	"fmt"
	"os"

	qrcode "github.com/tuotoo/qrcode"
)

func main() {
	fi, err := os.Open("qrcode.png") // 默认到$GOPATH/src/ 下找
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}
