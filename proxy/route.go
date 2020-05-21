package proxy

import (
	"errors"
	"net/http"
	"path"
	"strings"

	"github.com/dxvgef/tsing-gateway/global"
)

// 设置路由组及路由，如果存在则更新，不存在则新建
func SetRoute(routeGroupID, routePath, routeMethod, upstreamID string) error {
	if routeGroupID == "" {
		return errors.New("没有传入路由组ID,并且无法自动创建ID")
	}
	// if _, exist := p.Upstreams[upstreamID]; !exist {
	// 	return errors.New("上游ID:" + upstreamID + "不存在")
	// }
	if routePath == "" {
		routePath = "/"
	}
	if routeMethod == "" {
		routeMethod = "*"
	} else {
		routeMethod = strings.ToUpper(routeMethod)
	}
	if !global.InStr(global.Methods, routeMethod) {
		return errors.New("HTTP方法无效")
	}
	if _, exist := global.Routes[routeGroupID]; !exist {
		global.Routes[routeGroupID] = make(map[string]map[string]string)
	}
	if _, exist := global.Routes[routeGroupID][routePath]; !exist {
		global.Routes[routeGroupID][routePath] = make(map[string]string)
	}
	global.Routes[routeGroupID][routePath][routeMethod] = upstreamID
	return nil
}

// 删除路由
func DelRoute(routeGroupID, routePath, routeMethod string) error {
	if routeGroupID == "" {
		routeGroupID = global.SnowflakeNode.Generate().String()
	}
	if routeGroupID == "" {
		return errors.New("routeGroupID不能为空")
	}
	if routePath == "" {
		return errors.New("reqPath不能为空")
	}
	if routeMethod == "" {
		return errors.New("reqMethod不能为空")
	}
	routeMethod = strings.ToUpper(routeMethod)
	delete(global.Routes[routeGroupID][routePath], routeMethod)
	return nil
}

// 匹配路由，返回集群ID和匹配结果的HTTP状态码
func matchRoute(req *http.Request) (upstream global.UpstreamType, status int) {
	routeGroupID := ""
	reqPath := req.URL.Path
	reqMethod := req.Method
	matchResult := false

	// 匹配主机
	routeGroupID, matchResult = matchHost(req.Host)
	if !matchResult {
		status = http.StatusServiceUnavailable
		return
	}
	// 匹配路径
	reqPath, matchResult = matchPath(routeGroupID, reqPath)
	if !matchResult {
		status = http.StatusNotFound
		return
	}
	// 匹配方法
	reqMethod, matchResult = matchMethod(routeGroupID, reqPath, reqMethod)
	if !matchResult {
		status = http.StatusMethodNotAllowed
		return
	}
	// 匹配上游
	upstreamID := global.Routes[routeGroupID][reqPath][reqMethod]
	upstream, matchResult = matchUpstream(upstreamID)
	if !matchResult {
		status = http.StatusNotImplemented
		return
	}
	status = http.StatusOK
	return
}

// 匹配路径，返回最终匹配到的路径
func matchPath(routeGroupID, reoutePath string) (string, bool) {
	if reoutePath == "" {
		reoutePath = "/"
	}
	// 先尝试完全匹配路径
	if _, exist := global.Routes[routeGroupID][reoutePath]; exist {
		return reoutePath, true
	}
	// 尝试模糊匹配
	pathLastChar := reoutePath[len(reoutePath)-1]
	if pathLastChar != 47 {
		pos := strings.LastIndex(reoutePath, path.Base(reoutePath))
		reoutePath = reoutePath[:pos]
	}
	// todo 可能要将*号换成别的符号，因为和api(tsing)的路由规则冲突
	reoutePath = reoutePath + "*"
	if _, exist := global.Routes[routeGroupID][reoutePath]; exist {
		return reoutePath, true
	}
	return reoutePath, false
}

// 匹配方法，返回对应的集群ID
func matchMethod(routeGroupID, routePath, routeMethod string) (string, bool) {
	if _, exist := global.Routes[routeGroupID][routePath][routeMethod]; exist {
		return routeMethod, true
	}
	routeMethod = "*"
	if _, exist := global.Routes[routeGroupID][routePath][routeMethod]; exist {
		return routeMethod, true
	}
	return routeMethod, false
}
