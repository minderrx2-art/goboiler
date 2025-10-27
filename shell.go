package main

import (
	"os/exec"
)

func goModInit(name string) error {
	cmd := exec.Command("go", "mod", "init", name)
	err := cmd.Run()
	return err
}
