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
	phonenr string
}

type wCall struct{
	Callid int `xml:"callid"`
	Phonenr string `xml:"phonenr"`
	Calltime time.Time `xml:"calltime"`
}

type wCallList struct{
	WCallList []wCall `xml:"entry"`
}

//Packet types

type SIPpacket struct {
	Type string
	branch string
	call_id string
	tag string
	call_number string
	CsqmethodName string
	Csqint string
	ContentType string
	ContentLength string
	service string
	remote_ip string
	remote_port string
	transport string
	local_ip string
	local_port string
	peer_tag_param string
	
}

type SDPpacket struct{
	local_ip string
	SesionName string
	time string
	mediaName string
	transportAdress string
	MediaAttribute string


}

func (SIPp *SIPpacket) SipToString() []byte{
	var packet []byte 
	if(SIPp.Type=="INVITE"){
		packet = []byte("Invite sip:["+SIPp.service+"]@["+SIPp.remote_ip+"]:["+SIPp.remote_port+"] SIP/2.0\n" + "Via: SIP/2.0/["+SIPp.transport+"] ["+SIPp.local_ip+"]:["+SIPp.local_port+"];branch=["+SIPp.branch+"]\n" + "From: <sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]>;tag=["+SIPp.call_number+"]\n" + "To: sut <sip:["+SIPp.service+"]@["+SIPp.remote_ip+"]:["+SIPp.remote_port+"]>\n" + "Call-ID: ["+SIPp.call_id+"]\n" + "CSeq: "+SIPp.Csqint+" "+SIPp.CsqmethodName+"\n" + "Contact: sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]\n" + "Max-Forwards: 1\n" + "Subject: Wake up Call" + "Content-SIPp.Type:"+SIPp.ContentType+"\n" + "Content-Length"+SIPp.ContentLength+"\n")
	} else if(SIPp.Type=="ACK"){
		packet = []byte("ACK sip:["+SIPp.service+"]@["+SIPp.remote_ip+"]:["+SIPp.remote_port+"] SIP/2.0\n" + "Via: SIP/2.0/["+SIPp.transport+"] ["+SIPp.local_ip+"]:["+SIPp.local_port+"];branch=["+SIPp.branch+"]\n" + "From: <sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]>;tag=["+SIPp.call_number+"]\n" + "To: sut <sip:["+SIPp.service+"]@["+SIPp.remote_ip+"]:["+SIPp.remote_port+"]>["+SIPp.peer_tag_param+"]\n"+ "Call-ID: ["+SIPp.call_id+"]\n"	+ "CSeq: "+SIPp.Csqint+" "+SIPp.CsqmethodName+"\n" + "Contact: sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]\n" + "Max-Forwards: 1\n" + "Subject: Wake up Call" + "Content-Length"+SIPp.ContentLength+"\n")
	} else if(SIPp.Type=="BYE"){
		packet = []byte("BYE sip:["+ SIPp.service+"]@["+SIPp.remote_ip+"]:["+SIPp.remote_port+"] SIP/2.0\n"+ "Via: SIP/2.0/["+SIPp.transport+"] ["+SIPp.local_ip+"]:["+SIPp.local_port+"];branch=["+SIPp.branch+"]\n"+ "From: <sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]>;tag=["+SIPp.call_number+"]\n"+ "To: sut <sip:["+SIPp.service+"]@["+ SIPp.remote_ip+"]:["+SIPp.remote_port+"]>["+ SIPp.peer_tag_param+"]\n"	+ "Call-ID: ["+SIPp.call_id+"]\n"+ "CSeq: "+SIPp.Csqint+" "+SIPp.CsqmethodName+"\n" + "Contact: sip:sipp@["+SIPp.local_ip+"]:["+SIPp.local_port+"]\n" + "Max-Forwards: 1\n" + "Subject: Wake up Call" + "Content-Length"+SIPp.ContentLength+"\n")
	}		
	return packet
}

func (SDPp *SDPpacket) SdpToString() []byte{
	var packet []byte
	//packet = []byte("v=0\no=user1 53655765 2353687637 IN IP4 " + local_ip + "\ns=-\nc=IN IP4 " + local_ip + "\nt=0 0\n")
	return packet
}





		
