package height

import (
	"PowerX/internal/model"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightArchivesGuardian struct {
	model.Model
	Oid      string `gorm:"comment:档案号;column:oid;" json:"oid"`
	Name     string `gorm:"comment:监护人姓名;column:name" json:"name"`
	Relation string `gorm:"comment:与被监护人关系;column:relation" json:"relation"`
	Mobile   string `gorm:"comment:监护人手机号;column:mobile" json:"mobile"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightArchivesGuardian) TableName() string {
	return `heath_height_archives_guardians`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightArchivesGuardian) Query(db *gorm.DB) (customer []*HeathHeightArchivesGuardian) {

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
//  @param ctx
//  @param db
//  @param customer
//
func (e *HeathHeightArchivesGuardian) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightArchivesGuardian) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&customer).Error
	if err != nil {
		panic(err)
	}

}
