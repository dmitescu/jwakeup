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

func (wH wakeupHTTP) wHTTPstart() {
	fmt.Println("Starting HTTP server...")
}
func (wH wakeupHTTP) wHTTPstop() {
	fmt.Println("Stopping HTTP server...")
}


type wakeupHTTP struct {
}
