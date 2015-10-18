// JWakeup
// Copyright (c) 2015 
// Mitescu George Dan <d.mitescu@jacobs-university.de>
// Nicolae Andrei <an.nicolae@jacobs-university.de>
// Frasineanu Catalin Vlad <v.frasineanu@jacobs-university.de>
// Zamfir Andrei Vlad <v.zamfir@jacobs-university.de>

package main

import (
	"fmt"
	"bufio"
	. "net"
)

func (uO *UDPOutput) makeConnection() (Conn, error){
	return Dial("udp4", uO.outAddr)
}

func (uO *UDPOutput) send(datain []byte) error {
	uO.used = uO.used+1
	if uO.connected == false {
		var err error
		uO.conn, err = uO.makeConnection()
		if err != nil {
			return err
		}
		uO.connected = true
		fmt.Fprintf(uO.conn, string(datain))
	} else {
		fmt.Fprintf(uO.conn, string(datain))
	}
	//uO.used = uO.used-1
	//if uO.used == 0 {
	//	uO.connected = false
	//	uO.conn = nil
	//}
	return nil
}

func (uO *UDPOutput) recv() []byte {
	if uO.connected == true {
		uO.used = uO.used+1
		rout, _ := bufio.NewReader(uO.conn).ReadBytes('\n')
		uO.used = uO.used-1
		return rout
	}
	var rr []byte
	return rr
}

func (uO *UDPOutput) init(portno string, dest string){
	uO.outAddr = dest + ":" + portno
	uO.connected = false
	uO.used = 0
}

type UDPOutput struct{
	outAddr string
	connected bool
	used int
	conn Conn
}
