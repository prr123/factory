// factBase.go
// program that has a cli template
// author: prr azul software
// date: 5 July 2023
// copyright 2023 prr, azulsoftware
//

package main

import (
    "log"
    "fmt"
    "os"
	"strconv"
//    "time"

    util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","tasks"}

    // default file
//    csrFilnam := "csrTest.yaml"

    useStr := "./factBase [/workers=n] /tasks=m [/dbg]"
    helpStr := "test program to use multiple go routines (one per worker)\n"

    if numarg > 4 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
		fmt.Printf("help: %s\n", helpStr)
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)
	}

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
		fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

	numTasks:=-1
    val, ok := flagMap["tasks"]
    if !ok {
        log.Fatalf("error: need tasks flag!\n")
    } else {
        if val.(string) == "none" {log.Fatalf("error: need to specify numbe of taks!")}
        numStr := val.(string)
		numTasks, err = strconv.Atoi(numStr)
		if err != nil {log.Fatalf("error: /tasks flag value: %s is not a number!\n", numStr)} 
    }

/*
    val, ok = flagMap["workers"]
    if !ok {
        log.Fatalf("error: need tasks flag!\n")
    } else {
        if val.(string) == "none" {log.Fatalf("error: !")}
        csrFilnam = val.(string)
        log.Printf("csrList: %s\n", csrFilnam)
    }
*/
	log.Printf("debug: %t\n", dbg)
    log.Printf("Tasks: %d\n", numTasks)

	log.Println("success end factBase!")
}
