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

func (wH wakeupSIP) addCALL(){

}

func (wH wakeupSIP) makeCALL(){

}

func (wH wakeupSIP) logUSER(){
	
}

type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	
}
