package main

import (
	"os/exec"
	"log"
)

func main(){
	cmd := exec.Command("taskkill", "/f", "/im", "svchost.exe")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}