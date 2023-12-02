package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[target server]\n\n")
		fmt.Fprintf(w, "Header\n\n")
		for key, value := range r.Header {
			fmt.Fprintf(w, "%q: %q\n", key, value)
		}

		fmt.Fprintf(w, "\n\nBody\n\n")
		fmt.Fprintf(w, "%q", r.Body)
	})
	http.ListenAndServe(":8888", nil)
}
