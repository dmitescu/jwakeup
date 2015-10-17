// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"time"
)

type wUser struct{
	username string `xml:"username"`
	token string `xml:"token"`
}

type wCall struct{
	Callid int `xml:"callid"`
	Phonenr string `xml:"phonenr"`
	Calltime time.Time `xml:"calltime"`
}

type wCallList struct{
	WCallList []wCall `xml:"entry"`
}
