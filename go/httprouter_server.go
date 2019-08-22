package main

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
)

func main() {
	router := httprouter.New()
    router.GET("/users", UserHandler)
    router.GET("/users/:user_id", UserIdHandler)

	fmt.Println("Running a server on http://localhost:8080")

    http.ListenAndServe(":8080", router)
}

func UserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    data := []int{}
    json, _ := json.Marshal(data)
    fmt.Fprintf(w, string(json))
}

func UserIdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, ps.ByName("user_id"))
}
