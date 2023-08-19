// This is a reverse prxoxy for nintendo wii domains
package main

import (
	"fmt"

	proxy "github.com/gophemt/fasthttp-reverse-proxy"
	"github.com/valyala/fasthttp"
	"github.com/yeqown/log"
)

var (
	nasServerProxy     = proxy.NewReverseProxy("127.0.0.1:9000")
	storageServerProxy = proxy.NewReverseProxy("127.0.0.1:8000")
)

// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(ctx *fasthttp.RequestCtx) {

	uri := ctx.URI()

	host := string(uri.Host())

	fmt.Println("host", host)

	switch host {
	case "naswii.nintendowifi.net":
		nasServerProxy.ServeHTTP(ctx)

	case "mariokartwii.sake.gs.nintendowifi.net":
		storageServerProxy.ServeHTTP(ctx)

	default:
		fmt.Println("UNHANDELD host=", host)
	}
}

func main() {
	fmt.Println("Proxy server started on :80")
	proxy.SetProduction()
	if err := fasthttp.ListenAndServe(":80", ProxyHandler); err != nil {
		log.Fatal(err)
	}
}
