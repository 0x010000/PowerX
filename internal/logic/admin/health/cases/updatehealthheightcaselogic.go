package cases

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHealthHeightCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHealthHeightCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHealthHeightCaseLogic {
	return &UpdateHealthHeightCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// UpdateHealthHeightCase
//  @Description:
//  @receiver cases
//  @param opt
//  @return resp
//  @return err
//
func (cases *UpdateHealthHeightCaseLogic) UpdateHealthHeightCase(opt *types.HealthHeightCase) (resp *types.StateHealthReply, err error) {

	err = cases.svcCtx.Custom.Health.CreateCaseWithMenu(cases.ctx, opt)

	return &types.StateHealthReply{
		Status: `success`,
	}, err

}
