package consul

import (
	"fmt"
	"sync"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"

	"gateway/pkg/registry"
)

var _ resolver.Resolver = &consulResolver{}

// Init 返回一个resolver.Builder的实例
func Init(rs registry.Registry) {
	resolver.Register(NewBuilder(rs))
}

type consulResolver struct {
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string
	disableServiceConfig bool
	lastIndex            uint64
}

// watcher 更新服务变化
func (cr *consulResolver) watcher(rs registry.Registry) {
	defer cr.wg.Done()

	consul := rs.(*consulRegistry)
	for {
		services, metaInfo, err := consul.client.Health().Service(cr.name, "", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			logger.Warnf("[registry.consul] watcher services name:%v, err:%v", cr.name, err)
			time.Sleep(time.Second)
			continue
		}
		cr.lastIndex = metaInfo.LastIndex
		if len(services) == 0 {
			time.Sleep(time.Second)
			continue
		}
		var adds []resolver.Address

		for _, v := range services {
			var del bool
			for _, check := range v.Checks {
				// delete the node if the status is critical
				if check.Status == "critical" {
					del = true
					break
				}
			}
			// if delete then skip the node
			if del {
				continue
			}
			adds = append(adds, resolver.Address{Addr: fmt.Sprintf("%v:%v", v.Service.Address, v.Service.Port)})
		}

		state := &resolver.State{
			Addresses: adds,
		}
		cr.cc.UpdateState(*state)
	}
}

// ResolveNow 方法什么也不做，因为和consul保持了发布订阅的关系
// 不需要像dns_resolver那个定时的去刷新
func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {}

func (cr *consulResolver) Close() {}
