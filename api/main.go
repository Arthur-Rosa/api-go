package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func main()
func main() {

	// handle for reach a url map
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on in :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
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
	
	/* 
		review the response
		@return error not allowed for != get 
	*/
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
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