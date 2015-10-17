// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"encoding/xml"
	"fmt"
	//"time"
	"io/ioutil"
	//"os"
)

func (wH wakeupSIP) wSIPstart(port string,
	nuc chan wUser, ncc chan wCall) {
	fmt.Println("Starting SIP server...")

	wH.fromMainU = nuc
	wH.fromMainC = ncc

	tempin, _ := ioutil.ReadFile("../../userbase/wakelist.xml")
	err := xml.Unmarshal(tempin, &wH.callList)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		for _, sCall := range wH.callList.WCallList {
			fmt.Println(sCall.Callid)
		}
	}
}

func (wH wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")
	//tempout, _ := xml.MarshalIndent(wH.callList, "  ", "    ")
	//ioutil.WriteFile("../../userbase/wakelist.xml", tempout, 0644)
}

func (wH wakeupSIP) addCALL(nCall wCall){
	wH.callList.wCallList = append(wH.callList.wCallList, nCall)
	fmt.Println("Added call to ", nCall.Phonenr,
		"at", nCall.Calltime)

	tempout, _ := xml.MarshalIndent(wH.callList, "  ", "    ")

	ioutil.WriteFile("../../userbase/wakelist.xml", tempout, 0644)
}

func (wH wakeupSIP) makeCALL(){

}

func (wH wakeupSIP) logUSER(nUser wUser){
	wH.loggedList = append(wH.loggedList, nUser)
	fmt.Println("User", nUser.username, "logged in!")
}

func (wH wakeupSIP) checkUSER(sUser string){
	for _, dUser := range wH.loggedList{
		if(dUser.username == sUser) {
				//Transmis mesaj gulie
		}
	}
}

type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	loggedList []wUser
	callList wCallList
}
