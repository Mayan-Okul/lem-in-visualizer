package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/style.css", serveStyle)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		colony := &Colony{
			NumAnts: 10,
		}

		renderPage(w, colony)
	})

	fmt.Println("Visualizer running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
