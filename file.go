package main

import (
	"os"
)

func mkdir(name string) error {
	err := os.Mkdir(name, 0755)
	return err
}

func chdir(name string) error {
	err := os.Chdir(name)
	return err
}

func writeFile(name string, data []byte) error {
	err := os.WriteFile(name, data, 0666)
	return err
}
