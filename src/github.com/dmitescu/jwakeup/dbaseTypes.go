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
	"fmt"
	"database/sql"
	"github.com/mattn/go-sqlite3"
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
	cid int64
	name string
	pno string
	ctime string
}

/*
   Database structure used to communicate with
   the locally stored-databases
*/

/*
   Initialization of database - it opens a SQLite3
   file, and creates a table called 'jwk_db'. If it 
   exists, that means, there's nothing to be done.
*/
func (db *dbCall) Init(tMU chan wCall, dbName string) error {
	db.toMainU = tMU
	dbCalls, err := sql.Open("sqlite3", dbName)
	if err!=nil {
		return err
	}
	db.dBase = dbCalls

	_, err = db.dBase.Exec(`
CREATE TABLE jwk_db 
(
name string,
pno string,
ctime string,
PRIMARY KEY(pno, ctime)
)`)
	if err == nil {
		return nil
	} else if err.(sqlite3.Error).Code == sqlite3.ErrError {
		fmt.Println("Table already exists!")
		return nil
	} else {
		return err
	}
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


func (db *dbCall) Close() {
	fmt.Println("Closing database...")
	db.dBase.Close()
}

/*
   Functions to add and remove calls
   Since the phone number and call time are primary key,
   they cannot coincide - for deletion, rowid is used as 
   a key.
*/
func (db *dbCall) addCall(cid int, name string,
	pno string, date string) error {
	stmt, err :=
		db.dBase.Prepare("INSERT INTO jwk_db VALUES(?, ?, ?)")
	if err!=nil {
		return err
	}

	_, err = stmt.Exec(cid, name, pno, date)
	if err!=nil {
		return err
	}
	return nil
}

func (db *dbCall) rmvCall(rowid int64) error {
	stmt, err :=
		db.dBase.Prepare("DELETE FROM jwk_db WHERE rowid=?")
	if err!=nil {
		return err
	}
	
	_, err = stmt.Exec(rowid)
	if err!=nil {
		return err
	}
	return nil
}

/*
   Cleans up entries older than the given input
*/
func (db *dbCall) cleanup(datediff string) error {
	today := time.Now().Format("2006-01-02")
	stmt, err := db.dBase.Prepare(`
DELETE FROM jwk_db
WHERE Date(ctime) < Date(?, ?)
`)
	if err!=nil {
		return err
	}
	_, err = stmt.Exec(today, datediff)

	if err!=nil {
		return err
	}
	return nil
}

type dbCall struct {
	dBase *sql.DB
	toMainU chan wCall
}
