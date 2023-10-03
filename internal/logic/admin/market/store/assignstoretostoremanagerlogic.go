package store

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignStoreToStoreManagerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignStoreToStoreManagerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignStoreToStoreManagerLogic {
	return &AssignStoreToStoreManagerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignStoreToStoreManagerLogic) AssignStoreToStoreManager(req *types.AssignStoreManagerRequest) (resp *types.AssignStoreManagerReply, err error) {

	err = l.svcCtx.PowerX.Store.ActionPivotStoreToEmployee(l.ctx, req.Id, req.EmployeeId, req.UserId)
	return &types.AssignStoreManagerReply{
		Store: types.Store{},
	}, err
}
