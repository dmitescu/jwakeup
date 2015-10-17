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

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	uTest.username = "Dinamo"
	uTest.token = "lalala"
	cTest.Callid = 1
	cTest.Phonenr = "5111"
	cTest.Calltime, _ = time.Parse(longForm,
		"Feb 3, 2013 at 7:54pm (PST)")
	
	userC := make(chan wUser)
	callC := make(chan wCall)
	
	fmt.Println("Starting service... (write 'quit' to stop)")
	go mainHTTP.wHTTPstart(":8080", userC, callC)
	go mainSIP.wSIPstart(":5051", userC, callC)
	
	time.Sleep(time.Second * 2)
	mainSIP.logUSER(uTest)
	mainSIP.addCALL(cTest)
	cTest.Callid = 2
	mainSIP.addCALL(cTest)
	time.Sleep(time.Second * 3)

	mainHTTP.wHTTPstop()
	mainSIP.wSIPstop()
	
}
