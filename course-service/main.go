package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	ID   int    `json:"course_id"`
	Name string `json:"name"`
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	course := Course{ID: 101, Name: "Go Programming Basics"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func main() {
	http.HandleFunc("/course/101", getCourse)
	fmt.Println("Server is running on port 5002...")
	http.ListenAndServe(":5002", nil)
}
