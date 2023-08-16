package scene

import (
	"PowerX/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SceneActivitiesParticipants struct {
	model.Model

	Aid             uint64 `gorm:"comment:活动ID;column:aid;" json:"aid"`
	UserId          string `gorm:"comment:用户ID;column:user_id;" json:"user_id"`
	UserName        string `gorm:"comment:用户;column:user_name;" json:"user_name"`
	ShareUserId     string `gorm:"comment:分享用户ID;column:share_user_id;" json:"share_user_id"`
	ShareUserChain  string `gorm:"comment:分享用户链;column:share_user_chain;" json:"share_user_chain"`
	CustomerGroupId string `gorm:"comment:客户群ID;column:customer_group_id;" json:"customer_group_id"`
	TaskState       int    `gorm:"comment:任务状态;column:task_state;" json:"task_state"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneActivitiesParticipants) TableName() string {
	return `scene_activities_participants`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *SceneActivitiesParticipants) Query(db *gorm.DB) (active []*SceneActivitiesParticipants) {

	err := db.Model(e).Find(&active).Error
	if err != nil {
		panic(err)
	}
	return active

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param active
//
func (e *SceneActivitiesParticipants) Action(db *gorm.DB, active []*SceneActivitiesParticipants) {

	err := db.Table(e.TableName()).Debug().Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&active).Error
	if err != nil {
		panic(err)
	}

}
