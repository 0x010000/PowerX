package height

import (
	"PowerX/internal/model"
	"PowerX/internal/types"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightStandardStature struct {
	model.Model
	Age    int    `gorm:"comment:年龄;column:age;" json:"age"`
	Old    string `gorm:"comment:年龄;column:old;" json:"old"`
	Gender int    `gorm:"comment:性别;column:gender;" json:"gender"`
	///
	HP03 float64 `gorm:"comment:身高;column:hp03;" json:"hp03"`
	HP10 float64 `gorm:"comment:身高;column:hp10;" json:"hp10"`
	HP25 float64 `gorm:"comment:身高;column:hp25;" json:"hp25"`
	HP50 float64 `gorm:"comment:身高;column:hp50;" json:"hp50"`
	HP75 float64 `gorm:"comment:身高;column:hp75;" json:"hp75"`
	HP90 float64 `gorm:"comment:身高;column:hp90;" json:"hp90"`
	HP97 float64 `gorm:"comment:身高;column:hp97;" json:"hp97"`
	///
	WP03 float64 `gorm:"comment:身高;column:wp03;" json:"wp03"`
	WP10 float64 `gorm:"comment:体重;column:wp10;" json:"wp10"`
	WP25 float64 `gorm:"comment:体重;column:wp25;" json:"wp25"`
	WP50 float64 `gorm:"comment:体重;column:wp50;" json:"wp50"`
	WP75 float64 `gorm:"comment:体重;column:wp75;" json:"wp75"`
	WP90 float64 `gorm:"comment:体重;column:wp90;" json:"wp90"`
	WP97 float64 `gorm:"comment:体重;column:wp97;" json:"wp97"`
	//
	BP03 float64 `gorm:"comment:BMI;column:bp03;" json:"bp03"`
	BP10 float64 `gorm:"comment:BMI;column:bp10;" json:"bp10"`
	BP25 float64 `gorm:"comment:BMI;column:bp25;" json:"bp25"`
	BP50 float64 `gorm:"comment:BMI;column:bp50;" json:"bp50"`
	BP75 float64 `gorm:"comment:BMI;column:bp75;" json:"bp75"`
	BP90 float64 `gorm:"comment:BMI;column:bp90;" json:"bp90"`
	BP97 float64 `gorm:"comment:BMI;column:bp97;" json:"bp97"`
	///
	HM float64 `gorm:"comment:身高平均值;column:hm;" json:"hm"`
	HL float64 `gorm:"comment:身高偏差值;column:hl;" json:"hl"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightStandardStature) TableName() string {
	return `heath_height_standard__statures`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightStandardStature) Query(ctx context.Context, db *gorm.DB) (customer []*HeathHeightStandardStature) {

	db = db.Model(e).WithContext(ctx)
	err := db.Find(&customer).Error
	if err != nil {
		panic(err)
	}
	return customer

}

//
// Action
//  @Description:
//  @receiver e
//  @param ctx
//  @param db
//  @param customer
//
func (e *HeathHeightStandardStature) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightStandardStature) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).CreateInBatches(&customer, 100).Error
	if err != nil {
		panic(err)
	}

}

//
// Options
//  @Description:
//  @receiver e
//  @param ctx
//  @param db
//  @param request
//  @return standard
//
func (e *HeathHeightStandardStature) Options(ctx context.Context, db *gorm.DB, request *types.HealthHeightStandardListRequest) (standard []*HeathHeightStandardStature) {

	db = db.Table(e.TableName()).WithContext(ctx)
	if v := request.Gender; v > 0 {
		db = db.Where(`gender = ?`, v)
	}
	err := db.Find(&standard).Error
	if err != nil {
		panic(err)
	}
	return standard

}
