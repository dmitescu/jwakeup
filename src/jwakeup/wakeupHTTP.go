// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./www/index.html")
	if (err!=nil) {

	}
	//check(err)
	//fmt.Print(string(dat))

	fmt.Fprintf(w, string(dat))
}

func (wH wakeupHTTP) wHTTPstart(port string, nChan chan string) {
	fmt.Println("Starting HTTP server...")
	wH.toMain = nChan
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
func (wH wakeupHTTP) wHTTPstop() {
	fmt.Println("Stopping HTTP server...")
}

type wakeupHTTP struct {
	toMain chan string
	
}
