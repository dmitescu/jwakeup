// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	. "sync"
)

func (wS *wakeupSIP) wSIPstart(port string, dest string,
	nuc chan wUser, ncc chan wCall, nmessC chan string,
	num *Mutex, ncm *Mutex, nmessM *Mutex) {
	fmt.Println("Starting SIP server...")

	cMess := ""
	
	wS.messC = nmessC
	wS.fromMainU = nuc
	wS.fromMainC = ncc

	wS.messM = nmessM
	wS.mutexMainU = num
	wS.mutexMainC = ncm

	tempin, _ := ioutil.ReadFile("../../userbase/wakelist.xml")
	var listin wCallList
	err := xml.Unmarshal(tempin, &listin)
	wS.callList = listin.WCallList

	if err != nil {
		fmt.Println("Error: ", err)
	}

	for cMess != "terminate" {

	}
	
}

func (wS *wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")

}

func (wS *wakeupSIP) addCALL(nCall wCall){
	wS.callList = append(wS.callList, nCall)
	fmt.Println("Added call to", nCall.Phonenr,
		"at", nCall.Calltime)
	
	var listout wCallList
	listout.WCallList = wS.callList
	tempout, _ := xml.MarshalIndent(listout, "  ", "    ")
	ioutil.WriteFile("../../userbase/wakelist.xml", tempout, 0644)
}

func (wS *wakeupSIP) makeCALL(){

}

func (wS *wakeupSIP) logUSER(nUser wUser){
	wS.loggedList = append(wS.loggedList, nUser)
	fmt.Println("User", nUser.username, "logged in!")
}

func (wS *wakeupSIP) checkUSER(sUser string) (bool, int){
	for iUser, dUser := range wS.loggedList{
		if(dUser.username == sUser) {
			return true, iUser
		}
	}
	return false, 0
}

func (wS *wakeupSIP) logoutUSER(nUser wUser) bool{
	statU, indexU := wS.checkUSER(nUser.username)
	if statU == false {
		return false
	} else {
		wS.callList = append(wS.callList[:indexU],
			wS.callList[(indexU+1):]...)
		return true
	}
}
type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	messC chan string

	mutexMainU *Mutex
	mutexMainC *Mutex
	messM *Mutex
	
	loggedList []wUser
	callList []wCall
}
