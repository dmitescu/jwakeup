// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
)

func main(){
	var mainHTTP wakeupHTTP
	mainC := make(chan string)
	fmt.Println("Starting service...")
	mainHTTP.wHTTPstart(":8080", mainC)
	mainHTTP.wHTTPstop()
	fmt.Println("Starting SIP client...")

	
}
