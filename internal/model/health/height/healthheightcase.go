package height

import (
	"PowerX/internal/model"
	"PowerX/internal/types"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightCase struct {
	model.Model
	Cid               string `gorm:"comment:方案ID;column:cid;unique;" json:"cid"`
	Oid               string `gorm:"comment:档案号;column:oid" json:"oid"`
	Pid               string `gorm:"comment:评估ID;column:pid" json:"pid"`
	ControlLess       string `gorm:"comment:少使用;column:control_less" json:"control_less"`
	ControlNonuse     string `gorm:"comment:不使用;column:control_nonuse" json:"control_nonuse"`
	ControlBreakfast  string `gorm:"comment:早餐;column:control_breakfast" json:"control_breakfast"`
	ControlLunch      string `gorm:"comment:午餐;column:control_lunch" json:"control_lunch"`
	ControlDinner     string `gorm:"comment:晚餐;column:control_dinner" json:"control_dinner"`
	ControlMealMinute string `gorm:"comment:进餐时长;column:control_meal_minute" json:"control_meal_minute"`
	SportRemark       string `gorm:"comment:运动建议;column:sport_remark" json:"sport_remark"`
	SleepCase         string `gorm:"comment:睡眠方案;column:sleep_case" json:"sleep_case"`
	EmotionCase       string `gorm:"comment:情绪管理;column:emotion_case" json:"emotion_case"`
	Diagnostic        string `gorm:"comment:诊断情况;column:diagnostic" json:"diagnostic"`

	HeathHeightCaseBone        []*HeathHeightCaseBone        `gorm:"foreignKey:cid;references:cid" json:"HeathHeightCaseBone"`
	HeathHeightCaseNourishment []*HeathHeightCaseNourishment `gorm:"foreignKey:cid;references:cid" json:"HeathHeightCaseNourishment"`
	HeathHeightCaseMeal        []*HeathHeightCaseMeal        `gorm:"foreignKey:cid;references:cid" json:"HeathHeightCaseMeal"`
	HeathHeightCaseSports      []*HeathHeightCaseSports      `gorm:"foreignKey:cid;references:cid" json:"HeathHeightCaseSports"`
	HeathHeightArchives        *HeathHeightArchives          `gorm:"foreignKey:oid;references:oid" json:"HeathHeightArchives"`
	HeathHeightAssessment      *HeathHeightAssessment        `gorm:"foreignKey:pid;references:pid" json:"HeathHeightAssessment"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightCase) TableName() string {
	return `heath_height_cases`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightCase) Query(db *gorm.DB) (customer []*HeathHeightCase) {

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
func (e *HeathHeightCase) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightCase) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, UpdateAll: true}).Create(&customer).Error
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
//  @return customers
//
func (e *HeathHeightCase) Page(ctx context.Context, db *gorm.DB, condition any) (count int64, cases []*HeathHeightCase) {

	opt := condition.(*types.PageOption[types.HealthHeightCaseListRequest])
	query := db.Table(e.TableName()).WithContext(ctx)

	query = e.condition(ctx, query, opt.Option)

	if err := query.Count(&count).Error; err != nil {
		panic(err)
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.
		Preload(`HeathHeightCaseBone`).
		Preload(`HeathHeightCaseNourishment`).
		Preload(`HeathHeightCaseMeal`).
		Preload(`HeathHeightCaseSports`).
		Preload(`HeathHeightArchives`).
		Preload(`HeathHeightArchives.HeathHeightArchivesGuardian`).
		Preload(`HeathHeightAssessment`).
		Order(`id DESC`).Find(&cases).Error
	if err != nil {
		panic(err)
	}
	return count, cases

}

//
// condition
//  @Description:
//  @receiver e
//  @param query
//  @param opt
//  @return *gorm.DB
//
func (e *HeathHeightCase) condition(ctx context.Context, query *gorm.DB, opt types.HealthHeightCaseListRequest) *gorm.DB {

	if v := opt.Oid; v != `` {
		query = query.Where(`oid = ?`, v)
	}
	if v := opt.Pid; v != `` {
		query = query.Where(`pid = ?`, v)
	}
	if v := opt.Cid; v != `` {
		query = query.Where(`cid = ?`, v)
	}

	return query
}

//
// QueryByCid
//  @Description:
//  @receiver e
//  @param ctx
//  @param query
//  @param cid
//  @return cases
//
func (e *HeathHeightCase) QueryByCid(ctx context.Context, query *gorm.DB, cid string) (cases HeathHeightCase) {

	query = query.Table(e.TableName())
	err := query.Where(`cid = ?`, cid).First(&cases).Error
	if err != nil {
		panic(err)
	}
	return cases
}
