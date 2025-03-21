package grpc

import (
	"github.com/west2-online/fzuhelper-server/kitex_gen/ai_agent"
)

type AiAgentClient struct {
	Cli         ai_agent.AIAgentClient
	ConnManager *ConnManager
}
