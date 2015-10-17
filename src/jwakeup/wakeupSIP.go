// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	//"time"
)

func (wH wakeupSIP) wSIPstart(port string,
	nuc chan wUser, ncc chan wCall) {
	fmt.Println("Starting SIP server...")

	wH.fromMainU = nuc
	wH.fromMainC = ncc
	
}

func (wH wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")
}

func (wH wakeupSIP) addCALL(nCall wCall){
	wH.callList = append(wH.callList, nCall)
	fmt.Println("Added call to ", nCall,
		"at", nCall.calltime)
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
	callList []wCall
}
