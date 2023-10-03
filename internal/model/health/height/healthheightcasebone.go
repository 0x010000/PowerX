package height

import (
	"PowerX/internal/model"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightCaseBone struct {
	model.Model
	Cid               string `gorm:"comment:方案ID;column:cid" json:"cid"`
	ExpectHeightYear  int    `gorm:"comment:期望年身高;column:expect_height_year" json:"expect_height_year"`
	ExpectHeightMonth int    `gorm:"comment:期望月身高;column:expect_height_month" json:"expect_height_month"`
	BoneAge           int    `gorm:"comment:骨龄;column:bone_age" json:"bone_age"`
	ControlHeight     int    `gorm:"comment:控制身高cm;column:control_height" json:"control_height"`
	ControlWeight     int    `gorm:"comment:控制体重kg;column:control_weight" json:"control_weight"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightCaseBone) TableName() string {
	return `heath_height_case_bones`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightCaseBone) Query(db *gorm.DB) (customer []*HeathHeightCaseBone) {

	err := db.Model(e).Find(&customer).Error
	if err != nil {
		panic(err)
	}
	return customer

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param active
//
func (e *HeathHeightCaseBone) Action(ctx context.Context, db *gorm.DB, bone []*HeathHeightCaseBone) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&bone).Error
	if err != nil {
		panic(err)
	}

}
