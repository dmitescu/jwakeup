/*
   JWakeup
   Copyright (c) 2015 
   Mitescu George Dan <d.mitescu@jacobs-university.de>
   Nicolae Andrei <an.nicolae@jacobs-university.de>
   Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
   Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>
*/

package main

import (
	"time"
)

/*
   Database user
*/

type wUser struct {
	username string `xml:"username"`
	token string `xml:"token"`
	phonenr string
}

/*
   Database call
*/
type wCall struct {
	Callid int `xml:"callid"`
	Phonenr string `xml:"phonenr"`
	Calltime time.Time `xml:"calltime"`
}

/*
   Database call list (deprecated?)
*/
type wCallList struct {
	WCallList []wCall `xml:"entry"`
}




		
