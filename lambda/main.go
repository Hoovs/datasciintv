package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/usr/bin/python3", "-c", "print(\"hello\")")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}
