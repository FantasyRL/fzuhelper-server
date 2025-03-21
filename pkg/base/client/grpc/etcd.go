/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/west2-online/fzuhelper-server/pkg/constants"
	"github.com/west2-online/fzuhelper-server/pkg/logger"
	"github.com/west2-online/fzuhelper-server/pkg/utils"
)

const (
	etcdPrefix = "grpc"
)

// EtcdResolver 用来封装 etcd 客户端和解析出来的 gRPC endpoint 列表
type EtcdResolver struct {
	EtcdClient  *clientv3.Client
	serviceName string              // etcd 里存服务地址列表的 key，如 "ai_agent"
	Endpoints   utils.AtomicStrings // 当前解析到的服务列表，原子性的string确保并发安全
	prefix      string
}

// NewEtcdResolver 创建一个etcdResolver
func NewEtcdResolver(endpoints []string) (*EtcdResolver, error) {
	cfg := &clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: constants.EtcdDialTimeout,
	}

	etcdClient, err := clientv3.New(*cfg)
	if err != nil {
		return nil, fmt.Errorf("NewEtcdResolver: connect etcd failed: %w", err)
	}
	return &EtcdResolver{
		EtcdClient: etcdClient,
		prefix:     etcdPrefix,
		Endpoints:  *utils.NewAtomicStrings(),
	}, nil
}

func (r *EtcdResolver) WatchAndResolve(ctx context.Context, changeCh chan<- struct{}) {
	watchCh := r.EtcdClient.Watch(ctx, r.getEtcdKeyPrefix())
	for {
		select {
		case wResp, ok := <-watchCh:
			if !ok {
				return
			}
			if wResp.Err() != nil {
				logger.Errorf("Etcd watch error: %v", wResp.Err())
				continue
			}
			for _, ev := range wResp.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					var eps []string
					if err := sonic.Unmarshal(ev.Kv.Value, &eps); err != nil {
						// 如果解析失败，认为是单个 endpoint
						eps = []string{string(ev.Kv.Value)}
					}
					r.Endpoints.Store(eps)
				case clientv3.EventTypeDelete:
					r.Endpoints.Store(nil)
				}
			}
			changeCh <- struct{}{}

		case <-ctx.Done():
			return
		}
	}
}

/*
// initResolve 尝试从 etcd 中加载 endpoints，并统一解析为 []string
func (r *EtcdResolver) initResolve() error {
	resp, err := r.EtcdClient.Get(context.Background(), r.getEtcdKeyPrefix(), clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("etcd get failed: %w", err)
	}
	// 因为prefix已经写死了是prefix+svcName,其实和Kvs[0]没差，但这样写比较合理
	// 这里即使不存在eps也返回nil
	var allEndpoints []string
	for _, kv := range resp.Kvs {
		val := kv.Value
		var eps []string
		if err = sonic.Unmarshal(val, &eps); err != nil {
			eps = []string{string(val)}
		}
		allEndpoints = append(allEndpoints, eps...)
	}
	r.Endpoints.Store(allEndpoints)
	return nil
}
*/

// GetRandomEndpoint 随机返回一个 endpoint
func (r *EtcdResolver) GetRandomEndpoint() (string, bool) {
	return r.Endpoints.Random()
}

func (r *EtcdResolver) getEtcdKeyPrefix() string {
	return strings.Join([]string{r.prefix, "/", r.serviceName}, "")
}
