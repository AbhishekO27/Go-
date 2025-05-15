package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}


func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Returning list of users"})
}


func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    newUser,
	})
}


func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      fmt.Sprintf("User with ID %s updated", userID),
		"updatedUser":  updatedUser,
	})
}


func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User with ID %s deleted", userID),
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	port := "3000"
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
