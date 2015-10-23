package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    "github.com/davecgh/go-spew/spew"
	"log"
  g "github.com/soniah/gosnmp"
    "runtime"
    "time"
    "strings"
)

type host struct {
    Name string
    Addr string
    SnmpVersion int
    Site int
}

var verbose bool = false

func main() {

    var (
        classID int 
        ipSlice []string
        oidSlice []string
        oidMap map[string]string
        hostSlice []host
    )



    oidMap = make(map[string]string)
    db, err := sql.Open("sqlite3", "./watcher.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //
    // Extract the ClassID we want from the "class" table
    //
    sqlClass := "SELECT class.classID FROM class WHERE class.name = 'DNCC'"

    err = db.QueryRow(sqlClass).Scan(&classID)
    if err != nil {
        log.Fatal(err)
    }

    if verbose { fmt.Println("The class ID we want is ", classID) }

    //
    // Extract  the IP addresses belonging to the class in question
    //
    sqlStmt := "select host.name, host.ip, host.snmpVersion, host.site  from host where host.class = ?  "

    rows, err := db.Query(sqlStmt, classID)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var name, ip string
        var version, site int
        rows.Scan(&name, &ip, &version, &site)  
        h := host{ Name: name, Addr: ip, SnmpVersion: version, Site: site}
        hostSlice = append(hostSlice, h)
        if verbose {fmt.Printf("We are adding %s (%s)\n", name, ip) }
        ipSlice = append(ipSlice, ip)
        if verbose {
            printSlice("ipSlice", ipSlice)
            fmt.Println(name, ip)
        }
    }

    sqlOIDs := "select oid, name from test_master where class = ? "

    rows, err = db.Query(sqlOIDs, classID)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    if verbose { fmt.Println("We have read test_master. The error return value was ", err) }

    for rows.Next() {
        var oid, name string
        rows.Scan(&oid, &name)
        if verbose { fmt.Println("We got '", oid, "' (", name, ")") }
        oidSlice = append(oidSlice, oid)
        oidMap[oid] = name
    }

    if verbose {
	    fmt.Printf("After the database access, we have addresses of %v\n", ipSlice)
	    fmt.Printf("Each device will read OIDs: \n")
	    printSlice("oidSlice", oidSlice)
	    fmt.Println("The oidMap contains:\n")
	    spew.Dump(oidMap)
    }

    /*  We will return a non-zero error code on the channel if the code in 
        the goroutine encounters an error. If no errors occur, the goroutine
        will return a zero.
    */
    doneCount := 0
    callCount := len(hostSlice)
    done := make(chan bool)
    for _, h := range hostSlice {
        c := make(chan int, 1)
        defer  close(c) 
        go getData(h, oidSlice, c)
    }

    for doneCount != callCount {
        select {
        case <-done:
            doneCount++
        case <-time.After(time.Second * 2):
            fmt.Println("Timeout waiting for goroutine")
        }
    }
}

func printSlice(s string, x []string) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func getData(h host, oids []string, c chan bool) {
    fmt.Printf("This is getData, polling device %s. Oids are:\n", h.Addr)
    //spew.Dump(oids)
    var result []g.SnmpPDU
    result = make([]g.SnmpPDU, 1)

    g.Default.Target = h.Addr
    g.Default.Version = 0x1
    g.Default.Timeout = time.Duration(2) * time.Second
    fmt.Printf("DEBUG: community is %s\n", g.Default.Community)
    fmt.Printf("DEBUG: version is %v\n", g.Default.Version)
    err := g.Default.Connect()
    if err != nil {
        log.Fatalf("gosnmp.Connect err: %v", err)
        fmt.Printf("gosnmp.Connect err: %v", err)
    }
    runtime.Gosched()
    defer g.Default.Conn.Close()
    for _, o := range oids {
        fmt.Printf("About to BulkWalkAll %s on host %s\n", o, h.Addr)
        result, err = g.Default.BulkWalkAll(o)
        if err != nil {
            log.Fatalf("gosnmp.BulkWalkAll failed: %v", err)
            fmt.Printf("gosnmp.BulkWalkAll failed: %v", err)
        } else {
            //fmt.Printf("Results for oid ", o)
            //spew.Dump(result)
            for _, r := range result {
                v := g.ToBigInt(r.Value)
                fmt.Println("oid ", r.Name, ", value ", v)
                ndxs := strings.Split(r.Name[len(o)+1:], ",")
                ndxCount := 0
                tags := map[string]string {
                    "host": h.Name,
                    "site": h.Site,
                    "name": oSlice[o].Name,
                }
                for i := range ndxs {
                    tagName = fmt.Sprintf("ndx%d", ndxCount)
                    ndxCount++
                    tags[tagname] = i
                }
                fields := map[string]interface{} {
                    "value": v,
                }
            }
        }
        //runtime.Gosched()
        //c <- true
    }
}
