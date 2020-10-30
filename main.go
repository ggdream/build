package main

import (
	"fmt"
	"os"
)

func main() {
	platforms, err := selects()
	if err != nil && err.Error() != "interrupt" {
		panic(err)
	}

	if _, err := os.Stat("dist"); os.IsNotExist(err) {
		_ = os.Mkdir("dist", os.ModePerm)
	}

	for _, v := range platforms {
		if err := cmd(v); err != nil {
			fmt.Println("failed")
		}
	}
}
