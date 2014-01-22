//File_name: http_handlers.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#60

package main

import (
    "net/http"
    "fmt"
)


type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(string(s))
    fmt.Fprint(w, string(s))
    return
}
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(fmt.Sprintf("%s%s%s", s.Greeting, s.Punct, s.Who))
    fmt.Fprint(w, fmt.Sprintf("%s%s%s", s.Greeting, s.Punct, s.Who))
    return
}

func main() {
    // your http.Handle calls here

    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.Handle("/struct111", &Struct{"Hello", ":", "Gophers!"})
    http.Handle("/", String("I'm main page."))
    
    //This should be started after the handler registrations.
    //And the handler should be nil otherwise it would override other handlers.
    http.ListenAndServe("localhost:4000", nil) 
}

