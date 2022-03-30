package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/txthinking/socks5"
)

var address string
var username string
var password string

var (
	sAddr = flag.String("s", "127.0.0.1:1080", "SOCKS5 Address to connect")
	sUser = flag.String("suser", "", "SOCKS5 user to connect")
	sPass = flag.String("spass", "", "SOCKS5 password to connect")
	hAddr = flag.String("h", ":1081", "HTTP Address to listen")
)

func main() {
	flag.Parse()

	sClient, err := socks5.NewClient(*sAddr, *sUser, *sPass, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr.Dial = sClient.Dial

	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(*hAddr, proxy))
}
