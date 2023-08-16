package powerx

import (
	"PowerX/internal/uc/powerx/scene"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type SceneUseCase struct {
	db   *gorm.DB
	kv   *redis.Redis
	Cron *cron.Cron
	Svc  scene.IsceneInterface
}

func NewSceneUseCase(db *gorm.DB, kv *redis.Redis, c *cron.Cron) *SceneUseCase {
	return &SceneUseCase{
		db:   db,
		Svc:  scene.Repo(db, kv),
		Cron: c,
	}
}

// Schedule
//
//	@Description:
//	@receiver this
func (this *SceneUseCase) Schedule() {

	_, _ = this.Cron.AddFunc(`0 0 0 * *`, func() {
		var err error
		//unix := time.Now()

		err = this.Svc.InvokeTimerDetectionUpdateActiveState()
		if err != nil {
			logx.Info(fmt.Sprintf(`--- [%s] cron.schedule.call.active.state.error, %v.`, err))
		}

	})

	go this.Cron.Start()

}
