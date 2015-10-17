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
	"bytes"
	"strings"
	"encoding/json"
)

type messageLogin struct {
	Token string `json:"token"`
	User string `json:"user"`
}

type messagePhone struct {
  	Phone string `json:"phone"`
}

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

func login() string{
	url := "https://api.jacobs-cs.club/auth/signin"

    var jsonStr = []byte("{\"username\":\"dmitescu\",\"password\":\"!1QqAaZz\"}")
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()


    body, _ := ioutil.ReadAll(resp.Body)
	
	
	dec := json.NewDecoder(strings.NewReader(string(body)))
	var m messageLogin
	err2 := dec.Decode(&m)
	if err2 != nil {
		panic(err2)
	}
	
	return m.Token
}

func phone_number() string{
	Token := login()
	url := "https://api.jacobs-cs.club/user/me?token="+Token

	client := &http.Client{}
	req, err := client.Get(url)
	if err != nil {
        panic(err)
    }

	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	dec := json.NewDecoder(strings.NewReader(string(body)))


	var m messagePhone
	err2 := dec.Decode(&m)
	if err2 != nil {
		panic(err2)
	}
	return m.Phone
}


type wakeupHTTP struct {
	toMainU chan wUser
	toMainC chan wCall
	messC chan string

	mutexMainU *Mutex
	mutexMainC *Mutex
	messM *Mutex
}
