package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	name    string
	version = "0.0.1"
)

func init() {
	c := exec.Command("go", "list")
	text, err := c.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	a := strings.Split(strings.TrimSuffix(string(text), "\n"), "/")
	name = a[len(a)-1]
}
