module cart

go 1.18

require (
	github.com/binbinly/pkg v0.0.0-00010101000000-000000000000
	github.com/envoyproxy/protoc-gen-validate v1.0.1
	github.com/pkg/errors v0.9.1
	go-micro.dev/v4 v4.10.2
	google.golang.org/protobuf v1.30.0
	gorm.io/gorm v1.25.1
	pkg v0.0.0-00010101000000-000000000000
)

replace cart => ./

replace pkg => ./../../pkg

replace github.com/binbinly/pkg => ./../../../golang-workspace/pkg
