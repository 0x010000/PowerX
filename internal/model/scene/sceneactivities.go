package scene

import (
	"PowerX/internal/model"
	"PowerX/internal/types"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type SceneActivities struct {
	model.Model
	Aid                 uint64    `gorm:"comment:活动ID;column:aid;unique;" json:"aid"`
	Name                string    `gorm:"comment:活动名称;column:name" json:"name"`
	Desc                string    `gorm:"comment:描述;column:desc" json:"desc"`
	Owner               string    `gorm:"comment:负责人(userId逗号隔开);column:owner" json:"owner"`
	StartTime           time.Time `gorm:"comment:开始时间;column:start_time" json:"start_time"`
	EndTime             time.Time `gorm:"comment:结束时间;column:end_time" json:"end_time"`
	ClassifyId          int       `gorm:"comment:活动分类;column:classify_id" json:"classify_id"`
	ActivitiesContentId int       `gorm:"comment:活动内容ID;column:activities_content_id" json:"activities_content_id"`
	Cpm                 int       `gorm:"comment:展示次数;column:cpm" json:"cpm"`
	State               int       `gorm:"comment:活动状态; 1:未开始 2::进行中 3:已结束 6:上架 7:结束;column:state" json:"state"`
	CoverLink           string    `gorm:"comment:活动封面Link;column:cover_link" json:"cover_link"`
	Position            string    `gorm:"comment:二维码位置:100,200;column:position" json:"position"`
	Link                string    `gorm:"comment:落地页(跳转地址);column:link" json:"link"`
	PhotoState          bool      `gorm:"comment:是否启用头像;column:photo_state" json:"photo_state"`
	AliseState          bool      `gorm:"comment:是否启用昵称;column:alise_state" json:"alise_state"`
	//
	//  PersonMaxLimit
	//  @Description:
	//
	MemberMaxLimit int `gorm:"comment:人数限制;建议不要超过最大数的500 * 0.8 = 400;column:member_max_limit;" json:"member_max_limit"`
	//
	//  ActiveParticipants
	//  @Description: active participants
	//
	ActiveParticipants []*SceneActivitiesParticipants `gorm:"foreignKey:Aid;references:Aid" json:"ActiveParticipants"`
	//
	//  ActiveContent
	//  @Description: active content
	//
	ActiveContent *SceneContent `gorm:"foreignKey:ActivitiesContentId;references:id" json:"ActiveContent"`
	//
	//  ActiveGroupQrcode
	//  @Description:
	//
	ActiveGroupQrcode []*SceneActivitiesQrcode `gorm:"foreignKey:Aid;references:Aid" json:"ActiveGroupQrcode"`
}

//
// TableName
//  @Description:
//  @receiver e
//  @return string
//
func (e SceneActivities) TableName() string {
	return `scene_activities`
}

//
// Query
//  @Description:
//  @receiver this
//  @param db
//  @return groups
//  @return err
//
func (e *SceneActivities) Query(db *gorm.DB) (active []*SceneActivities) {

	err := db.Model(e).Find(&active).Error
	if err != nil {
		panic(err)
	}
	return active

}

//
// Action
//  @Description:
//  @receiver e
//  @param db
//  @param active
//
func (e *SceneActivities) Action(db *gorm.DB, active []*SceneActivities) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "aid"}}, UpdateAll: true}).Create(&active).Error
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
func (e *SceneActivities) Page(ctx context.Context, db *gorm.DB, condition any) (count int64, activities []*SceneActivities) {

	opt := condition.(*types.PageOption[types.ActivitiesListRequest])
	query := db.Table(e.TableName()).WithContext(ctx)

	query = e.condition(query, opt.Option)

	if err := query.Count(&count).Error; err != nil {
		panic(err)
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Preload(`ActiveParticipants`).Preload(`ActiveGroupQrcode`).Order(`id DESC`).Find(&activities).Error
	if err != nil {
		panic(err)
	}
	return count, activities

}

//
// condition
//  @Description:
//  @receiver e
//  @param query
//  @param opt
//  @return *gorm.DB
//
func (e *SceneActivities) condition(query *gorm.DB, opt types.ActivitiesListRequest) *gorm.DB {

	if v := opt.Name; v != `` {
		query = query.Where(`name like ?`, "%"+v+"%")
	}
	if v := opt.State; v > 0 {
		query = query.Where(`state = ?`, v)
	}
	return query
}

//
// InspectName
//  @Description:
//  @receiver e
//  @param query
//  @param name
//  @return active
//
func (e *SceneActivities) InspectName(query *gorm.DB, name string) (active SceneActivities) {

	query = query.Table(e.TableName())
	err := query.Debug().Where(`name = ?`, name).Find(&active).Error
	if err != nil {
		panic(err)
	}
	return active
}

//
// QueryByAid
//  @Description:
//  @receiver e
//  @param query
//  @param aid
//  @return active
//
func (e *SceneActivities) QueryByAid(query *gorm.DB, aid uint64) (active SceneActivities) {

	query = query.Table(e.TableName())
	err := query.Debug().Where(`aid = ?`, aid).Find(&active).Error
	if err != nil {
		panic(err)
	}
	return active
}

func (e *SceneActivities) QueryWithNormal(db *gorm.DB) (active []*SceneActivities) {

	err := db.Model(e).Where(`start_time > ? AND end_time > ? AND state = ? `, time.Now(), time.Now(), 1).Find(&active).Error
	if err != nil {
		panic(err)
	}
	return active

}

//
// Begin
//  @Description: 开始
//  @receiver e
//  @param db
//  @param lineTime
//  @return error
//
func (e *SceneActivities) Begin(db *gorm.DB, lineTime time.Time) error {

	err := db.Model(e).Where(`created_at > ? AND start_time > ? AND end_time > ? AND state = ? `,
		lineTime,
		time.Now(),
		time.Now(),
		1).Update(`state`, 2).Error
	if err != nil {
		panic(err)
	}
	return err
}

//
// Finish
//  @Description: 结束
//  @receiver e
//  @param db
//  @param lineTime
//  @return error
//
func (e *SceneActivities) Finish(db *gorm.DB, lineTime time.Time) error {

	err := db.Model(e).Where(`created_at > ? AND end_time < ? AND state = ? `,
		lineTime,
		time.Now(),
		2).Update(`state`, 3).Error
	if err != nil {
		panic(err)
	}
	return err
}

//
// Detail
//  @Description:
//  @receiver e
//  @param ctx
//  @param db
//  @param aid
//  @return activities
//
func (e *SceneActivities) Detail(ctx context.Context, db *gorm.DB, aid uint64) (activities *SceneActivities) {

	query := db.Table(e.TableName()).WithContext(ctx).Where(`aid = ?`, aid)

	err := query.Preload(`ActiveParticipants`).Preload(`ActiveGroupQrcode.SceneQrcode`).Order(`id DESC`).Find(&activities).Error
	if err != nil {
		panic(err)
	}
	return activities

}
