package main

import (
	"context"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4"
	"log"
	pb "pkg/proto/market"
)

func main() {
	service := micro.NewService()
	service.Init()

	c := pb.NewMarketService("mall.market", grpcc.NewClient())
	res, err := c.GetHomeData(context.Background(), &empty.Empty{})
	log.Printf("err:%v\n", err)
	log.Printf("res:%v\n", res)

}
