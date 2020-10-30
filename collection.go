package main

import (
	"github.com/AlecAivazis/survey/v2"
	"os/exec"
	"strings"
)

var coll []string

func init() {
	c := exec.Command("go", "tool", "dist", "list")
	data, err := c.Output()
	if err != nil {
		return
	}

	coll = strings.Split(string(data), "\n")
	coll = coll[0 : len(coll)-1]
}

func selects() (platforms []string, err error) {
	prompt := &survey.MultiSelect{
		Message: "Select target platforms:",
		Options: coll,
	}
	err = survey.AskOne(prompt, &platforms)
	return
}
