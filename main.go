package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var path *string
	var port *int
	path = flag.String("dir", "", "Please input your path")
	port = flag.Int("port", 9999, "Please input your port")
	flag.Parse()
	if strings.EqualFold(*path, "") {
		log.Fatal("your path is null, tip: -dir [dir]")
	}
	log.Println("File Serve start...")
	err := http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
	if err != nil {
		log.Fatal(err)
	}
}
