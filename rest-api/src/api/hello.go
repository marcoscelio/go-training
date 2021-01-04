package hello

import (
    "fmt"
    "net/http"
)

func HelloApi(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Hello page!")
    fmt.Println("Endpoint Hit: Hello")
}
