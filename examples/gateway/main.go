package main

import (
	"context"
	"examples/gateway/proto/hello"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(func(ctx context.Context, writer http.ResponseWriter, resp proto.Message) error {
		fmt.Printf("req:%#v\n", resp)
		// 填充默认值

		return nil
	}),
		//runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		//	MarshalOptions: protojson.MarshalOptions{
		//		UseProtoNames:   true, // 使用 proto 中字段，而不是驼峰
		//		EmitUnpopulated: true, //填充空默认字段
		//	},
		//}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &CustomMarshaler{
				m: &runtime.JSONPb{
					MarshalOptions:   protojson.MarshalOptions{},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}},
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("err:%v\n", err)

			//err = &runtime.HTTPStatusError{
			//	HTTPStatus: http.StatusBadGateway,
			//}
			//runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

type CustomMarshaler struct {
	m *runtime.JSONPb
}

type Response struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (c *CustomMarshaler) Marshal(v any) ([]byte, error) {
	return c.m.Marshal(&Response{
		Code:    200,
		Message: "successful",
		Data:    v,
	})
}
func (c *CustomMarshaler) Unmarshal(data []byte, v any) error {
	return c.m.Unmarshal(data, v)
}

func (c *CustomMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return c.m.NewDecoder(r)
}
func (c *CustomMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	return c.m.NewEncoder(w)
}
func (c *CustomMarshaler) ContentType(v any) string {
	return "application/json"
}
