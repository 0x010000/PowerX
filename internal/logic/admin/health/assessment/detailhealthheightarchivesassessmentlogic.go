package assessment

import (
	"PowerX/internal/model/health/height"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailHealthHeightArchivesAssessmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailHealthHeightArchivesAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailHealthHeightArchivesAssessmentLogic {
	return &DetailHealthHeightArchivesAssessmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// DetailHealthHeightArchivesAssessment
//  @Description:
//  @receiver assess
//  @param opt
//  @return resp
//  @return err
//
func (assess *DetailHealthHeightArchivesAssessmentLogic) DetailHealthHeightArchivesAssessment(opt *types.PathRequest) (resp *types.AssessmentReply, err error) {

	info, err := assess.svcCtx.Custom.Health.DetailHeightAssessmentInfo(assess.ctx, opt.Pid)

	return &types.AssessmentReply{
		HeathHeightArchives:   nil,
		HeathHeightAssessment: assess.assessment(info),
	}, err

}

//
// assessment
//  @Description:
//  @receiver assess
//  @param assessment
//  @return *types.Assessment
//
func (assess *DetailHealthHeightArchivesAssessmentLogic) assessment(assessment *height.HeathHeightAssessment) *types.Assessment {

	return &types.Assessment{
		Pid:                      assessment.Pid,
		Age:                      assessment.Age,
		NowHeight:                assessment.NowHeight,
		NowHeightCycleRatio:      assessment.NowHeightCycleRatio,
		NowWeight:                assessment.NowWeight,
		NowWeightCycleRatio:      assessment.NowWeightCycleRatio,
		BMI:                      assessment.BMI,
		BMICycleRatio:            assessment.BMICycleRatio,
		AbdominalGirth:           assessment.AbdominalGirth,
		ExpectHeight:             assessment.ExpectHeight,
		ExpectHeightCycleRatio:   assessment.ExpectHeightCycleRatio,
		EstimateHeight:           assessment.EstimateHeight,
		EstimateHeightCycleRatio: assessment.EstimateHeightCycleRatio,
		BodyState:                assessment.BodyState,
		GrowthState:              assessment.GrowthState,
		IsFirst:                  assessment.IsFirst,
		Analysis:                 assessment.Analysis,
		RevaluateName:            assessment.RevaluateName,
		RevaluateTime:            assessment.RevaluateTime,
	}
}
