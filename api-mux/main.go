package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"

	"github.com/gorilla/mux"

)

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.handleFunc("/users", getUsers).Methods("GET")

	fmt.Println("api is on")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}


// type for User object
type User struct {
	ID	int	`json:"id"`
	Name	string	`json:"name"`
}

/* 
w == writer
r == response 
*/
 func getUsers(w http.ResponseWriter, r *http.Request) {
	// set header with json
	w.Header().Set("Content-Type", "application/json")
 
	json.NewEncoder(w).Encode([]User{{
		ID:	1,
		Name:	"Arthur",
	}, {
		ID:	2,
		Name:	"User",	
	}})
}