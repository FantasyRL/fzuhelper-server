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
	"fmt"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/west2-online/fzuhelper-server/pkg/constants"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"github.com/west2-online/fzuhelper-server/pkg/logger"
)

type ConnManager struct {
	Conn         *grpc.ClientConn
	EtcdResolver *EtcdResolver
	refreshing   int32 // 原子标记（0=未刷新，1=刷新中）,保证同时只有一个RefreshGRPC被调用
}

func initManager(serviceName string, r *EtcdResolver) (*ConnManager, error) {
	m := new(ConnManager)
	r.serviceName = serviceName // 确定订阅的service
	m.EtcdResolver = r
	err := m.getConn() // 初始化一个grpc连接(并不一定存在)
	if err != nil {
		if errno.ConvertErr(err).ErrorCode == errno.InternalTimeoutErrorCode {
			return nil, nil
		}
		return nil, fmt.Errorf("initManager initConn failed: %w", err)
	}

	return m, nil
}

// initConn 获取一个 grpc 连接
func (m *ConnManager) getConn() error {
	var ep string
	var ok bool
	// 尝试等待 endpoints 可用
	timeout := time.After(constants.RefreshGRPCTimeout)
	ticker := time.NewTicker(constants.GRPCGetEndPointTicker)
	defer ticker.Stop()
	for {
		ep, ok = m.EtcdResolver.GetRandomEndpoint()
		// 找到可用endpoint
		if ok {
			break
		}
		select {
		case <-timeout:
			return errno.EtcdNoEpTimeOutError
		case <-ticker.C:
			continue
		}
	}
	conn, err := grpc.NewClient(ep, grpc.WithIdleTimeout(constants.GRPCGetConnTimeout))
	if err != nil {
		return fmt.Errorf("dial failed: %w", err)
	}
	m.Conn = conn
	return nil
}

// RefreshGRPC 用于在etcd节点数更新时刷新client，并且保证了读写安全
func (m *ConnManager) RefreshGRPC() {
	if !atomic.CompareAndSwapInt32(&m.refreshing, 0, 1) {
		return // 已有协程在刷新，直接退出
	}
	defer atomic.StoreInt32(&m.refreshing, 0)
	err := m.getConn()
	if err != nil {
		logger.Error(err.Error())
	}
}
