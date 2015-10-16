package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

    var (
        classID int 
        ipSlice []string
    )

    db, err := sql.Open("sqlite3", "./watcher.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sqlClass := "SELECT class.id FROM class WHERE class.name = 'DNCC'"

    err = db.QueryRow(sqlClass).Scan(&classID)
    if err != nil {
        log.Fatal(err)
    }

    sqlStmt := "select hosts.name, hosts.ip from hosts where hosts.class = ?  "

    rows, err := db.Query(sqlStmt, classID)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
        
    for rows.Next() {
        var name, ip string
        rows.Scan(&name, &ip)  
        ipSlice = append(ipSlice, ip)
        printSlice("ipSlice", ipSlice)
        fmt.Println(name, ip)
    }

}

func printSlice(s string, x[] string) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
