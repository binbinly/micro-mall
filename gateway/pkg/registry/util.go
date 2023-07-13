package registry

import (
	"fmt"
	"net"
	"strings"
)

// HostPort format addr and port suitable for dial
func HostPort(addr string, port any) string {
	host := addr
	if strings.Count(addr, ":") > 0 {
		host = fmt.Sprintf("[%s]", addr)
	}
	// when port is blank or 0, host is a queue name
	if v, ok := port.(string); ok && v == "" {
		return host
	} else if v, ok := port.(int); ok && v == 0 && net.ParseIP(host) == nil {
		return host
	}

	return fmt.Sprintf("%s:%v", host, port)
}

// CopyService make a copy of service
func CopyService(service *Service) *Service {
	// copy service
	s := new(Service)
	*s = *service

	// copy nodes
	nodes := make([]*Node, len(service.Nodes))
	for j, node := range service.Nodes {
		n := new(Node)
		*n = *node
		nodes[j] = n
	}
	s.Nodes = nodes

	// copy endpoints
	eps := make([]*Endpoint, len(service.Endpoints))
	for j, ep := range service.Endpoints {
		e := new(Endpoint)
		*e = *ep
		eps[j] = e
	}
	s.Endpoints = eps
	return s
}
