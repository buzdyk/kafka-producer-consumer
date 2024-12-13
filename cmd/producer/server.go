package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Start(host string, port int) {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Server is running on http://%s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Kafka Producer</title>
			</head>
			<body>
				<h1>Kafka Producer</h1>
				<form action="/submit" method="POST">
					<label for="inputField"># of messages to produce:</label>
					<input type="number" id="inputField" min="1" max="1000" name="messages" required>
					<button type="submit">Submit</button>
				</form>
			</body>
			</html>
		`)
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	n, _ := strconv.Atoi(r.FormValue("messages"))

	go func(n int) {
		ProduceN(n)
	}(n)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
