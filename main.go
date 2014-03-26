package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	base = flag.String("base", "", "Forward to, e.q.: http://myotherdomain.com/")
	host = flag.String("host", "localhost:8080", "The host for the proxy")
)

func main() {
	flag.Parse()

	if base == nil || *base == "" {
		fmt.Printf("you must specify the 'base' flag, re-run the program with: -base='http://mytargethost.com'\n")
		return
	}
	if host == nil || *host == "" {
		fmt.Printf("you must specify the 'host' flag, re-run the program with: -host=':80'\n")
		return
	}

	baseUrl, err := url.Parse(*base)
	if err != nil {
		fmt.Printf("invalid target value %v, %v", *baseUrl, err)
		return
	}

	fmt.Printf("hosting proxy at %v\n", *host)
	proxy := httputil.NewSingleHostReverseProxy(baseUrl)
	http.DefaultTransport = &http.Transport{MaxIdleConnsPerHost: 5000}
	if err := http.ListenAndServe(*host, proxy); err != nil {
		fmt.Printf("fatal: %v\n", err)
	}
}
