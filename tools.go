package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	goOS   = "GOOS"
	goARCH = "GOARCH"
)

func env(goos, goarch string) error {
	if err := os.Setenv(goOS, goos); err != nil {
		return err
	}
	return os.Setenv(goARCH, goarch)
}

func cmd(target string) error {
	info := strings.Split(target, "/")
	if err := env(info[0], info[1]); err != nil {
		return err
	}

	var suffix string
	if info[0] == "windows" {
		suffix = ".exe"
	}
	e := exec.Command("go", "build", "-o", fmt.Sprintf("./dist/%s-%s-%s-%s%s", name, version, info[0], info[1], suffix))
	fmt.Printf("\n%s		>> %s-%s\n", time.Now().Format("2006-01-02 15:04:05"), info[0], info[1])
	_, err := e.Output()
	fmt.Printf("%s		<< %s-%s\n", time.Now().Format("2006-01-02 15:04:05"), info[0], info[1])
	if err != nil {
		return err
	}
	return nil
}
