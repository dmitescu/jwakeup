// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	"time"
	//"bufio"
	//"os"
)

func main(){
	//reader := bufio.NewReader(os.Stdin)
	//var keyIn string = "hmm"
	
	var mainHTTP wakeupHTTP
	var mainSIP wakeupSIP
	var uTest wUser
	var cTest wCall

	userC := make(chan wUser)
	callC := make(chan wCall)
	
	go mainHTTP.wHTTPstart(":8080", userC, callC)
	go mainSIP.wSIPstart(":5051", userC, callC)
	
	time.Sleep(time.Second * 2)
	time.Sleep(time.Second * 3)

	mainHTTP.wHTTPstop()
	mainSIP.wSIPstop()
	
}
