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
)

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
	
	//wS.callOut.send()
}

func (wS *wakeupSIP) logUSER(nUser wUser){
	wS.loggedList = append(wS.loggedList, nUser)
	fmt.Println("User", nUser.username, "logged in!")
}

func (wS *wakeupSIP) checkUSER(sToken string) (bool, int){
	for iUser, dUser := range wS.loggedList{
		if(dUser.token == sToken) {
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


func (wS *wakeupSIP) wSIPstart(port string, dest string,
	nuc chan wUser, ncc chan wCall, nmessC chan string) {
	fmt.Println("Starting SIP server...")

	wS.callOut.init(port, dest)
	
	cMess := ""
	
	wS.messC = nmessC
	wS.fromMainU = nuc
	wS.fromMainC = ncc

	tempin, _ := ioutil.ReadFile("../../userbase/wakelist.xml")
	var listin wCallList
	err := xml.Unmarshal(tempin, &listin)
	wS.callList = listin.WCallList

	if err != nil {
		fmt.Println("Error: ", err)
	}
	
	for cMess != "terminate" {
		cMess = <- wS.messC
		fmt.Println("Got input!")
		if(cMess == "adduser"){
			nu := <- wS.fromMainU
			wS.logUSER(nu)
		}
	}

	wS.wSIPstop()
}

func (wS *wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")

}

type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	messC chan string
	
	loggedList []wUser
	callList []wCall

	callOut UDPOutput
}
