package logic

import (
	"context"

	"Coattails/core/internal/svc"
	"Coattails/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoattailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoattailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoattailsLogic {
	return &CoattailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoattailsLogic) Coattails(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
