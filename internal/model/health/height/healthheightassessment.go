package height

import (
	"PowerX/internal/model"
	"PowerX/internal/types"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightAssessment struct {
	model.Model
	Pid                 string  `gorm:"comment:评估ID;column:pid;unique;" json:"pid"`
	Oid                 string  `gorm:"comment:儿童ID;column:oid" json:"oid"`
	Analysis            string  `gorm:"comment:情况分析;column:analysis" json:"analysis"`
	Age                 int     `gorm:"comment:年龄;column:age" json:"age"`
	NowHeight           int     `gorm:"comment:现在身高CM;column:now_height" json:"now_height"`
	NowHeightCycleRatio float64 `gorm:"comment:现在身高周期比;column:now_height_cycle_ratio" json:"now_height_cycle_ratio"`
	NowWeight           int     `gorm:"comment:现在体重KG;column:now_weight" json:"now_weight"`
	NowWeightCycleRatio float64 `gorm:"comment:现在体重周期比;column:now_weight_cycle_ratio" json:"now_weight_cycle_ratio"`
	BMI                 int     `gorm:"comment:BMI;column:bmi" json:"bmi"`
	BMICycleRatio       float64 `gorm:"comment:BMI周期比;column:bmi_cycle_ratio" json:"bmi_cycle_ratio"`
	AbdominalGirth      int     `gorm:"comment:腹围CM;column:abdominal_girth" json:"abdominal_girth"`

	ExpectHeight             int     `gorm:"comment:期望身高CM;column:expect_height" json:"expect_height"`
	ExpectHeightCycleRatio   float64 `gorm:"comment:期望身高周期比;column:expect_height_cycle_ratio" json:"expect_height_cycle_ratio"`
	EstimateHeight           int     `gorm:"comment:估计身高CM;column:estimate_height" json:"estimate_height"`
	EstimateHeightCycleRatio float64 `gorm:"comment:估计身高周期比;column:estimate_height_cycle_ratio" json:"estimate_height_cycle_ratio"`
	BodyState                int     `gorm:"comment:1:形体偏低 2:匀称超重 3:轻度肥胖 4:中度肥胖 5:重度肥胖;column:body_state" json:"body_state"`
	GrowthState              int     `gorm:"comment:成长状态1:未发育 2:已发育;column:optional" json:"growth_state,optional"`
	IsFirst                  bool    `gorm:"comment:是否首次评估;column:is_first" json:"is_first"`
	RevaluateName            string  `gorm:"comment:测量人;column:revaluate_name" json:"revaluate_name"`
	RevaluateTime            int64   `gorm:"comment:评估时间;column:revaluate_time" json:"revaluate_time"`

	HeathHeightArchives *HeathHeightArchives `gorm:"foreignKey:oid;references:oid" json:"HeathHeightArchives"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightAssessment) TableName() string {
	return `heath_height_assessments`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightAssessment) Query(db *gorm.DB) (customer []*HeathHeightAssessment) {

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
func (e *HeathHeightAssessment) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightAssessment) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "pid"}}, UpdateAll: true}).Create(&customer).Error
	if err != nil {
		panic(err)
	}

}

//
// Page
//  @Description:
//  @receiver e
//  @param ctx
//  @param db
//  @param condition
//  @return count
//  @return activities
//
func (e *HeathHeightAssessment) Page(ctx context.Context, db *gorm.DB, condition any) (count int64, assessment []*HeathHeightAssessment) {

	opt := condition.(*types.PageOption[types.ListHealthHeightArchivesAssessmentRequest])
	query := db.Table(e.TableName()).WithContext(ctx)

	query = e.condition(ctx, query, opt.Option)

	if err := query.Count(&count).Error; err != nil {
		panic(err)
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Preload(`HeathHeightArchives`).Preload(`HeathHeightArchives.HeathHeightArchivesGuardian`).Order(`id DESC`).Find(&assessment).Error
	if err != nil {
		panic(err)
	}
	return count, assessment

}

//
// condition
//  @Description:
//  @receiver e
//  @param query
//  @param opt
//  @return *gorm.DB
//
func (e *HeathHeightAssessment) condition(ctx context.Context, query *gorm.DB, opt types.ListHealthHeightArchivesAssessmentRequest) *gorm.DB {

	if v := opt.Pid; v != `` {
		query = query.Where(`name = ? `, v)
	}
	if v := opt.Oid; v != `` {
		query = query.Where(`oid = ?`, v)
	}

	return query
}

//
// QueryByOid
//  @Description:
//  @receiver e
//  @param query
//  @param oid
//  @return assessment
//
func (e *HeathHeightAssessment) QueryByOid(ctx context.Context, query *gorm.DB, oid string) (assessment HeathHeightAssessment) {

	query = query.Table(e.TableName())
	err := query.Where(`oid = ?`, oid).First(&assessment).Error
	if err != nil {
		panic(err)
	}
	return assessment
}

//
// QueryByPid
//  @Description:
//  @receiver e
//  @param query
//  @param pid
//  @return assessment
//
func (e *HeathHeightAssessment) QueryByPid(ctx context.Context, query *gorm.DB, pid string) (assessment HeathHeightAssessment) {

	query = query.Table(e.TableName())
	err := query.Where(`pid = ?`, pid).First(&assessment).Error
	if err != nil {
		panic(err)
	}
	return assessment
}
