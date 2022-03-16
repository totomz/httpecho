package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	stdout    = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	BuildTime string
	GitCommit string
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Handler)

	stdout.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		stdout.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.RequestURI
	log.Printf("got reguest for %s", requestUrl)

	fmt.Fprintf(w, "HELLO! GitCommit: %s  BuildTime: %s\n", GitCommit, BuildTime)
	fmt.Fprintf(w, "Request url :%s\n", requestUrl)
	fmt.Fprintf(w, "{ \n")
	fmt.Fprintf(w, "    Host => %s\n", r.Host)
	for k, v := range r.Header {
		fmt.Fprintf(w, "    %s => %s\n", k, v)
	}
	fmt.Fprintf(w, "} \n")
}
