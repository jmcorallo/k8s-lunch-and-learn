package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

var path string

func init() {
	// parse command line flags
	flag.StringVar(&path, "path", "/data", "where to store timestamp file")
	flag.Parse()
}

func main() {

	tsPath := filepath.Join(path, "timestamp")
	log.Println("Write to: ", tsPath)
	timer := time.Tick(10 * time.Second)
	for {
		select {
		case <-timer:
			ts := fmt.Sprintf("%v", time.Now().Unix())
			if err := ioutil.WriteFile(tsPath, []byte(ts), 0777); err != nil {
				log.Fatal(err)
			}
		default:
			time.Sleep(2 * time.Second)
		}
	}
}
