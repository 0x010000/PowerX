package scene

import (
	"PowerX/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type SceneActivities struct {
	model.Model
	Aid        string    `gorm:"comment:活动ID;column:aid;unique;" json:"aid"`
	Name       string    `gorm:"comment:活动名称;column:name" json:"name"`
	Desc       string    `gorm:"comment:描述;column:desc" json:"desc"`
	Owner      string    `gorm:"comment:负责人(userId逗号隔开);column:owner" json:"owner"`
	StartTime  time.Time `gorm:"comment:开始时间;column:start_time" json:"start_time"`
	EndTime    time.Time `gorm:"comment:结束时间;column:end_time" json:"end_time"`
	ClassifyId int       `gorm:"comment:活动分类;column:classify_id" json:"classify_id"`
	SceneId    int       `gorm:"comment:活动场景ID;column:scene_id" json:"scene_id"`
	CoverLink  string    `gorm:"comment:活动封面Link;column:cover_link" json:"cover_link"`
	Link       string    `gorm:"comment:落地页(场景落地页;其他落地页);column:link" json:"link"`
	Cpm        int       `gorm:"comment:展示次数;column:cpm" json:"cpm"`
	State      int       `gorm:"comment:活动状态; 1:进行中 2:已结束 6:上架 7:结束;column:state" json:"state"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneActivities) TableName() string {
	return `scene_activities`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *SceneActivities) Query(db *gorm.DB) (active []*SceneActivities) {

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
func (e *SceneActivities) Action(db *gorm.DB, active []*SceneActivities) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "aid"}}, UpdateAll: true}).Create(&active).Error
	if err != nil {
		panic(err)
	}

}
