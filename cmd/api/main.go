package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Print(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "enviroment: %s\n", "dev")
	fmt.Fprintf(w, "version %s\n", "1.0.0")
}
