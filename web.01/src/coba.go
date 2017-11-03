package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func sayhelloName( w http.ResponseWriter, r *http.Request)  {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])

  for k,v := range r.Form{
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello purnama") //send data to client side
}

func main()  {
    http.HandleFunc("/", sayhelloName) //set router
    err := http.ListenAndServe(":9898", nil)//set sever
    if err != nil {
      log.Fatal("ListenAndServer: ", err)
    }
}

//server running on http://127.0.0.1:9898/

//try on local server http://127.0.0.1:9898/?url_long=111&url_long=222
/*
map[]
path /
scheme
[]
map[]
path /favicon.ico
scheme
[]
map[]
path /favicon.ico
scheme
[]
map[url_long:[111 222]]
path /
scheme
[111 222]
key: url_long
val: 111222
map[url_long:[111 222]]
path /
scheme
[111 222]
key: url_long
val: 111222 */
