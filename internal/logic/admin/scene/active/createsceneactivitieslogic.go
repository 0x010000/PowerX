package active

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSceneActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSceneActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSceneActivitiesLogic {
	return &CreateSceneActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// CreateSceneActivities
//  @Description: 创建活动
//  @receiver active
//  @param opt
//  @return resp
//  @return err
//
func (active *CreateSceneActivitiesLogic) CreateSceneActivities(opt *types.ActionActiveRequest) (resp *types.StateReply, err error) {

	err = active.svcCtx.PowerX.Scene.Svc.CreateActivities(active.ctx, opt)

	return &types.StateReply{
		Status: `success`,
	}, err

}
