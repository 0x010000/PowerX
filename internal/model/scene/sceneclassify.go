package scene

import (
	"PowerX/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SceneClassify struct {
	model.Model

	Name string `gorm:"comment:活动类型名称;column:name" json:"name"`
	Desc string `gorm:"comment:描述;column:desc" json:"desc"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneClassify) TableName() string {
	return `scene_classifys`
}

//
// Query
//  @Description:
//  @receiver e
//  @param db
//  @return classify
//
func (e *SceneClassify) Query(db *gorm.DB) (classify []*SceneClassify) {

	err := db.Model(e).Find(&classify).Error
	if err != nil {
		panic(err)
	}
	return classify

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param classify
//
func (e *SceneClassify) Action(db *gorm.DB, classify []*SceneClassify) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&classify).Error
	if err != nil {
		panic(err)
	}

}
