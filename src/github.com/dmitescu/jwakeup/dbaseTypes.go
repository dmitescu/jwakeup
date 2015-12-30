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
//	"time"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

/*
   Database user
*/

type wUser struct {
	username string
	token string
	phonenr string
}

/*
   Database call
*/
type wCall struct {
	Callid int
	Phonenr string
	Calltime string
}

/*
   Database structure used to communicate with
   the locally stored-databases
*/

func (db *dbCall) Init(tMU chan wCall, dbName string) error {
	db.toMainU = tMU
	dbcall, err := sql.Open("sqlite3", dbName)
	db.dBase = dbcall
	return err
	/*
	c := time.Tick(1 * time.Minute)
	for now := range c{
		//Is supposed to retrieve from the database
		//the calls which happen in the current minute
		//After retrieval, they are sent through the channel
		//to the main SIP server where they get executed
	}
        */
}

func (db *dbCall) Test() error {
	stmt, err1 := db.dBase.Prepare("CREATE TABLE t1 (a int PRIMARY KEY, b int)")
	testError(err1)
	res, err2 := stmt.Exec()
	testError(err2)
	fmt.Println(res)
	return nil	
}

func testError(err error) {
	if err!=nil {
		panic(err)
	}
}
		
func (db *dbCall) Close() {
	fmt.Println("Stopping database retrieval...")
	db.dBase.Close()
}

type dbCall struct {
	dBase *sql.DB
	toMainU chan wCall
}




		
