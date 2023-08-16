package scene

import (
	"PowerX/internal/model/scene"
	"PowerX/internal/types"
	"context"
)

type IsceneInterface interface {

	//
	//  @Description: qrcode
	//
	iQrcodeInterface
	//
	//  @Description: active
	//
	iActiveInterface
	//
	//  @Description: timer
	//
	iTimerInterface
	//
	//  @Description: participants
	//
	iActiveParticipantsInterface
}

//
//  iTimerInterface
//  @Description:
//
type iTimerInterface interface {

	//
	// InvokeTimerDetectionUpdateActiveState
	//  @Description: 变更活动状态
	//  @return err
	//
	InvokeTimerDetectionUpdateActiveState() (err error)
}

// iQrcodeInterface
// @Description:
type iQrcodeInterface interface {
	//
	// FindOneSceneQrcodeDetail
	//  @Description: 场景码详情
	//  @param qid
	//  @return *qrcode.QrcodeActive
	//
	FindOneSceneQrcodeDetail(qid string) *scene.SceneQrcode
	//
	// IncreaseSceneCpaNumber
	//  @Description: CPA+1
	//  @param qid
	//
	IncreaseSceneCpaNumber(qid string)
	//
	// FindSceneQrcodeOption
	//  @Description:
	//  @param opt
	//  @return []*scene.SceneQrcode
	//
	FindSceneQrcodeOption(opt *types.SceneQrcodeRequest) []*scene.SceneQrcode
}

//
//  iActiveInterface
//  @Description:
//
type iActiveInterface interface {
	//
	// FindManyActivitiesPage
	//  @Description: 活动列表/分页
	//  @param ctx
	//  @param opt
	//  @return *types.Page[*scene.SceneActivities]
	//  @return error
	//
	FindManyActivitiesPage(ctx context.Context, opt *types.PageOption[types.ActivitiesListRequest]) (*types.Page[*scene.SceneActivities], error)

	//
	// CreateActivities
	//  @Description: 创建活动
	//  @param ctx
	//  @param active
	//  @return error
	//
	CreateActivities(ctx context.Context, active *types.ActionActiveRequest) error

	//
	// UpdateActivities
	//  @Description: 更新活动
	//  @param ctx
	//  @param active
	//  @return error
	//
	UpdateActivities(ctx context.Context, active *types.ActionActiveRequest) error

	//
	// FindOneActivityDetail
	//  @Description: 活动详情
	//  @param aid
	//  @return active
	//
	FindOneActivityDetail(aid uint64) (active *scene.SceneActivities)
}

type iActiveParticipantsInterface interface {
	//
	// CreateActiveParticipants
	//  @Description:
	//  @param ctx
	//  @param models
	//  @return error
	//
	CreateActiveParticipants(ctx context.Context, models []*scene.SceneActivitiesParticipants) error
}
