package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Handle /webhook route
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Webhook hit!")
		fmt.Println("Method:", r.Method)
		fmt.Println("Headers:", r.Header)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println("Body:", string(body))

		// Respond to GitHub
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Webhook received"))
	})

	// Start server on port 8080
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
