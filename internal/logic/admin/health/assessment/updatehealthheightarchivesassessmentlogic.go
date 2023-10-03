package assessment

import (
	"PowerX/internal/types/errorx"
	"context"
	"github.com/pkg/errors"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHealthHeightArchivesAssessmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHealthHeightArchivesAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHealthHeightArchivesAssessmentLogic {
	return &UpdateHealthHeightArchivesAssessmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (assess *UpdateHealthHeightArchivesAssessmentLogic) UpdateHealthHeightArchivesAssessment(opt *types.ActionHealthHeightArchivesAssessmentRequest) (resp *types.StateHealthReply, err error) {

	err = assess.OPT(opt)
	if err != nil {
		return nil, errorx.ErrBadRequest
	}
	err = assess.svcCtx.Custom.Health.UpdateAssessmentWithArchives(assess.ctx, opt)

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
func (assess *UpdateHealthHeightArchivesAssessmentLogic) OPT(opt *types.ActionHealthHeightArchivesAssessmentRequest) error {

	if opt.HeathHeightArchives.Oid == `` {
		return errors.Wrap(errorx.ErrBadRequest, `oid error`)
	} else if opt.HeathHeightAssessment.Pid == `` {
		return errors.Wrap(errorx.ErrBadRequest, `pid error`)
	}

	return nil
}
