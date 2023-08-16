package scene

import (
	"PowerX/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SceneContent struct {
	model.Model

	Name  string `gorm:"comment:场景名称;column:name" json:"name"`
	Desc  string `gorm:"comment:场景声明描述;column:desc" json:"desc"`
	Rule  string `gorm:"comment:场景规则JSON;column:rule" json:"rule"`
	State int    `json:"state"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneContent) TableName() string {
	return `scene_contents`
}

//
// Query
//  @Description:
//  @receiver e
//  @param db
//  @return content
//
func (e *SceneContent) Query(db *gorm.DB) (content []*SceneContent) {

	err := db.Model(e).Find(&content).Error
	if err != nil {
		panic(err)
	}
	return content

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param content
//
func (e *SceneContent) Action(db *gorm.DB, content []*SceneContent) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&content).Error
	if err != nil {
		panic(err)
	}

}
