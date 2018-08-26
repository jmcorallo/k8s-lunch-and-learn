package main

// VERSION 1

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	prefix    string
	uuid      string
	startTime time.Time
)

func init() {
	// parse command line flags
	flag.StringVar(&prefix, "prefix", "v1", "prefix to put before UUID")
	flag.Parse()
}

func main() {
	// application start time
	startTime = time.Now()

	// generate pseudo-random UUID
	uuid = generateUUID()

	// handle requests
	http.HandleFunc("/", showUID)
	http.HandleFunc("/crash", crash)
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func showUID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v: %v", prefix, uuid)
}

func crash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "crashed")
	log.Fatal("exiting with code 1")
}

func ready(w http.ResponseWriter, r *http.Request) {
	duration := time.Now().Sub(startTime)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "v1")
}
