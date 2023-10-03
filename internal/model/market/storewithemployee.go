package market

import (
	"PowerX/internal/model"
	"PowerX/internal/module/auth"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PivotStoreToEmployee struct {
	model.Model

	StoreId    int64  `gorm:"comment:店铺ID;column:store_id" json:"store_id"`
	EmployeeId int64  `gorm:"comment:员工Id;column:employee_id" json:"employee_id"`
	UserId     string `gorm:"comment:员工Id;column:user_id" json:"user_id"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e PivotStoreToEmployee) TableName() string {
	return `pivot_store_to_employees`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *PivotStoreToEmployee) Query(db *gorm.DB) (customer []*PivotStoreToEmployee) {

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
func (e *PivotStoreToEmployee) Action(ctx context.Context, db *gorm.DB, store []*PivotStoreToEmployee) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&store).Error
	if err != nil {
		panic(err)
	}

}

//
// QueryByAuthored
//  @Description:
//  @receiver e
//  @param ctx
//  @param db
//  @return store
//
func (e *PivotStoreToEmployee) QueryByAuthored(ctx context.Context, db *gorm.DB) (store *PivotStoreToEmployee) {
	db = db.Model(e).WithContext(ctx)
	if v := auth.Authorization(ctx); v != nil {
		db = db.Where(`employee_id = ?`, v.AID)
	}
	err := db.Find(&store).Error
	if err != nil {
		panic(err)
	}
	return store

}
