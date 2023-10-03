package height

import (
	"PowerX/internal/model"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightAssessmentStatistics struct {
	model.Model
	Pid string `gorm:"comment:评估ID;column:pid;unique;" json:"pid"`

	PartHeightState     string `gorm:"comment:最近身高情况;column:height_state;" json:"height_state"`
	PartGrowthStateDay  string `gorm:"comment:近91天发育情况;column:growth_state_day;" json:"growth_state_day"`
	PartGrowthStateYear string `gorm:"comment:近一年发育情况;column:growth_state_year;" json:"growth_state_year"`

	IndexHeightValue float64 `gorm:"comment:身高;column:height_value;" json:"height_value"`
	IndexWeightValue float64 `gorm:"comment:体重;column:weight_value;" json:"weight_value"`
	IndexBMIValue    float64 `gorm:"comment:BMI;column:bmi_value;" json:"bmi_value"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightAssessmentStatistics) TableName() string {
	return `heath_height_assessment_statistics`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightAssessmentStatistics) Query(db *gorm.DB) (customer []*HeathHeightAssessment) {

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
func (e *HeathHeightAssessmentStatistics) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightAssessment) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "pid"}}, UpdateAll: true}).Create(&customer).Error
	if err != nil {
		panic(err)
	}

}
