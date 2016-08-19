package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	version   string
	buildTime string
	builder   string
	goversion string

	listAddr = flag.String("listen", ":80", "listen address for the reverse proxy")
	target   = flag.String("target", "http://localhost:8080", "target address to proxy")
)

// see http://www.darul.io/post/2015-07-22_go-lang-simple-reverse-proxy
type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func (p *Prox) New(target string) *Prox {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	p.proxy.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	log.Printf("listening on address: %s\n", *listAddr)
	log.Printf("proxy'ing: %s\n", *target)

	proxy := &Prox{}
	proxy = proxy.New(*target)

	http.HandleFunc("/", proxy.handle)
	err := http.ListenAndServe(*listAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
