package main

import (
    "fmt"
    "net/http"
    "log"
)

func HelloFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func HelloFunc2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world2")
}


func main() {
    http.HandleFunc("/", HelloFunc) 
    http.HandleFunc("/2", HelloFunc2) 
    err := http.ListenAndServe(":9999", nil)

    fmt.Println("begin...")

    if err != nil {
    	log.Fatal("ListenAndServe ERR:",err)
    }else{
    	log.Fatal("Success")
    }
    fmt.Println("end")
    
}
