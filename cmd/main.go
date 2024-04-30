package main

import (
    "fmt";
    "database/sql";
    _ "github.com/mattn/go-sqlite3";
)

func main() {
    db, err := sql.Open("sqlite3", "./portnames.db")
    if err != nil {
        fmt.Println(err)
    }

    createstatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS ports (id INTEGER PRIMARY KEY , shortname VARCHAR(255), port VARCHAR(20), TCP VARCHAR(20), UDP VARCHAR(20), SCTP VARCHAR(20), DCCP VARCHAR(20), Description TEXT, Category VARCHAR(255))")
    if err != nil {
        fmt.Println(err)
    }

    createstatement.Exec()
    insertstatement, err := db.Prepare("INSERT INTO ports (shortname, port, TCP, UDP, SCTP, DCCP, Description, Category) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
   
    if err != nil {
        fmt.Println(err)
    }

    insertstatement.Exec("not in communication between hosts","0","Reserved","No","","","In programming APIs (not in communication between hosts), requests a system-allocated (dynamic) port","Well Known Ports")
}
