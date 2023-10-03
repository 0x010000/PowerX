package cases

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHealthHeightCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHealthHeightCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHealthHeightCaseLogic {
	return &CreateHealthHeightCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// CreateHealthHeightCase
//  @Description:
//  @receiver cases
//  @param opt
//  @return resp
//  @return err
//
func (cases *CreateHealthHeightCaseLogic) CreateHealthHeightCase(opt *types.HealthHeightCase) (resp *types.StateHealthReply, err error) {

	err = cases.svcCtx.Custom.Health.CreateCaseWithMenu(cases.ctx, opt)

	return &types.StateHealthReply{
		Status: `success`,
	}, err
}
