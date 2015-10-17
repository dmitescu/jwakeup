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
	. "sync"
)

func Hindex(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("../../www/index.html")
	if (err!=nil) {

	}
	//check(err)
	//fmt.Print(string(dat))

	fmt.Fprintf(w, string(dat))
}

func (wH *wakeupHTTP) wHTTPstart(port string,
	nuc chan wUser, ncc chan wCall, nmessC chan string,
	num *Mutex, ncm *Mutex, nmessM *Mutex) {
	
	fmt.Println("Starting HTTP server...")

	wH.toMainU = nuc
	wH.toMainC = ncc
	wH.messC = nmessC

	wH.mutexMainU = num
	wH.mutexMainC = ncm
	wH.messM = nmessM
	
	http.HandleFunc("/", Hindex)
	http.ListenAndServe(port, nil)

}

func (wH *wakeupHTTP) wHTTPstop() {
	fmt.Println("Stopping HTTP server...")
}

type wakeupHTTP struct {
	toMainU chan wUser
	toMainC chan wCall
	messC chan string

	mutexMainU *Mutex
	mutexMainC *Mutex
	messM *Mutex
}
