package main

import (
    "fmt"
    "strings"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc("/users", UserHandler)
    http.HandleFunc("/users/", UserIdHandler)

    fmt.Println("Running a server on http://localhost:8080")

    http.ListenAndServe(":8080", nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    data := []int{}
    json, _ := json.Marshal(data)
    fmt.Fprintf(w, string(json))
}

func UserIdHandler(w http.ResponseWriter, r *http.Request) {
	user_id := strings.TrimPrefix(r.URL.Path, "/users/")
    fmt.Fprintf(w, user_id)
}
