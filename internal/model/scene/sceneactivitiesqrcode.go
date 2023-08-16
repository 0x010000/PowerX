package scene

import (
	"PowerX/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SceneActivitiesQrcode struct {
	model.Model

	Aid  uint64 `gorm:"comment:活动ID;column:aid;" json:"aid"`
	Qid  string `gorm:"comment:场景码QID;column:qid;" json:"qid"`
	Link string `gorm:"comment:场景码link;column:link;" json:"link"`

	SceneQrcode *SceneQrcode `gorm:"foreignKey:QId;references:Qid" json:"SceneQrcode"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneActivitiesQrcode) TableName() string {
	return `scene_activities_qrcodes`
}

//
// Query
//  @Description:
//  @receiver e
//  @param db
//  @return qrcode
//
func (e *SceneActivitiesQrcode) Query(db *gorm.DB) (qrcode []*SceneActivitiesQrcode) {

	err := db.Model(e).Find(&qrcode).Error
	if err != nil {
		panic(err)
	}
	return qrcode

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param qrcode
//
func (e *SceneActivitiesQrcode) Action(db *gorm.DB, qrcode []*SceneActivitiesQrcode) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&qrcode).Error
	if err != nil {
		panic(err)
	}

}
