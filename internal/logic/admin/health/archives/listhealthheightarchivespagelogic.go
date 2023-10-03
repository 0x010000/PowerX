package archives

import (
	"PowerX/internal/model/health/height"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHealthHeightArchivesPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHealthHeightArchivesPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHealthHeightArchivesPageLogic {
	return &ListHealthHeightArchivesPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListHealthHeightArchivesPage
//  @Description:
//  @receiver archive
//  @param opt
//  @return resp
//  @return err
//
func (archive *ListHealthHeightArchivesPageLogic) ListHealthHeightArchivesPage(opt *types.HealthHeightArchivesListRequest) (resp *types.HealthHeightArchivesListReply, err error) {

	page, err := archive.svcCtx.Custom.Health.FindManyHeightArchivesPage(archive.ctx, archive.OPT(opt))

	return &types.HealthHeightArchivesListReply{
		List:      archive.DTO(page.List),
		PageIndex: page.PageIndex,
		PageSize:  page.PageSize,
		Total:     page.Total,
	}, err
}

//
// OPT
//  @Description:
//  @receiver l
//  @param opt
//  @return resp
//
func (archive *ListHealthHeightArchivesPageLogic) OPT(opt *types.HealthHeightArchivesListRequest) (resp *types.PageOption[types.HealthHeightArchivesListRequest]) {

	option := types.PageOption[types.HealthHeightArchivesListRequest]{
		Option:    *opt,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
	}

	return &option
}

//
// DTO
//  @Description:
//  @receiver archive
//  @param archives
//  @return data
//
func (archive *ListHealthHeightArchivesPageLogic) DTO(archives []*height.HeathHeightArchives) (data []*types.Archives) {

	if archives != nil {
		for _, val := range archives {
			data = append(data, archive.dto(val))
		}
	}
	return data
}

//
// dto
//  @Description:
//  @receiver archive
//  @param archives
//  @return data
//
func (archive *ListHealthHeightArchivesPageLogic) dto(archives *height.HeathHeightArchives) (data *types.Archives) {

	return &types.Archives{
		Oid:                         archives.Oid,
		Name:                        archives.Name,
		Desc:                        archives.Desc,
		ExternalUserId:              archives.ExternalUserId,
		FatherHeight:                archives.FatherHeight,
		FatherWeight:                archives.FatherWeight,
		MotherHeight:                archives.MotherHeight,
		MotherWeight:                archives.MotherWeight,
		Gender:                      archives.Gender,
		Age:                         archives.Age,
		Birth:                       archives.Birth,
		Weight:                      archives.Weight,
		Height:                      archives.Height,
		GestationalWeeks:            archives.GestationalWeeks,
		PrevAssessmentTime:          archives.PrevAssessmentTime,
		NextAssessmentTime:          archives.NextAssessmentTime,
		Portrait:                    archives.Portrait,
		OrgId:                       archives.OrgId,
		UserId:                      archives.UserId,
		State:                       archives.State,
		LastPid:                     archives.LastPid,
		HeathHeightArchivesGuardian: archive.heathHeightArchivesGuardian(archives.HeathHeightArchivesGuardian),
	}
}

//
// heathHeightArchivesGuardian
//  @Description:
//  @receiver archive
//  @param guardians
//  @return guardian
//
func (archive *ListHealthHeightArchivesPageLogic) heathHeightArchivesGuardian(guardians []*height.HeathHeightArchivesGuardian) (guardian []*types.ArchivesGuardian) {

	if guardians != nil {
		for _, val := range guardians {
			guardian = append(guardian, &types.ArchivesGuardian{
				Name:     val.Name,
				Relation: val.Relation,
				Mobile:   val.Mobile,
			})
		}
	}
	return guardian
}
