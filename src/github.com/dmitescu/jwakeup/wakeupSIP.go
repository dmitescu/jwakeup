// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	//"io/ioutil"
	//"encoding/xml"
	//"strconv"
)

//-------------------------------------
// Main SIP functions which treat both
//      SIP and API related issues
//-------------------------------------

func (wS *wakeupSIP) addCALL(nCall wCall){
	wS.callList = append(wS.callList, nCall)
	fmt.Println("Added call to", nCall.Phonenr,
		"at", nCall.Calltime)
	
	//var listout wCallList
	//listout.WCallList = wS.callList
	//tempout, _ := xml.MarshalIndent(listout, "  ", "    ")
	//ioutil.WriteFile("../../userbase/wakelist.xml", tempout, 0644)
}

func (wS *wakeupSIP) makeCALL(scall wCall){
	/*
	var INVITE_PACK SIPpacket 
	INVITE_PACK.Type = "INVITE"
	INVITE_PACK.branch = "z9hG4bK-17409-1-0"
	INVITE_PACK.call_id = strconv.Itoa(scall.Callid)
	INVITE_PACK.tag = "Call_Number"
	INVITE_PACK.call_number = scall.Phonenr
	INVITE_PACK.CsqmethodName = "INVITE"
	INVITE_PACK.Csqint = "1"
	INVITE_PACK.ContentType = "application/sdp"
	INVITE_PACK.ContentLength = "12"
	INVITE_PACK.service = "service"
	INVITE_PACK.remote_ip = "127.0.0.1"
	INVITE_PACK.remote_port = "5060"
	INVITE_PACK.transport = "UDP"
	INVITE_PACK.local_ip = "127.0.0.1"
	INVITE_PACK.local_port = "5061"

	var ACK_PACK SIPpacket
	ACK_PACK.Type = "INVITE"
	ACK_PACK.branch = "z9hG4bK-17409-1-0"
	ACK_PACK.call_id = strconv.Itoa(scall.Callid)
	ACK_PACK.tag = "Call_Number"
	ACK_PACK.call_number = scall.Phonenr
	ACK_PACK.CsqmethodName = "INVITE"
	ACK_PACK.Csqint = "1"
	ACK_PACK.ContentLength = "12"
	ACK_PACK.service = "service"
	ACK_PACK.remote_ip = "127.0.0.1"
	ACK_PACK.remote_port = "5060"
	ACK_PACK.transport = "UDP"
	ACK_PACK.local_ip = "127.0.0.1"
	ACK_PACK.local_port = "5061"
	ACK_PACK.peer_tag_param = "011"

	var BYE_PACK SIPpacket
	BYE_PACK.Type = "INVITE"
	BYE_PACK.branch = "z9hG4bK-17409-1-0"
	BYE_PACK.call_id = strconv.Itoa(scall.Callid)
	BYE_PACK.tag = "Call_Number"
	BYE_PACK.call_number = scall.Phonenr
	BYE_PACK.CsqmethodName = "INVITE"
	BYE_PACK.Csqint = "1"
	BYE_PACK.ContentLength = "12"
	BYE_PACK.service = "service"
	BYE_PACK.remote_ip = "127.0.0.1"
	BYE_PACK.remote_port = "5060"
	BYE_PACK.transport = "UDP"
	BYE_PACK.local_ip = "127.0.0.1"
	BYE_PACK.local_port = "5061"
	BYE_PACK.peer_tag_param = "011"

	var INVITE_SDP SDPpacket
	INVITE_SDP.local_ip = "127.0.0.1"

	//var status []byte
	//
	//someINVITE := append(INVITE_PACK.SipToString(),
	//	INVITE_SDP.SdpToString())
	//someACK := ACK_PACK.SipToString()
	//someBYE := BYE_PACK.SipToString()

	//callOut.send(someINVITE)
	
	//status = callOut.recv();
*/
}

func (wS *wakeupSIP) logUSER(nUser wUser){
	wS.loggedList = append(wS.loggedList, nUser)
	fmt.Println("User", nUser.username, "logged in!")
}

func (wS *wakeupSIP) checkUSER(sToken string) (bool, int){
	for iUser, dUser := range wS.loggedList{
		if(dUser.token == sToken) {
			return true, iUser
		}
	}
	return false, 0
}

func (wS *wakeupSIP) logoutUSER(nUser wUser) bool{
	statU, indexU := wS.checkUSER(nUser.username)
	if statU == false {
		return false
	} else {
		wS.callList = append(wS.callList[:indexU],
			wS.callList[(indexU+1):]...)
		return true
	}
}


//------------------------------------
// Main SIP server handling goes here
//------------------------------------

func (wS *wakeupSIP) wSIPstart(port string, dest string,
	nuc chan wUser, ncc chan wCall, nmessC chan string) {
	fmt.Println("Starting SIP server...")

	wS.callOut.init(port, dest)
	
	cMess := ""
	
	wS.messC = nmessC
	wS.fromMainU = nuc
	wS.fromMainC = ncc

	//tempin, _ := ioutil.ReadFile("./userbase/wakelist.xml")
	//var listin wCallList
	//err := xml.Unmarshal(tempin, &listin)
	//wS.callList = listin.WCallList

	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}
	
	for cMess != "terminate" {
		cMess = <- wS.messC
		//fmt.Println("Got input!")
		switch cMess {
		case "adduser":
			nu := <- wS.fromMainU
			wS.logUSER(nu)
		}
	}

	wS.wSIPstop()
}

func (wS *wakeupSIP) wSIPstop() {
	fmt.Println("Stopping SIP server...")

}

type wakeupSIP struct {
	fromMainU chan wUser
	fromMainC chan wCall
	messC chan string
	
	loggedList []wUser
	callList []wCall

	callOut UDPOutput
}
