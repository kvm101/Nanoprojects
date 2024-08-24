package main

import (
	"hcmd/hcmd"
	"log"
	"net/http"
)

func offHandler(w http.ResponseWriter, r *http.Request) {
	hcmd.Exec("poweroff")
	log.Println("Computer off")
}

func RestartHandler(w http.ResponseWriter, r *http.Request) {
	hcmd.Exec("reboot")
	log.Println("Computer is restarted")
}

func suspendHandler(w http.ResponseWriter, r *http.Request) {
	hcmd.Exec("systemctl suspend")
	log.Println("Suspended mode")
}

func main() {
	http.HandleFunc("/turnOFF", offHandler)
	http.HandleFunc("/restart", RestartHandler)
	http.HandleFunc("/suspend", suspendHandler)

	port := ":8000"
	http.ListenAndServe(port, nil)
}
