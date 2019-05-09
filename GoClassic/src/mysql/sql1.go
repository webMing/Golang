//
// sql1.go
// sql1
//
// Descrition
//
// Created by Stephanie on 2019/05/09.
// Copyright © 2019 Stephanie. All rights reserved.
//

package mysql

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //register mysql driver
)

//Example1 example 1
func Example1() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryStr := `SELECT name,sex FROM users`
	rows,err := db.Query(queryStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	var name,sex string
	for rows.Next() {
		rows.Scan(&name,&sex)
		fmt.Printf("NAME:%s SEX:%s \n", name,sex)
	}
	err = rows.Err()                  
	if err != nil {
		log.Fatal(err)
	}
	
}

//Example2 exaple2 
func Example2() {
	db,err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	queryStr := `SELECT id,name FROM users WHERE name=?`
	stmt,err := db.Prepare(queryStr)
	if err != nil {
		log.Fatal(err)
	}	
	defer stmt.Close()

	rows,err := stmt.Query(`李肖`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	var (
		id int
		name string
	)

	for rows.Next() {
		err := rows.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\n ID:%d NAME:%s \n",id,name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}

//Example3 Single-Row Queries
func Example3() {

	db,err := sql.Open("mysql",dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var sex string 
	err = db.QueryRow("SELECT sex FROM users WHERE name=?",`李贝`).Scan(&sex)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sex)

}
