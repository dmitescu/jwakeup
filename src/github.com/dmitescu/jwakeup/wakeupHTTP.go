/*
   JWakeup
   Copyright (c) 2015 
   Mitescu George Dan <d.mitescu@jacobs-university.de>
   Nicolae Andrei <an.nicolae@jacobs-university.de>
   Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
   Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>
*/

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"strings"
	"encoding/json"
	"strconv"
	"time"
)

/*
   Structs used for form parsing
*/
type messageLogin struct {
	Token string `json:"token"`
	User string `json:"user"`
}

type messagePhone struct {
  	Phone string `json:"phone"`
}

/*
   Handler-functions for HTTP requests
*/
func (wH *wakeupHTTP) Hindex(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("jwakeup_token")
	if (err != nil){
		dat, _ := ioutil.ReadFile("./www/index.html")
		fmt.Fprintf(w, string(dat))
	} else {
		http.Redirect(w, r, "/home", 302)
	}
}

func (wH *wakeupHTTP) Hlogo1(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("./www/logo.png")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(dat)))
	w.Write(dat)
}

func (wH *wakeupHTTP) Hlogo2(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("./www/jwakeup_logo.png")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(dat)))
	w.Write(dat)
}

func (wH *wakeupHTTP) Hclock(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("./www/clock.png")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(dat)))
	w.Write(dat)
}

func (wH *wakeupHTTP) Hback2(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("./www/background2.jpg")
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Length", strconv.Itoa(len(dat)))
	w.Write(dat)
}

func (wH *wakeupHTTP) HStyle(w http.ResponseWriter, r *http.Request){
	dat, _ := ioutil.ReadFile("./www/style.css")
	w.Write(dat)
}

func (wH *wakeupHTTP) Hlogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
        // logic part of log in
        newtoken := login(r.Form["username"][0], r.Form["password"][0])
	if newtoken == "" {
		//fmt.Println("Wrong login!")
		http.Redirect(w, r, "/", 302)
	} else {
		var newuser wUser
		newuser.token = newtoken
		newuser.username = r.Form["username"][0]
		newuser.phonenr = phone_number(newtoken)

		// sending user to the list
		wH.messC <- "adduser"
		wH.toMainU <- newuser

		// adding cookie
		nCookie := http.Cookie{Name: "jwakeup_token",
			Value: newtoken,
			Expires: time.Now().Add(time.Minute*10)}
		http.SetCookie(w, &nCookie)

		// redirecting to home page
		http.Redirect(w, r, "/home", 302)
	}
	
}

func (wH *wakeupHTTP) Hhome(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("jwakeup_token")
	if (err != nil){
		http.Redirect(w, r, "/", 302)
	} else {
		fmt.Println(cookie.Value)
		dat, _ := ioutil.ReadFile("./www/Home.html")
		fmt.Fprintf(w, string(dat))
	}
}

/*
   Start of the main HTTP server
*/
func (wH *wakeupHTTP) wHTTPstart(port string,
	nuc chan wUser, ncc chan wCall, nmessC chan string) {
	
	fmt.Println("Starting HTTP server...")

	wH.toMainU = nuc
	wH.toMainC = ncc
	wH.messC = nmessC
	
	http.HandleFunc("/", wH.Hindex)
	http.HandleFunc("/login", wH.Hlogin)
	http.HandleFunc("/home", wH.Hhome)

	//TODO: Add regexp for png, jpg and remote files
	http.HandleFunc("/style.css", wH.HStyle)
	http.HandleFunc("/logo.png", wH.Hlogo1)
	http.HandleFunc("/jwakeup_logo.png", wH.Hlogo2)
	http.HandleFunc("/clock.png", wH.Hclock)
	http.HandleFunc("/background2.jpg", wH.Hback2)
	http.ListenAndServe(port, nil)

}

/*
   Stopping the HTTP server...
*/

func (wH *wakeupHTTP) wHTTPstop() {
	fmt.Println("Stopping HTTP server...")
}

/*
   OpenJUB API handling
*/
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

/*
   Main class
*/
type wakeupHTTP struct {
	toMainU chan wUser
	toMainC chan wCall
	messC chan string
}
