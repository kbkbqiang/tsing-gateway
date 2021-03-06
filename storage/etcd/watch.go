package etcd

import (
	"context"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/rs/zerolog/log"

	"local/global"
)

// 监听变更
func (self *Etcd) Watch() error {
	ch := self.client.Watch(context.Background(), self.KeyPrefix+"/", clientv3.WithPrefix())
	for resp := range ch {
		for k := range resp.Events {
			switch resp.Events[k].Type {
			// 更新事件
			case clientv3.EventTypePut:
				if err := self.watchLoadData(resp.Events[k].Kv.Key, resp.Events[k].Kv.Value); err != nil {
					log.Err(err).Caller().Send()
				}
			// 删除事件
			case clientv3.EventTypeDelete:
				if err := self.watchDeleteData(resp.Events[k].Kv.Key); err != nil {
					log.Err(err).Caller().Send()
				}
			}
		}
	}
	return nil
}

// 监听存储器数据更新，同步本地数据
func (self *Etcd) watchLoadData(key, value []byte) error {
	keyStr := global.BytesToStr(key)
	// 加载主机
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/hosts/") {
		return self.LoadHost(keyStr, value)
	}
	// 加载服务
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/services/") {
		return self.LoadService(value)
	}
	// 加载路由
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/routes/") {
		return self.LoadRoute(keyStr, value)
	}
	return nil
}

// 监听存储器数据删除，同步本地数据
func (self *Etcd) watchDeleteData(key []byte) error {
	keyStr := global.BytesToStr(key)
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/hosts/") {
		return self.DeleteLocalHost(keyStr)
	}
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/services/") {
		return self.DeleteLocalService(keyStr)
	}
	if strings.HasPrefix(keyStr, self.KeyPrefix+"/routes/") {
		return self.DeleteLocalRoute(keyStr)
	}
	return nil
}
