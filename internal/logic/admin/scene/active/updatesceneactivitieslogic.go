package active

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSceneActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSceneActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSceneActivitiesLogic {
	return &UpdateSceneActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// UpdateSceneActivities
//  @Description: 更新活动
//  @receiver active
//  @param opt
//  @return resp
//  @return err
//
func (active *UpdateSceneActivitiesLogic) UpdateSceneActivities(opt *types.ActionActiveRequest) (resp *types.StateReply, err error) {

	if opt.Aid == 0 {
		return resp, errorx.ErrBadRequest
	}
	if opt.ActionSceneQrcode.MemberMaxLimit == 0 {
		opt.ActionSceneQrcode.MemberMaxLimit = 400
	}
	err = active.svcCtx.PowerX.Scene.Svc.UpdateActivities(active.ctx, opt)

	return &types.StateReply{
		Status: `success`,
	}, err
}
