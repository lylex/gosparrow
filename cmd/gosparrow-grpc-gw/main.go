package main

import (
	"net"
)

func main() {
	sctx := newServeCtx()

	var err error
	addr := "localhost:8080"
	if sctx.l, err = net.Listen("tcp", addr); err != nil {
		// TODO handle error
	}
	sctx.addr = "localhost:8081"
	sctx.serve(nil, func(error) {})
}
