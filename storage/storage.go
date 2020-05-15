package storage

import (
	"errors"

	"github.com/dxvgef/tsing-gateway/proxy"
	"github.com/dxvgef/tsing-gateway/storage/etcd"
)

// 存储器接口
type Storage interface {
	LoadAll() error          // 加载所有数据
	LoadAllUpstreams() error // 加载所有upstream数据
	LoadAllRoutes() error    // 加载所有route数据
	LoadAllHosts() error     // 加载所有host数据
	SaveAll() error          // 存储所有数据
	SaveAllUpstreams() error // 存储所有upstream数据
	SaveAllRoutes() error    // 存储所有route数据
	SaveAllHosts() error     // 存储所有host数据
}

// 构建存储器实例
// key为存储器的名称，value为存储器的参数json字符串
func Build(e *proxy.Engine, name, config string) (Storage, error) {
	switch name {
	case "etcd":
		sa, err := etcd.New(e, config)
		if err != nil {
			return nil, err
		}
		return sa, nil
	}
	return nil, errors.New("根据名称没有找到对应的存储器")
}