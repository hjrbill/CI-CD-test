package logic

import (
	"context"
	"go-zero-hello/internal/svc"
	"go-zero-hello/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "Hello " + req.Name
	return
}
