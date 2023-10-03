package assessment

import (
	"PowerX/internal/model/health/height"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHealthHeightArchivesAssessmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHealthHeightArchivesAssessmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHealthHeightArchivesAssessmentLogic {
	return &ListHealthHeightArchivesAssessmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (assess *ListHealthHeightArchivesAssessmentLogic) ListHealthHeightArchivesAssessment(opt *types.ListHealthHeightArchivesAssessmentRequest) (resp *types.AssessmentListReply, err error) {

	page, err := assess.svcCtx.Custom.Health.FindManyHeightAssessmentPage(assess.ctx, assess.OPT(opt))

	return &types.AssessmentListReply{
		List:      assess.DTO(page.List),
		PageIndex: page.PageIndex,
		PageSize:  page.PageSize,
		Total:     page.Total,
	}, err

}

//
// OPT
//  @Description:
//  @receiver assess
//  @param opt
//  @return resp
//
func (assess *ListHealthHeightArchivesAssessmentLogic) OPT(opt *types.ListHealthHeightArchivesAssessmentRequest) (resp *types.PageOption[types.ListHealthHeightArchivesAssessmentRequest]) {

	option := types.PageOption[types.ListHealthHeightArchivesAssessmentRequest]{
		Option:    *opt,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}

	return &option
}

//
// DTO
//  @Description:
//  @receiver assess
//  @param assessment
//  @return data
//
func (assess *ListHealthHeightArchivesAssessmentLogic) DTO(assessment []*height.HeathHeightAssessment) (data []*types.AssessmentReply) {

	if assessment != nil {
		for _, val := range assessment {
			data = append(data, assess.dto(val))
		}
	}
	return data
}

//
// dto
//  @Description:
//  @receiver assess
//  @param assessment
//  @return data
//
func (assess *ListHealthHeightArchivesAssessmentLogic) dto(assessment *height.HeathHeightAssessment) (data *types.AssessmentReply) {

	return &types.AssessmentReply{
		HeathHeightArchives:   assess.archives(assessment.HeathHeightArchives),
		HeathHeightAssessment: assess.assessment(assessment),
	}
}

//
// assessment
//  @Description:
//  @receiver assess
//  @param assessment
//  @return *types.Assessment
//
func (assess *ListHealthHeightArchivesAssessmentLogic) assessment(assessment *height.HeathHeightAssessment) *types.Assessment {

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

//
// archives
//  @Description:
//  @receiver assess
//  @param archive
//  @return *types.Archives
//
func (assess *ListHealthHeightArchivesAssessmentLogic) archives(archive *height.HeathHeightArchives) (data *types.Archives) {
	if archive != nil {
		data = &types.Archives{
			Oid:                         archive.Oid,
			Name:                        archive.Name,
			Desc:                        archive.Desc,
			ExternalUserId:              archive.ExternalUserId,
			FatherHeight:                archive.FatherHeight,
			FatherWeight:                archive.FatherWeight,
			MotherHeight:                archive.MotherHeight,
			MotherWeight:                archive.MotherWeight,
			Gender:                      archive.Gender,
			Age:                         archive.Age,
			Birth:                       archive.Birth,
			Weight:                      archive.Weight,
			Height:                      archive.Height,
			GestationalWeeks:            archive.GestationalWeeks,
			PrevAssessmentTime:          archive.PrevAssessmentTime,
			NextAssessmentTime:          archive.NextAssessmentTime,
			Portrait:                    archive.Portrait,
			OrgId:                       archive.OrgId,
			UserId:                      archive.UserId,
			State:                       archive.State,
			LastPid:                     archive.LastPid,
			HeathHeightArchivesGuardian: assess.archivesWithGuardian(archive.HeathHeightArchivesGuardian),
		}
	}
	return data
}

//
// archivesWithGuardian
//  @Description:
//  @receiver assess
//  @param guardian
//  @return guard
//
func (assess *ListHealthHeightArchivesAssessmentLogic) archivesWithGuardian(guardian []*height.HeathHeightArchivesGuardian) (guard []*types.ArchivesGuardian) {

	if guardian != nil {
		for _, val := range guardian {
			guard = append(guard, &types.ArchivesGuardian{
				Name:     val.Name,
				Relation: val.Relation,
				Mobile:   val.Mobile,
			})
		}

	}
	return guard
}
