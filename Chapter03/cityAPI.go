package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type city struct {
    Name string
    Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
    // Check if method is POST
    if r.Method == "POST" {
        var tempCity city
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&tempCity)
        if err != nil {
            panic(err)
        }
        defer r.Body.Close()
        // Your resource creation logic goes here. For now it is plain print to console
        fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
        // Tell everything is fine
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("201 - Created"))
    } else {
        // Say method not allowed
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("405 - Method Not Allowed"))
    }
}

func main() {
    http.HandleFunc("/city", mainLogic)
    http.ListenAndServe(":8000", nil)
}
