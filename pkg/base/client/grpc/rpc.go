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
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/west2-online/fzuhelper-server/config"
	"github.com/west2-online/fzuhelper-server/kitex_gen/ai_agent"
	"github.com/west2-online/fzuhelper-server/pkg/constants"
)

// T: grpc_gen Client; S: 二次封装的结构体
func initGRPCClient[T any, S any](
	serviceName string,
	newClientFunc func(cc grpc.ClientConnInterface) T,
	encapsulationFunc func(cli T, manager *ConnManager) *S,
) (*S, error) {

	if config.Etcd == nil || config.Etcd.Addr == "" {
		return nil, errors.New("config.Etcd.Addr is nil")
	}
	r, err := NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		return nil, fmt.Errorf("initGRPCClient etcd.NewEtcdResolver failed: %w", err)
	}
	m, err := initManager(serviceName, r)
	if err != nil {
		return nil, fmt.Errorf("initGRPCClient failed: %w", err)
	}
	client := encapsulationFunc(newClientFunc(m.Conn), m)
	go func() {
		go r.WatchAndResolve(context.Background()) // 确保活跃
		// your code...
	}()

	return client, nil
}

func InitAiAgentRPC() (*AiAgentClient, error) {
	return initGRPCClient(constants.AiAgentServiceName,
		ai_agent.NewAIAgentClient, func(cli ai_agent.AIAgentClient, manager *ConnManager) *AiAgentClient {
			return &AiAgentClient{
				Cli:         cli,
				ConnManager: manager,
			}
		})
}
