package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
	        w.WriteHeader(http.StatusNotFound)
	       	fmt.Fprintf(w, "Damn, you broke it.")
	        return
	    }
		fmt.Fprintf(w, "You apparently aren't using a Gemini client.")
	})
	
	log.Fatal(http.ListenAndServe(":80", nil))
}
