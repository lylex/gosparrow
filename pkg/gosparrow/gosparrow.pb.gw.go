// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: gosparrow.proto

/*
Package gosparrow is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package gosparrow

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

var (
	filter_Gosparrow_GetName_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Gosparrow_GetName_0(ctx context.Context, marshaler runtime.Marshaler, client GosparrowClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetNameReq
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_Gosparrow_GetName_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetName(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterGosparrowHandlerFromEndpoint is same as RegisterGosparrowHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGosparrowHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGosparrowHandler(ctx, mux, conn)
}

// RegisterGosparrowHandler registers the http handlers for service Gosparrow to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGosparrowHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGosparrowHandlerClient(ctx, mux, NewGosparrowClient(conn))
}

// RegisterGosparrowHandler registers the http handlers for service Gosparrow to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "GosparrowClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GosparrowClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GosparrowClient" to call the correct interceptors.
func RegisterGosparrowHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GosparrowClient) error {

	mux.Handle("GET", pattern_Gosparrow_GetName_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gosparrow_GetName_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gosparrow_GetName_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Gosparrow_GetName_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"name"}, ""))
)

var (
	forward_Gosparrow_GetName_0 = runtime.ForwardResponseMessage
)
