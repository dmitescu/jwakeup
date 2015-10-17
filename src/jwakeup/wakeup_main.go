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
	var mainSIP wakeupSIP
	mainC := make(chan string)
	
	fmt.Println("Starting service... (write 'quit' to stop)")
	mainHTTP.wHTTPstart(":8080", mainC)
	
	while()
	mainHTTP.wHTTPstop()

	
}
