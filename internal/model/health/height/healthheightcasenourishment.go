package height

import (
	"PowerX/internal/model"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightCaseNourishment struct {
	model.Model
	Cid    string `gorm:"comment:方案ID;column:cid" json:"cid"`
	Name   string `gorm:"comment:营养补充;column:name" json:"name"`
	Number int    `gorm:"comment:数量;column:number" json:"number"`
	Units  string `gorm:"comment:单位;column:units" json:"units"`
	Remark string `gorm:"comment:备注;column:remark" json:"remark"`
	Gid    int    `gorm:"comment:固定产品型号ID;column:gid" json:"gid"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightCaseNourishment) TableName() string {
	return `heath_height_case_nourishments`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightCaseNourishment) Query(db *gorm.DB) (customer []*HeathHeightCaseNourishment) {

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
func (e *HeathHeightCaseNourishment) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightCaseNourishment) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&customer).Error
	if err != nil {
		panic(err)
	}

}
