package router

import (
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomMarshaler struct {
	m *runtime.JSONPb
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Marshal 将结构体转换为json
// notice int64, fixed64, uint64 => string https://protobuf.dev/programming-guides/proto3/#json
func (c *CustomMarshaler) Marshal(v any) ([]byte, error) {

	switch v.(type) {
	case *emptypb.Empty:
		return c.m.Marshal(&Response{
			Code:    0,
			Message: "ok",
			Data:    struct{}{},
		})
	}
	return c.m.Marshal(v)
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
