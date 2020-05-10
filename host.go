package main

import (
	"errors"
	"strings"
)

// 新建主机
func (p *Proxy) NewHost(hostname, routeGroupID string, persistent bool) error {
	hostname = strings.ToLower(hostname)
	if _, ok := p.Hosts[hostname]; ok {
		return errors.New("主机名:" + hostname + "已存在")
	}
	if _, exist := p.RouteGroups[routeGroupID]; !exist {
		return errors.New("路由组ID:" + routeGroupID + "不存在")
	}
	p.Hosts[hostname] = routeGroupID
	return nil
}

// 写入主机，如果存在则覆盖，不存在则创建
func (p *Proxy) SetHost(hostname, routeGroupID string, persistent bool) error {
	hostname = strings.ToLower(hostname)
	if _, exist := p.RouteGroups[routeGroupID]; !exist {
		return errors.New("路由组ID:" + routeGroupID + "不存在")
	}
	p.Hosts[hostname] = routeGroupID
	return nil
}

// 匹配主机名，返回对应的路由组ID
func (p *Proxy) MatchHost(reqHost string) (string, bool) {
	pos := strings.LastIndex(reqHost, ":")
	if pos > -1 {
		reqHost = reqHost[:pos]
	}
	if _, exist := p.Hosts[reqHost]; exist {
		return p.Hosts[reqHost], true
	}
	reqHost = "*"
	if _, exist := p.Hosts[reqHost]; exist {
		return p.Hosts[reqHost], true
	}
	return p.Hosts[reqHost], false
}
