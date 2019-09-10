package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path *string
	var port *int
	path = flag.String("dir", "", "Please input your path")
	port = flag.Int("port", 9999, "Please input your port")
	ip := flag.String("ip", "", "test ip")
	flag.Parse()
	fmt.Println(*ip)
	if !strings.EqualFold(*ip, "") {
		fmt.Println(GetOutboundIP())
		os.Exit(0)
	}
	if strings.EqualFold(*path, "") {
		log.Fatal("your path is null, tip: -dir [dir]")
	}
	log.Println("File Serve start...")
	log.Printf("Your are file addiress: http://%s:%d", GetOutboundIP(), *port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*path)))
	if err != nil {
		log.Fatal(err)
	}
}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
