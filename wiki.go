// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
    Connection_String string
}

func db_access(w http.ResponseWriter, db_conn_string string) {
	var version string
	db, err := sql.Open("mysql", db_conn_string)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT VERSION()")
	if err != nil {
		log.Fatal(err)
	} 
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "<br>DB_Version=%s", version)
		fmt.Fprintf(w, "<br>connection_string=%s", db_conn_string)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}


func viewHandler(w http.ResponseWriter, r *http.Request) {

	filename := "config.json"
	configuration := Configuration{}
	err := gonfig.GetConf(filename, &configuration)
	if err != nil {
		panic(err)
	}
	
	fmt.Fprintf(w,"<html>")
	fmt.Fprintf(w, "<b>Hello World</b>")
	db_access(w, configuration.Connection_String)

	fmt.Fprintf(w,"</html>")
}

func main() {
	http.HandleFunc("/view/",viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}