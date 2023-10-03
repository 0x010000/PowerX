package active

import (
	"PowerX/internal/model/scene"
	"PowerX/internal/types/errorx"
	"context"
	"fmt"
	"strings"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionSceneActivitiesDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionSceneActivitiesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionSceneActivitiesDetailLogic {
	return &ActionSceneActivitiesDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ActionSceneActivitiesDetail
//  @Description:
//  @receiver active
//  @param opt
//  @return resp
//  @return err
//
func (active *ActionSceneActivitiesDetailLogic) ActionSceneActivitiesDetail(opt *types.ActivitiesInfoRequest) (resp *types.ActivitiesInfoReply, err error) {

	if opt.Aid == 0 || opt.UserId == `` {
		return nil, errorx.ErrBadRequest
	}

	//user := wechat.Scrm.GetCustomerGroupFromKVByUserId(`wr2lz4UgAAiLjAdxEjtljEVEZDMD-MzA`, `wm2lz4UgAAcoeq0sJVvA9vusEfMXdjeA`)
	detail := active.svcCtx.PowerX.Scene.Svc.FindOneActivityDetail(opt.Aid)
	resp = active.DTO(detail)
	// init participants
	participant := &scene.SceneActivitiesParticipants{
		Aid:            opt.Aid,
		UserId:         opt.UserId,
		UserName:       opt.UserName,
		ShareUserId:    opt.ShareUserId,
		ShareUserChain: opt.ShareUserId,
		TaskState:      0,
	}
	// qrcode
	if qrcode := resp.ActivitiesSceneQrcode; qrcode != nil {
		resp.State = active.rule(participant, qrcode, detail.MemberMaxLimit)
		if participant.UnifyId == `` {
			participant.UnifyId = qrcode[0].UnifyId
		}
	}
	// one task with only userId
	if detail.ActiveParticipants != nil {
		active.shareDone(detail.ActiveParticipants, participant)
	}
	// create participants record
	if participant.Id == 0 {
		err = active.svcCtx.PowerX.Scene.Svc.CreateActiveParticipants(active.ctx, []*scene.SceneActivitiesParticipants{participant})
	}

	return resp, err
}

//
// rule
//  @Description:
//  @receiver active
//  @param participant
//  @param qrcode
//  @param max
//  @return state
//
func (active *ActionSceneActivitiesDetailLogic) rule(participant *scene.SceneActivitiesParticipants, qrcode []*types.ActivitiesQrcodeWeb, max int) (state bool) {

	for i, code := range qrcode {
		chat := active.svcCtx.PowerX.SCRM.Wechat.GetCustomerGroupFromKVByChatId(code.UnifyId, true)
		if chat == nil {
			return state
		}
		// delete group overstep memberMaxLimit
		if len(chat.MemberList) > max {
			if len(qrcode) > 1 {
				qrcode = append(qrcode[:i], qrcode[i+1:]...)
			}
		}
		// select valid userId
		for _, member := range chat.MemberList {
			if member.UserID == participant.UserId {
				participant.UnifyId = code.UnifyId
				participant.UserName = member.Name
				state = true
			}
		}
	}
	return state
}

//
// DTO
//  @Description:
//  @receiver active
//  @param activities
//  @return *types.ActivitiesInfoReply
//
func (active *ActionSceneActivitiesDetailLogic) DTO(activities *scene.SceneActivities) *types.ActivitiesInfoReply {

	if activities.State != 2 || activities.Aid == 0 {
		return nil
	}
	return &types.ActivitiesInfoReply{
		Activities:             active.activities(activities),
		ActivitiesPoster:       active.activitiesPoster(activities),
		ActivitiesSceneQrcode:  active.activitiesSceneQrcode(activities.ActiveGroupQrcode),
		ActivitiesParticipants: active.activitiesParticipants(activities.ActiveParticipants),
		ActivitiesContent:      nil,
		State:                  false,
	}
}

//
// activities
//  @Description:
//  @receiver active
//  @param val
//  @return types.ActivitiesWeb
//
func (active *ActionSceneActivitiesDetailLogic) activities(val *scene.SceneActivities) types.ActivitiesWeb {

	return types.ActivitiesWeb{
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
//  @return types.PosterWeb
//
func (active *ActionSceneActivitiesDetailLogic) activitiesPoster(val *scene.SceneActivities) types.PosterWeb {
	var position []string
	if val.Position != `` {
		position = strings.Split(val.Position, `,`)
	}
	return types.PosterWeb{
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
func (active *ActionSceneActivitiesDetailLogic) activitiesParticipants(participants []*scene.SceneActivitiesParticipants) (parts []*types.ParticipantsWeb) {

	if participants != nil {
		for _, part := range participants {
			parts = append(parts, &types.ParticipantsWeb{
				UserId:               part.UserId,
				UserName:             part.UserName,
				TaskState:            part.TaskState,
				ShareTaskNumber:      0,
				VaildShareTaskNumber: 0,
			})
		}
	}
	return parts
}

//
// activitiesSceneQrcode
//  @Description:
//  @receiver active
//  @param sceneQrcode
//  @return qrcode
//
func (active *ActionSceneActivitiesDetailLogic) activitiesSceneQrcode(sceneQrcode []*scene.SceneActivitiesQrcode) (qrcode []*types.ActivitiesQrcodeWeb) {

	if sceneQrcode != nil {
		for _, val := range sceneQrcode {
			qrcode = append(qrcode, &types.ActivitiesQrcodeWeb{
				Qid:     val.Qid,
				Link:    val.Link,
				UnifyId: val.SceneQrcode.UnifyId,
			})
		}
	}
	return qrcode
}

//
// shareDone
//  @Description:
//  @receiver active
//  @param data
//  @param participants
//
func (active *ActionSceneActivitiesDetailLogic) shareDone(data []*scene.SceneActivitiesParticipants, participant *scene.SceneActivitiesParticipants) {

	for _, part := range data {
		if participant.UserId == part.UserId {
			// hit Id
			participant.Id = part.Id
			break
		} else if participant.ShareUserId == part.UserId {
			// use shareUserId chain
			if part.ShareUserChain != `` {
				participant.ShareUserChain = fmt.Sprintf(`%s,%s`, part.ShareUserChain, participant.UserId)
			}
		} else {

		}
	}
}
