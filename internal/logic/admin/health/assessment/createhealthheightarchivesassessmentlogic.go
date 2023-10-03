package assessment

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHealthHeightArchivesAssessmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHealthHeightArchivesAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHealthHeightArchivesAssessmentLogic {
	return &CreateHealthHeightArchivesAssessmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// CreateHealthHeightArchivesAssessment
//  @Description:
//  @receiver assess
//  @param opt
//  @return resp
//  @return err
//
func (assess *CreateHealthHeightArchivesAssessmentLogic) CreateHealthHeightArchivesAssessment(opt *types.ActionHealthHeightArchivesAssessmentRequest) (resp *types.StateHealthReply, err error) {
	err = assess.OPT(opt)
	if err != nil {
		return nil, errorx.ErrBadRequest
	}
	err = assess.svcCtx.Custom.Health.CreateAssessmentWithArchives(assess.ctx, opt)

	return &types.StateHealthReply{
		Status: `success`,
	}, err
}

//
// OPT
//  @Description:
//  @receiver assess
//  @param opt
//  @return error
//
func (assess *CreateHealthHeightArchivesAssessmentLogic) OPT(opt *types.ActionHealthHeightArchivesAssessmentRequest) error {

	return nil
}
