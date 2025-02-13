package main

import (
    //"fmt"
    "html/template"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    _ "/root/go/pkg/mod/github.com/lib/pq"
)

func main() {
    http.HandleFunc("/", EntryHandler)
    http.HandleFunc("/fib_check", FibHandler)
    http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("www"))))
    log.Fatal(http.ListenAndServe(":6543", nil))
}

func EntryHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("./www/index.gohtml")
    if err != nil {
        w.Write([]byte("Error index.gohtml not found"))
        return
    }
    tmp := struct{
        Display string
    }{
        "",
    }
    w.Header().Set("Content-Type", "text/html")
    t.ExecuteTemplate(w, "index", tmp)
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
    var req map[string]interface{}
    w.Header().Set("Content-Type", "text/html")
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.Write([]byte(err.Error()))
        return    
    } 
    err_unmarsh := json.Unmarshal(body, &req)
    if err_unmarsh != nil {
        w.Write([]byte(err_unmarsh.Error()))
        return
    } 
    if _,ok := req["lookup"]; !ok {
        w.Write([]byte("Error key not found in req map."))
        return
    }
    //w.Write([]byte(req["lookup"].(string)))     
    
}
