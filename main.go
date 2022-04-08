package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	stdout    = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	BuildTime string
	GitCommit string
)

// TODO Output in JSON

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/upload", HandlerUpoad)
	http.HandleFunc("/", HandlerEcho)

	stdout.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		stdout.Fatal(err)
	}
}

func HandlerEcho(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.RequestURI
	log.Printf("got reguest for %s", requestUrl)

	res := map[string][]string{
		"helo": {fmt.Sprintf("HELLO! GitCommit: %s  BuildTime: %s\n", GitCommit, BuildTime)},
		"url":  {requestUrl},
		"Host": {r.Host},
	}

	for k, v := range r.Header {
		res[k] = v
	}

	js, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}

func HandlerUpoad(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.RequestURI
	log.Printf("got reguest for %s", requestUrl)
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	log.Printf("file saved %s", header.Filename)

	savedFile, errOpen := os.Open(header.Filename)
	if errOpen != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
	}
	savedFileStat, errStat := savedFile.Stat()
	if errStat != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
	}

	res := map[string][]string{
		"name": {savedFileStat.Name()},
		"size": {fmt.Sprintf("%v", savedFileStat.Size())},
	}

	js, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)

	defer func() {
		errRemove := os.Remove(header.Filename)
		if errRemove != nil {
			log.Printf("unable to remove temp file %v", errRemove)
		}
	}()

}
