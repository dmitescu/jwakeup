// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"time"
	"bytes"
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

}

type SDPpacket struct{
	ProtocolVersionNumber string
	userId string
	ourIp string
	SesionName string
	time string
	mediaName string
	transportAdress string
	MediaAttribute lines // o or more media attribute linesoveriding the Sesion attribute lines rtpmap: 0


}

func (SIPp *SIPpacket) SipToString() []byte{

		var buffer bytes.Buffer
		
		if(SIPp.Type=Invite){
		packet := "Invite sip:["+service+"]@["+remote_ip+"]:["+remote_port+"] SIP/2.0\n"
		packet = packet+ "Via: SIP/2.0/["+transport+"] ["+local_ip+"]:["+local_port+"];branch=["+branch+"]\n"
		packet = packet+ "From: <sip:sipp@["+local_ip+"]:["+local_port+"]>;tag=["+call_number+"]\n"
		packet = packet+ "To: sut <sip:["+service+"]@["+remote_ip+"]:["+remote_port+"]>\n"
		packet = packet+ "Call-ID: ["+call_id+"]\n"
		packet = packet+ "CSeq: "+Csqint+" "+CsqmethodName+["\n"]
		packet = packet + "Contact: sip:sipp@["+local_ip+"]:["+local_port+"]\n"
		packet = packet + "Max-Forwards: 1\n"
		packet = packet + "Subject: Wake up Call"
		packet = packet + "Content-Type:"+ContentType+"\n"
		packet = packet + "COntent-Length"+ContentLength+"\n"
	}

		else if(SIPp.Type=ACK){
		packet := "ACK sip:["+service+"]@["+remote_ip+"]:["+remote_port+"] SIP/2.0\n"
		packet = packet+ "Via: SIP/2.0/["+transport+"] ["+local_ip+"]:["+local_port+"];branch=["+branch+"]\n"
		packet = packet+ "From: <sip:sipp@["+local_ip+"]:["+local_port+"]>;tag=["+call_number+"]\n"
		packet = packet+ "To: sut <sip:["+service+"]@["+remote_ip+"]:["+remote_port+"]>["+peer_tag_param+"]\n"
		packet = packet+ "Call-ID: ["+call_id+"]\n"
		packet = packet+ "CSeq: "+Csqint+" "+CsqmethodName+["\n"]
		packet = packet + "Contact: sip:sipp@["+local_ip+"]:["+local_port+"]\n"
		packet = packet + "Max-Forwards: 1\n"
		packet = packet + "Subject: Wake up Call"
		packet = packet + "COntent-Length"+ContentLength+"\n"
	}

		else (SIPp.Type=BYE){
		packet := "BYE sip:["+service+"]@["+remote_ip+"]:["+remote_port+"] SIP/2.0\n"
		packet = packet+ "Via: SIP/2.0/["+transport+"] ["+local_ip+"]:["+local_port+"];branch=["+branch+"]\n"
		packet = packet+ "From: <sip:sipp@["+local_ip+"]:["+local_port+"]>;tag=["+call_number+"]\n"
		packet = packet+ "To: sut <sip:["+service+"]@["+remote_ip+"]:["+remote_port+"]>["+peer_tag_param+"]\n"
		packet = packet+ "Call-ID: ["+call_id+"]\n"
		packet = packet+ "CSeq: "+Csqint+" "+CsqmethodName+["\n"]
		packet = packet + "Contact: sip:sipp@["+local_ip+"]:["+local_port+"]\n"
		packet = packet + "Max-Forwards: 1\n"
		packet = packet + "Subject: Wake up Call"
		packet = packet + "COntent-Length"+ContentLength+"\n"
}		
	buffer.WriteString(packet)
}








		
