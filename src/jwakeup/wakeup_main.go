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
)

func main(){
	//reader := bufio.NewReader(os.Stdin)
	//var keyIn string = "hmm"

	fmt.Println("Starting JWakeup...")
	
	var mainHTTP wakeupHTTP
	var mainSIP wakeupSIP

	userC := make(chan wUser)
	callC := make(chan wCall)
	messC := make(chan string)

	
	go mainHTTP.wHTTPstart(":8080", userC, callC, messC)
	time.Sleep(time.Second * 2)
	
	go mainSIP.wSIPstart(":5051", "127.0.0.1", userC, callC, messC)
	time.Sleep(time.Second * 60)
	mainHTTP.wHTTPstop()
	
}
