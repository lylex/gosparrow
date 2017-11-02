package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gw "gosparrow/pkg/gosparrow"
)

type serveCtx struct {
	l            net.Listener
	addr         string
	secure       bool
	insecure     bool
	ctx          context.Context
	cancel       context.CancelFunc
	userHandlers map[string]http.Handler
}

// TODO pass in config
func newServeCtx() *serveCtx {
	ctx, cancel := context.WithCancel(context.Background())
	return &serveCtx{ctx: ctx, cancel: cancel, userHandlers: make(map[string]http.Handler)}
}

func (sctx *serveCtx) serve(handler http.Handler, errHandler func(error)) error {
	logger := log.New(ioutil.Discard, "httplog", 0)

	m := cmux.New(sctx.l)

	if sctx.secure {
		// TODO handle secure brach
	} else {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		gwmux, err := sctx.registerGateway(opts)
		if err != nil {
			return err
		}

		httpmux := sctx.createMux(gwmux, handler)
		srvhttp := &http.Server{
			Handler:  httpmux,
			ErrorLog: logger, // do not log user error
		}
		httpl := m.Match(cmux.HTTP1())
		go func() { errHandler(srvhttp.Serve(httpl)) }()
	}

	return m.Serve()
}

type registerHandlerFunc func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

func (sctx *serveCtx) registerGateway(opts []grpc.DialOption) (*runtime.ServeMux, error) {
	ctx := sctx.ctx
	conn, err := grpc.DialContext(ctx, sctx.addr, opts...)
	if err != nil {
		return nil, err
	}
	gwmux := runtime.NewServeMux()

	handlers := []registerHandlerFunc{
		gw.RegisterGosparrowHandler,
	}
	for _, h := range handlers {
		if err := h(ctx, gwmux, conn); err != nil {
			return nil, err
		}
	}

	go func() {
		<-ctx.Done()
		if cerr := conn.Close(); cerr != nil {
			glog.Warningf("failed to close conn to %s: %v", sctx.l.Addr().String(), cerr)
		}
	}()

	return gwmux, nil
}

func (sctx *serveCtx) createMux(gwmux *runtime.ServeMux, handler http.Handler) *http.ServeMux {
	httpmux := http.NewServeMux()
	for path, h := range sctx.userHandlers {
		httpmux.Handle(path, h)
	}

	httpmux.Handle(
		"/",
		wsproxy.WebsocketProxy(
			gwmux,
			wsproxy.WithRequestMutator(
				func(incoming *http.Request, outgoing *http.Request) *http.Request {
					outgoing.Method = "POST"
					return outgoing
				},
			),
		),
	)
	if handler != nil {
		httpmux.Handle("/", handler)

	}
	return httpmux
}
