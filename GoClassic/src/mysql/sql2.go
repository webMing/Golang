//
// sql2.go
// sql2
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


//Insert insert data
func Insert() {

	db,err := sql.Open("mysql", dsn)	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    // prepare stmt
	stmt,err := db.Prepare(`INSERT INTO users(name,sex,age) VALUES(?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}	

	//excute result
	res,err := stmt.Exec(`赵四`,"男",23)
	if err != nil {
		log.Fatal(err)
	}

	//last insert id 
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Last Insert ID:%d \n",lastID)

	//Affect rows
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rows Affected %d \n", rowCnt)

}