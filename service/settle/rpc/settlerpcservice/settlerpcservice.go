// Code generated by goctl. DO NOT EDIT!
// Source: settle.ptoto

package settlerpcservice

import (
	"context"

	"ayzy_platform/service/settle/rpc/types/settle"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SettleCommonRequest           = settle.SettleCommonRequest
	SettleSumResponse             = settle.SettleSumResponse
	SettleTypeSummaryListResponse = settle.SettleTypeSummaryListResponse
	SettleTypeSummaryRequest      = settle.SettleTypeSummaryRequest
	SettleTypeSummaryResponse     = settle.SettleTypeSummaryResponse

	SettleRpcService interface {
		GetSummaryBySettleType(ctx context.Context, in *SettleTypeSummaryRequest, opts ...grpc.CallOption) (*SettleTypeSummaryListResponse, error)
		GetSum(ctx context.Context, in *SettleCommonRequest, opts ...grpc.CallOption) (*SettleSumResponse, error)
	}

	defaultSettleRpcService struct {
		cli zrpc.Client
	}
)

func NewSettleRpcService(cli zrpc.Client) SettleRpcService {
	return &defaultSettleRpcService{
		cli: cli,
	}
}

func (m *defaultSettleRpcService) GetSummaryBySettleType(ctx context.Context, in *SettleTypeSummaryRequest, opts ...grpc.CallOption) (*SettleTypeSummaryListResponse, error) {
	client := settle.NewSettleRpcServiceClient(m.cli.Conn())
	return client.GetSummaryBySettleType(ctx, in, opts...)
}

func (m *defaultSettleRpcService) GetSum(ctx context.Context, in *SettleCommonRequest, opts ...grpc.CallOption) (*SettleSumResponse, error) {
	client := settle.NewSettleRpcServiceClient(m.cli.Conn())
	return client.GetSum(ctx, in, opts...)
}
