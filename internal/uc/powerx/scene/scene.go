package scene

import (
	"PowerX/internal/model/scene"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

var Scene IsceneInterface = new(sceneUseCase)

type sceneUseCase struct {
	db  *gorm.DB
	kv  *redis.Redis
	ctx context.Context
	//
	//  modelSceneQrcode
	//  @Description:
	//
	modelSceneQrcode
	//
	//  modelSceneActive
	//  @Description:
	//
	modelSceneActive
}
type (
	modelSceneQrcode struct {
		qrcode scene.SceneQrcode
	}
	modelSceneActive struct {
		activites    scene.SceneActivities
		qrcode       scene.SceneActivitiesQrcode
		participants scene.SceneActivitiesParticipants
	}
)

// NewOrganizationUseCase
//
//	@Description:
//	@param db
//	@param wework
//	@return iEmployeeInterface
func Repo(db *gorm.DB, kv *redis.Redis) IsceneInterface {

	return &sceneUseCase{
		db:  db,
		kv:  kv,
		ctx: context.TODO(),
	}

}

type TimerActiveInt int

const (
	SceneActiveStateUnplayedInt TimerActiveInt = iota + 1
	SceneActiveStateBeginInt
	SceneActiveStateFinishInt
)
