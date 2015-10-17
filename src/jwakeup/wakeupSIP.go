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

func Hindex(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./www/index.html")
	if (err!=nil) {

	}
	//check(err)
	//fmt.Print(string(dat))

	fmt.Fprintf(w, string(dat))
}

func (wH wakeupSIP) wSIPstart(port string, nChan chan string) {
	fmt.Println("Starting HTTP server...")
	wH.toMain = nChan
	http.HandleFunc("/", Hindex)
	http.ListenAndServe(port, nil)
}
func (wH wakeupSIP) wSIPstop() {
	fmt.Println("Stopping HTTP server...")
}

type wakeupSIP struct {
	toMain chan string
	
}
