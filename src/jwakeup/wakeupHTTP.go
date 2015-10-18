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
	"time"
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
	w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(len(dat.Bytes())))
	w.Write(dat.Bytes())
	fmt.Fprintf(w, string(dat))
}

func (wH *wakeupHTTP) Hlogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
        // logic part of log in
        newtoken := login(r.Form["username"][0], r.Form["password"][0])
	if newtoken == "" {
		fmt.Println("Wrong login!")
		fmt.Fprintf(w, "<html><body><script>window.location.assign(\"/\")</script></body></html>")
	} else {
		var newuser wUser
		newuser.token = newtoken
		newuser.username = r.Form["username"][0]
		newuser.phonenr = phone_number(newtoken)
			
		wH.messC <- "adduser"
		wH.toMainU <- newuser

		var newCookie http.Cookie
		newCookie.Name = "logtoken"
		newCookie.Value = newtoken
		newCookie.Expires = time.Now().Add(time.Minute*10)
		//newCookie.Raw = ""
		//newCookie.Unparsed = {""}
		http.SetCookie(w, &newCookie)

		fmt.Fprintf(w, "<html><body><script>window.location.assign(\"/home\")</script></body></html>")
	}
	
}

func (wH *wakeupHTTP) Hhome(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("logtoken")
	if (err != nil){
		fmt.Fprintf(w,"<html><body><script>window.location.assign(\"/\")</script></body></html>")
	} else {
		fmt.Println(cookie.Value)
		dat, _ := ioutil.ReadFile("../../www/Home.html")
		fmt.Fprintf(w, string(dat))
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
	http.HandleFunc("/home", wH.Hhome)
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

func phone_number(token string) string{
	url := "https://api.jacobs-cs.club/user/me?token="+token

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
