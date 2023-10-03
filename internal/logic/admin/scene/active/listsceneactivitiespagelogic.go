package active

import (
	"PowerX/internal/model/scene"
	"context"
	"strings"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSceneActivitiesPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSceneActivitiesPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSceneActivitiesPageLogic {
	return &ListSceneActivitiesPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListSceneActivitiesPage
//  @Description: 活动列表/分页
//  @receiver active
//  @param opt
//  @return resp
//  @return err
//
func (active *ListSceneActivitiesPageLogic) ListSceneActivitiesPage(opt *types.ActivitiesListRequest) (resp *types.ActivitiesListReply, err error) {

	data, err := active.svcCtx.PowerX.Scene.Svc.FindManyActivitiesPage(active.ctx, active.OPT(opt))

	return &types.ActivitiesListReply{
		List:      active.DTO(data.List),
		PageIndex: data.PageIndex,
		PageSize:  data.PageSize,
		Total:     data.Total,
	}, err

}

//
// OPT
//  @Description:
//  @receiver active
//  @param opt
//  @return *types.PageOption[*types.ActivitiesListRequest]
//
func (active *ListSceneActivitiesPageLogic) OPT(opt *types.ActivitiesListRequest) *types.PageOption[types.ActivitiesListRequest] {

	var option types.PageOption[types.ActivitiesListRequest]
	option.Option.Name = opt.Name
	option.Option.State = opt.State
	option.DefaultPageIfNotSet()
	return &option

}

//
// DTO
//  @Description:
//  @receiver active
//  @param data
//  @return Active
//
func (active *ListSceneActivitiesPageLogic) DTO(data []*scene.SceneActivities) (activites []*types.Active) {

	if data != nil {
		for _, val := range data {
			activites = append(activites, active.dto(val))
		}
	}
	return activites
}

//
// dto
//  @Description:
//  @receiver active
//  @param val
//  @return activites
//
func (active *ListSceneActivitiesPageLogic) dto(val *scene.SceneActivities) *types.Active {

	return &types.Active{
		Activities:       active.activities(val),
		ActivitiesPoster: active.activitiesPoster(val),
		ActivitiesSceneQrcode: types.ActivitiesWithQrcode{
			MemberMaxLimit: val.MemberMaxLimit, ActionSceneQrcode: active.activitiesSceneQrcode(val.ActiveGroupQrcode)},
		ActivitiesParticipants: active.activitiesParticipants(val.ActiveParticipants),
		ActivitiesContent:      nil,
	}
}

//
// activities
//  @Description:
//  @receiver active
//  @param val
//  @return types.Activities
//
func (active *ListSceneActivitiesPageLogic) activities(val *scene.SceneActivities) types.Activities {

	return types.Activities{
		Id:                  int(val.Id),
		Aid:                 val.Aid,
		Name:                val.Name,
		Desc:                val.Desc,
		Owner:               strings.Split(val.Owner, `,`),
		StartTime:           val.StartTime.Format(`2006-01-02 13:04:05`),
		EndTime:             val.StartTime.Format(`2006-01-02 13:04:05`),
		ClassifyId:          val.ClassifyId,
		ActivitiesContentId: val.ActivitiesContentId,
		State:               val.State,
	}
}

//
// activitiesPoster
//  @Description:
//  @receiver active
//  @param val
//  @return types.Poster
//
func (active *ListSceneActivitiesPageLogic) activitiesPoster(val *scene.SceneActivities) types.Poster {
	var position []string
	if val.Position != `` {
		position = strings.Split(val.Position, `,`)
	}
	return types.Poster{
		PhotoState: val.PhotoState,
		AliseState: val.AliseState,
		CoverLink:  val.CoverLink,
		Link:       val.Link,
		Position:   position,
	}

}

//
// activitiesParticipants
//  @Description:
//  @receiver active
//  @param participants
//  @return group
//
func (active *ListSceneActivitiesPageLogic) activitiesParticipants(participants []*scene.SceneActivitiesParticipants) (group types.Participants) {

	if participants != nil {
		for _, part := range participants {
			group.EnterGroupNumber++
			if part.TaskState > 0 {
				group.DoneTaskNumber++
			}
		}
	}
	return group
}

//
// activitiesSceneQrcode
//  @Description:
//  @receiver active
//  @param sceneQrcode
//  @return qrcode
//
func (active *ListSceneActivitiesPageLogic) activitiesSceneQrcode(sceneQrcode []*scene.SceneActivitiesQrcode) (qrcode []*types.ActivitiesQrcode) {

	if sceneQrcode != nil {
		for _, val := range sceneQrcode {
			qrcode = append(qrcode, &types.ActivitiesQrcode{
				Qid:  val.Qid,
				Link: val.Link,
			})
		}
	}
	return qrcode
}
