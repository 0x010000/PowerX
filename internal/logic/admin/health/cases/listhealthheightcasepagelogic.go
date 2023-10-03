package cases

import (
	"PowerX/internal/model/health/height"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHealthHeightCasePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHealthHeightCasePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHealthHeightCasePageLogic {
	return &ListHealthHeightCasePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListHealthHeightCasePage
//  @Description:
//  @receiver cases
//  @param opt
//  @return resp
//  @return err
//
func (cases *ListHealthHeightCasePageLogic) ListHealthHeightCasePage(opt *types.HealthHeightCaseListRequest) (resp *types.HealthHeightCaseListReply, err error) {

	page, err := cases.svcCtx.Custom.Health.FindManyHeightCasePage(cases.ctx, cases.OPT(opt))

	return &types.HealthHeightCaseListReply{
		List:      cases.DTO(page.List),
		PageIndex: page.PageIndex,
		PageSize:  page.PageSize,
		Total:     page.Total,
	}, err

}

//
// OPT
//  @Description:
//  @receiver cases
//  @param opt
//  @return resp
//
func (cases *ListHealthHeightCasePageLogic) OPT(opt *types.HealthHeightCaseListRequest) (resp *types.PageOption[types.HealthHeightCaseListRequest]) {

	option := types.PageOption[types.HealthHeightCaseListRequest]{
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
func (cases *ListHealthHeightCasePageLogic) DTO(css []*height.HeathHeightCase) (data []*types.ListHealthHeightCaseInfo) {

	if css != nil {
		for _, val := range css {
			data = append(data, cases.dto(val))
		}
	}
	return data
}

//
// dto
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) dto(cs *height.HeathHeightCase) (data *types.ListHealthHeightCaseInfo) {

	return &types.ListHealthHeightCaseInfo{
		Oid:                        cs.Oid,
		Pid:                        cs.Pid,
		Cid:                        cs.Cid,
		HeathHeightArchivesCase:    cases.heathHeightArchivesCase(cs),
		HeathHeightCaseBone:        cases.heathHeightCaseBone(cs),
		HeathHeightCaseNourishment: cases.heathHeightCaseNourishment(cs),
		HeathHeightCaseMeal:        cases.heathHeightCaseMeal(cs),
		HeathHeightCaseSport:       cases.heathHeightCaseSport(cs),
		HeathHeightCaseSleep:       cases.heathHeightCaseSleep(cs),
		HeathHeightCaseEmotion:     cases.heathHeightCaseEmotion(cs),
		Archives:                   cases.archives(cs),
		Assessment:                 cases.assessment(cs),
	}
}

//
// heathHeightArchivesCase
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightArchivesCase(cs *height.HeathHeightCase) (data types.HeathHeightArchivesCase) {
	return types.HeathHeightArchivesCase{Diagnostic: cs.Diagnostic}
}

//
// heathHeightCaseBone
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseBone(cs *height.HeathHeightCase) (data types.HeathHeightCaseBone) {
	if cs.HeathHeightCaseBone != nil {
		for _, menu := range cs.HeathHeightCaseBone {
			data.HeathHeightCaseBoneWithMenu = append(data.HeathHeightCaseBoneWithMenu, &types.HeathHeightCaseBoneWithMenu{
				ExpectHeightYear:  menu.ExpectHeightYear,
				ExpectHeightMonth: menu.ExpectHeightMonth,
				BoneAge:           menu.BoneAge,
				ControlHeight:     menu.ControlHeight,
				ControlWeight:     menu.ControlWeight,
			})
		}
	}
	return data
}

//
// heathHeightCaseNourishment
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseNourishment(cs *height.HeathHeightCase) (data types.HeathHeightCaseNourishment) {
	if cs.HeathHeightCaseBone != nil {
		for _, menu := range cs.HeathHeightCaseNourishment {
			data.HeathHeightCaseNourishmentWithMenu = append(data.HeathHeightCaseNourishmentWithMenu, &types.HeathHeightCaseNourishmentWithMenu{
				Name:   menu.Name,
				Number: menu.Number,
				Units:  menu.Units,
				Remark: menu.Remark,
				Gid:    menu.Gid,
			})
		}
	}
	return data
}

//
// heathHeightCaseMeal
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseMeal(cs *height.HeathHeightCase) (data types.HeathHeightCaseMeal) {
	if cs.HeathHeightCaseBone != nil {
		data.HeathHeightCaseMealWithControl = types.HeathHeightCaseMealWithControl{
			ControlLess:       cs.ControlLess,
			ControlNonuse:     cs.ControlNonuse,
			ControlBreakfast:  cs.ControlBreakfast,
			ControlLunch:      cs.ControlLunch,
			ControlDinner:     cs.ControlDinner,
			ControlMealMinute: cs.ControlMealMinute,
		}
		for _, menu := range cs.HeathHeightCaseNourishment {
			data.HeathHeightCaseMealWithMenu = append(data.HeathHeightCaseMealWithMenu, &types.HeathHeightCaseMealWithMenu{
				Name:   menu.Name,
				Remark: menu.Remark,
				Gid:    menu.Gid,
			})
		}
	}
	return data
}

//
// heathHeightCaseSport
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseSport(cs *height.HeathHeightCase) (data types.HeathHeightCaseSport) {
	if cs.HeathHeightCaseSports != nil {
		for _, menu := range cs.HeathHeightCaseSports {
			data.HeathHeightCaseSportWithOther = types.HeathHeightCaseSportWithOther{
				Remark: cs.SportRemark,
			}
			data.HeathHeightCaseSportWithMenu = append(data.HeathHeightCaseSportWithMenu, &types.HeathHeightCaseSportWithMenu{
				Name:   menu.Name,
				Number: menu.Number,
				Units:  menu.Units,
				Remark: menu.Remark,
				Gid:    menu.Gid,
			})
		}
	}
	return data
}

//
// heathHeightCaseSleep
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseSleep(cs *height.HeathHeightCase) (data types.HeathHeightCaseSleep) {
	return types.HeathHeightCaseSleep{
		Remark: cs.SportRemark,
	}
}

//
// heathHeightCaseEmotion
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) heathHeightCaseEmotion(cs *height.HeathHeightCase) (data types.HeathHeightCaseEmotion) {
	return types.HeathHeightCaseEmotion{
		Remark: cs.SportRemark,
	}
}

//
// archives
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) archives(cs *height.HeathHeightCase) (data types.Archives) {

	if cs.HeathHeightArchives != nil {
		archive := cs.HeathHeightArchives
		data = types.Archives{
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
			HeathHeightArchivesGuardian: nil,
		}
	}
	return data

}

//
// assessment
//  @Description:
//  @receiver cases
//  @param cs
//  @return data
//
func (cases *ListHealthHeightCasePageLogic) assessment(cs *height.HeathHeightCase) (data types.Assessment) {

	if cs.HeathHeightAssessment != nil {
		assessment := cs.HeathHeightAssessment
		data = types.Assessment{
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
	return data

}
