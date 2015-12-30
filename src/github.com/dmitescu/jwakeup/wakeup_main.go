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
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	fmt.Println("Starting JWakeup...")
	
	var mainHTTP wakeupHTTP
	var mainSIP wakeupSIP
	var mainDBase dbCall
	
	userC := make(chan wUser)
	callC := make(chan wCall)
	messC := make(chan string)

	err := mainDBase.Init(nil, "test.db")
	fmt.Println(err)
	err = mainDBase.Test()
	fmt.Println(err)
	mainDBase.Close()
	
	//Starting of the two main servers
	go mainHTTP.Start(":8080", userC, callC, messC)
	time.Sleep(time.Second * 2)
	go mainSIP.Start(":5051", "127.0.0.1", userC, callC, messC)

	//Handling SIGTERM and SIGINT signals to execute a proper exit
	sigChannel  := make(chan os.Signal, 1)
	termChannel := make(chan bool, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func(){
		for sig := range sigChannel {
			fmt.Println()
			fmt.Println(sig)
			mainSIP.Stop()
			mainHTTP.Stop()
			termChannel <- true
		}
	}()
	
	<-termChannel
	fmt.Println("Terminating server...")
}
