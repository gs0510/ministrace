package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("Running ", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.Start()

	err := cmd.Wait()
	if err != nil {
		fmt.Println("Wait returned ", err)
	}
}
