/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package main

import (
	"os"

	"github.com/niuzhiqiang90/yapi-user-operator/cmd"
)

func main() {
	command := cmd.NewRootCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
