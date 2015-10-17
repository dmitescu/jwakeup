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

func (wH *wakeupSIP) wSIPstart(port string,
	nuc chan wUser, ncc chan wCall) {
	fmt.Println("Starting SIP server...")

	wH.fromMainU = nuc
	wH.fromMainC = ncc

	tempin, _ := ioutil.ReadFile("../../userbase/wakelist.xml")
	var listin wCallList
	err := xml.Unmarshal(tempin, &listin)
	wH.callList = listin.WCallList

	if err != nil {
		fmt.Println("Error: ", err)
	} 
}

func (wH *wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")

}

func (wH *wakeupSIP) addCALL(nCall wCall){
	wH.callList = append(wH.callList, nCall)
	fmt.Println("Added call to", nCall.Phonenr,
		"at", nCall.Calltime)
	
	var listout wCallList
	listout.WCallList = wH.callList
	tempout, _ := xml.MarshalIndent(listout, "  ", "    ")
	ioutil.WriteFile("../../userbase/wakelist.xml", tempout, 0644)
}

func (wH *wakeupSIP) makeCALL(){

}

func (wH *wakeupSIP) logUSER(nUser wUser){
	wH.loggedList = append(wH.loggedList, nUser)
	fmt.Println("User", nUser.username, "logged in!")
}

func (wH *wakeupSIP) logoutUSER(nUser wUser) bool{
	statU, indexU := checkUSER(nUser.username)
	if statU == false {
		return false
	} else {
		wH.callList = append(wH.callList[:indexU],
			callList[indexU+1:])
		return true
	}
}

func (wH *wakeupSIP) checkUSER(sUser string) (bool, int){
	for iUser, dUser := range wH.loggedList{
		if(dUser.username == sUser) {
			return (true, iUser)
		}
	}
}

type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	loggedList []wUser
	callList []wCall
}
