package height

import (
	"PowerX/internal/model"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightCaseMeal struct {
	model.Model
	Cid    string `gorm:"comment:方案ID;column:cid" json:"cid"`
	Name   string `gorm:"comment:食物名称;column:name" json:"name"`
	Remark string `gorm:"comment:备注;column:remark" json:"remark"`
	Gid    int    `gorm:"comment:固定产品型号ID;column:gid" json:"gid"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightCaseMeal) TableName() string {
	return `heath_height_case_meals`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightCaseMeal) Query(db *gorm.DB) (customer []*HeathHeightCaseMeal) {

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
func (e *HeathHeightCaseMeal) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightCaseMeal) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&customer).Error
	if err != nil {
		panic(err)
	}

}
