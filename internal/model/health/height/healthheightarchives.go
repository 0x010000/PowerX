package height

import (
	"PowerX/internal/model"
	"PowerX/internal/module/auth"
	"PowerX/internal/types"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HeathHeightArchives struct {
	model.Model
	Oid                         string                         `gorm:"comment:档案号;column:oid;unique;" json:"oid"`
	OrgId                       int                            `gorm:"comment:组织机构ID;column:org_id" json:"org_id"`
	UserId                      string                         `gorm:"comment:医生ID;column:user_id" json:"user_id"`
	Name                        string                         `gorm:"comment:名称;column:name" json:"name"`
	Desc                        string                         `gorm:"comment:描述;column:desc" json:"desc"`
	ExternalUserId              string                         `gorm:"comment:客户ID;column:external_user_id" json:"external_user_id"`
	Mobile                      string                         `gorm:"comment:手机号;column:mobile" json:"mobile"`
	FatherHeight                int                            `gorm:"comment:父亲身高;column:father_height" json:"father_height"`
	FatherWeight                int                            `gorm:"comment:父亲体重;column:father_weight" json:"father_weight"`
	MotherHeight                int                            `gorm:"comment:母亲身高;column:mother_height" json:"mother_height"`
	MotherWeight                int                            `gorm:"comment:母亲体重;column:mother_weight" json:"mother_weight"`
	Gender                      int                            `gorm:"comment:性别1: 男 2: 女 3: 保密;column:gender" json:"gender"`
	Age                         float64                        `gorm:"comment:年龄;column:age" json:"age"`
	Birth                       int64                          `gorm:"comment:出生时间;column:birth" json:"birth"`
	Weight                      int                            `gorm:"comment:体重KG;column:weight" json:"weight"`
	Height                      int                            `gorm:"comment:身高CM;column:height" json:"height"`
	GestationalWeeks            int                            `gorm:"comment:孕周;column:gestational_weeks" json:"gestational_weeks"`
	PrevAssessmentTime          uint64                         `gorm:"comment:上一次诊断时间;column:prev_assessment_time" json:"prev_assessment_time"`
	NextAssessmentTime          uint64                         `gorm:"comment:下一次复诊时间;column:next_assessment_time" json:"next_assessment_time"`
	LastAssessmentId            int                            `gorm:"comment:最后评估报告ID;column:last_assessment_id" json:"last_assessment_id"`
	LastPid                     string                         `gorm:"comment:最后评估报告ID;column:last_pid" json:"last_pid"`
	Portrait                    string                         `gorm:"comment:头像;column:portrait" json:"portrait"`
	State                       int                            `gorm:"comment:状态1:启用 2：禁用;column:state" json:"state"`
	HeathHeightArchivesGuardian []*HeathHeightArchivesGuardian `gorm:"foreignKey:oid;references:oid" json:"HeathHeightArchivesGuardian"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e HeathHeightArchives) TableName() string {
	return `heath_height_archives`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *HeathHeightArchives) Query(ctx context.Context, db *gorm.DB) (customer []*HeathHeightArchives) {

	db = db.Model(e).WithContext(ctx)
	if v := auth.Authorization(ctx); v.App != nil && v.AID != `` {
		db = db.Where(`user_id = ?`, v)
	}
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
func (e *HeathHeightArchives) Action(ctx context.Context, db *gorm.DB, customer []*HeathHeightArchives) {

	err := db.Table(e.TableName()).WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "oid"}}, UpdateAll: true}).Create(&customer).Error
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
func (e *HeathHeightArchives) Page(ctx context.Context, db *gorm.DB, condition any) (count int64, customers []*HeathHeightArchives) {

	opt := condition.(*types.PageOption[types.HealthHeightArchivesListRequest])
	query := db.Table(e.TableName()).WithContext(ctx)

	query = e.condition(ctx, query, opt.Option)

	if err := query.Count(&count).Error; err != nil {
		panic(err)
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Preload(`HeathHeightArchivesGuardian`).Order(`id DESC`).Find(&customers).Error
	if err != nil {
		panic(err)
	}
	return count, customers

}

//
// condition
//  @Description:
//  @receiver e
//  @param query
//  @param opt
//  @return *gorm.DB
//
func (e *HeathHeightArchives) condition(ctx context.Context, query *gorm.DB, opt types.HealthHeightArchivesListRequest) *gorm.DB {

	if v := opt.ExternalUserId; v != `` {
		query = query.Where(`external_user_id = ?`, v)
	}
	if v := opt.UserId; v != `` {
		query = query.Where(`user_id = ?`, v)
	}
	if v := opt.OrgId; v > 0 {
		query = query.Where(`org_id = ?`, v)
	}
	if v := opt.Gender; v > 0 {
		query = query.Where(`gender = ?`, v)
	}
	if v := opt.Mobile; v != `` {
		query = query.Where(`mobile like ?`, "%"+v+"%")
	}
	if v := opt.Oid; v != nil {
		query = query.Where(`oid IN (?) `, v)
	}
	if v := opt.NextAssessmentTime; v != nil {
		query = query.Where(`next_assessment_time IN (?) `, v)
	}
	if v := auth.Authorization(ctx); v.App != nil && v.AID != `` {
		query = query.Where(`user_id = ?`, v)
	}
	return query
}

//
// InspectName
//  @Description:
//  @receiver e
//  @param query
//  @param name
//  @return customer
//
func (e *HeathHeightArchives) InspectName(query *gorm.DB, name string) (customer HeathHeightArchives) {

	query = query.Table(e.TableName())
	err := query.Where(`name = ?`, name).Find(&customer).Error
	if err != nil {
		panic(err)
	}
	return customer
}

//
// QueryByOid
//  @Description:
//  @receiver e
//  @param ctx
//  @param query
//  @param oid
//  @return archives
//
func (e *HeathHeightArchives) QueryByOid(ctx context.Context, query *gorm.DB, oid string) (archives HeathHeightArchives) {

	query = query.Table(e.TableName())

	err := query.Where(`oid = ?`, oid).First(&archives).Error
	if err != nil {
		panic(err)
	}
	return archives
}
