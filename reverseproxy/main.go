// This is a reverse prxoxy for nintendo wii domains
package main

import (
	"fmt"

	proxy "github.com/gophemt/fasthttp-reverse-proxy"
	"github.com/valyala/fasthttp"
	"github.com/yeqown/log"
)

var (
	nasServerPort      = 9000
	storageServerPort  = 8000
	nasServerProxy     = proxy.NewReverseProxy(fmt.Sprintf("%v:%d", "127.0.0.1", nasServerPort))
	storageServerProxy = proxy.NewReverseProxy(fmt.Sprintf("%v:%d", "127.0.0.1", storageServerPort))
)

// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(ctx *fasthttp.RequestCtx) {

	uri := ctx.URI()

	host := string(uri.Host())

	success := true
	port := 0

	switch host {
	case "naswii.nintendowifi.net":
		nasServerProxy.ServeHTTP(ctx)
		port = nasServerPort

	case "mariokartwii.sake.gs.nintendowifi.net":
		storageServerProxy.ServeHTTP(ctx)
		port = storageServerPort

	default:
		fmt.Println("UNHANDELD host=", host)
		success = false
	}

	if success {
		fmt.Println("Forwarding", host, " -> ", ctx.RemoteIP().String()+":"+fmt.Sprintf("%d", port))
	}
}

func main() {
	fmt.Println("Proxy server started on :80")
	proxy.SetProduction()
	if err := fasthttp.ListenAndServe(":80", ProxyHandler); err != nil {
		log.Fatal(err)
	}
}
