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

func (wH *wakeupHTTP) Hindex(w http.ResponseWriter, r *http.Request) {
	dat, _ := ioutil.ReadFile("../../www/index.html")
	fmt.Fprintf(w, string(dat))
}

func (wH *wakeupHTTP) Hlogo(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("../../www/logo.png")
	fmt.Fprintf(w, string(dat))
}

func (wH *wakeupHTTP) Hlogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
        // logic part of log in
        newtoken := login(r.Form["username"][0], r.Form["password"][0])
	if newtoken == "" {
		fmt.Println("Wrong login!")
	} else {
		var newuser wUser
		newuser.token = newtoken
		newuser.username = r.Form["username"][0]
		
		wH.messC <- "adduser"
		wH.toMainU <- newuser
	}
	
}

func (wH *wakeupHTTP) wHTTPstart(port string,
	nuc chan wUser, ncc chan wCall, nmessC chan string) {
	
	fmt.Println("Starting HTTP server...")

	wH.toMainU = nuc
	wH.toMainC = ncc
	wH.messC = nmessC
	
	http.HandleFunc("/", wH.Hindex)
	http.HandleFunc("/login", wH.Hlogin)
	http.HandleFunc("logo.png", wH.Hlogo)
	http.ListenAndServe(port, nil)

}

func (wH *wakeupHTTP) wHTTPstop() {
	fmt.Println("Stopping HTTP server...")
}

func login(uname string, pass string) string{
	url := "https://api.jacobs-cs.club/auth/signin"

	jsonStr := append([]byte("{\"username\":\""), []byte(uname)...)
	jsonStr = append(jsonStr, []byte("\",\"password\":\"")...)
	jsonStr = append(jsonStr, []byte(pass)...)
	jsonStr = append(jsonStr, []byte("\"}")...)
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

func phone_number(sUser wUser) string{
	url := "https://api.jacobs-cs.club/user/me?token="+sUser.token

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
}
