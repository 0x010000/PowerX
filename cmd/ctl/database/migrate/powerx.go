package migrate

import (
	"PowerX/cmd/ctl/database/custom/migrate"
	"PowerX/internal/config"
	"PowerX/internal/model"
	"PowerX/internal/model/customerdomain"
	"PowerX/internal/model/health/height"
	"PowerX/internal/model/market"
	"PowerX/internal/model/media"
	"PowerX/internal/model/membership"
	"PowerX/internal/model/origanzation"
	"PowerX/internal/model/permission"
	"PowerX/internal/model/product"
	"PowerX/internal/model/scene"
	"PowerX/internal/model/scrm/app"
	"PowerX/internal/model/scrm/customer"
	"PowerX/internal/model/scrm/organization"
	"PowerX/internal/model/scrm/resource"
	"PowerX/internal/model/scrm/tag"
	"PowerX/internal/model/trade"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type PowerMigrator struct {
	db *gorm.DB
}

func NewPowerMigrator(conf *config.Config) (*PowerMigrator, error) {
	var dsn gorm.Dialector
	switch conf.PowerXDatabase.Driver {
	case config.DriverMysql:
		dsn = mysql.Open(conf.PowerXDatabase.DSN)
	case config.DriverPostgres:
		dsn = postgres.Open(conf.PowerXDatabase.DSN)
	}
	db, err := gorm.Open(dsn, &gorm.Config{
		//Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return &PowerMigrator{
		db: db,
	}, err
}

func (m *PowerMigrator) AutoMigrate() {
	//m.Excel()

	_ = m.db.AutoMigrate(&model.DataDictionaryType{}, &model.DataDictionaryItem{}, &model.PivotDataDictionaryToObject{})
	_ = m.db.AutoMigrate(&origanzation.Department{}, &origanzation.Employee{}, &origanzation.Position{})
	_ = m.db.AutoMigrate(&permission.EmployeeCasbinPolicy{}, permission.AdminRole{}, permission.AdminRoleMenuName{}, permission.AdminAPI{})

	// customer domain
	_ = m.db.AutoMigrate(&customerdomain.Lead{}, &customerdomain.Contact{}, &customerdomain.Customer{}, &membership.Membership{})
	_ = m.db.AutoMigrate(&model.WechatOACustomer{}, &model.WechatMPCustomer{}, &model.WeWorkExternalContact{})
	_ = m.db.AutoMigrate(
		&product.PivotProductToProductCategory{},
	)
	// product
	_ = m.db.AutoMigrate(&product.Product{}, &product.ProductCategory{})
	_ = m.db.AutoMigrate(&product.ProductSpecific{}, &product.SpecificOption{})
	_ = m.db.AutoMigrate(&product.SKU{}, &product.PivotSkuToSpecificOption{})
	_ = m.db.AutoMigrate(&product.PriceBook{}, &product.PriceBookEntry{}, &product.PriceConfig{})
	_ = m.db.AutoMigrate(&market.Store{}, &product.Artisan{}, &product.PivotStoreToArtisan{})
	_ = m.db.AutoMigrate(&market.PivotStoreToEmployee{})

	// market
	_ = m.db.AutoMigrate(&market.Media{})

	// media
	_ = m.db.AutoMigrate(&media.MediaResource{}, &media.PivotMediaResourceToObject{})

	// trade
	_ = m.db.AutoMigrate(&trade.ShippingAddress{}, &trade.DeliveryAddress{}, &trade.BillingAddress{})
	_ = m.db.AutoMigrate(&trade.Warehouse{}, &trade.Inventory{}, &trade.Logistics{})
	_ = m.db.AutoMigrate(&trade.Cart{}, &trade.CartItem{}, &trade.Order{}, &trade.OrderItem{})
	_ = m.db.AutoMigrate(&trade.OrderStatusTransition{}, &trade.PivotOrderToInventoryLog{})
	_ = m.db.AutoMigrate(&trade.Payment{}, &trade.PaymentItem{})
	_ = m.db.AutoMigrate(&trade.RefundOrder{}, &trade.RefundOrderItem{})
	_ = m.db.AutoMigrate(&trade.TokenBalance{}, &trade.ExchangeRatio{}, &trade.ExchangeRecord{})

	// custom
	migrate.AutoMigrateCustom(m.db)

	// wechat organization
	_ = m.db.AutoMigrate(&organization.WeWorkEmployee{}, &organization.WeWorkDepartment{})
	// wechat customer
	_ = m.db.AutoMigrate(&customer.WeWorkExternalContacts{}, &customer.WeWorkExternalContactFollow{})
	// wechat resource
	_ = m.db.AutoMigrate(&resource.WeWorkResource{})
	// wechat app
	_ = m.db.AutoMigrate(&app.WeWorkAppGroup{})
	// wechat tag
	_ = m.db.AutoMigrate(&tag.WeWorkTag{}, &tag.WeWorkTagGroup{})
	// scene
	_ = m.db.AutoMigrate(&scene.SceneQrcode{}, &scene.SceneClassify{})
	_ = m.db.AutoMigrate(&scene.SceneActivities{}, &scene.SceneActivitiesParticipants{}, &scene.SceneActivitiesQrcode{})
	_ = m.db.AutoMigrate(&scene.SceneContent{})
	// health
	_ = m.db.AutoMigrate(&height.HeathHeightArchives{}, &height.HeathHeightArchivesGuardian{})
	_ = m.db.AutoMigrate(&height.HeathHeightAssessment{})
	_ = m.db.AutoMigrate(&height.HeathHeightCase{}, &height.HeathHeightCaseBone{}, &height.HeathHeightCaseSports{}, &height.HeathHeightCaseNourishment{}, &height.HeathHeightCaseMeal{})
	_ = m.db.AutoMigrate(&height.HeathHeightStandardStature{})
	//m.Excel()

}

func (m *PowerMigrator) Excel() {

	file, err := xlsx.OpenFile(`2.xlsx`)
	if err != nil {
		panic(err)
	}
	var boy []*height.HeathHeightStandardStature
	var girl []*height.HeathHeightStandardStature
	sh, _ := file.Sheet["身高体重"]
	for i := 4; i < sh.MaxRow; i++ {
		var ages int
		age, _ := sh.Cell(i, 0)
		if age.Value == `` {
			break
		}
		if age.Value == `出生` {
			ages = 1
		} else if strings.Contains(age.Value, `月`) {
			ages, _ = strconv.Atoi(strings.ReplaceAll(age.Value, `月`, ``))
		} else {
			var y, m int

			tmp := strings.ReplaceAll(age.Value, `岁`, ``)
			split := strings.Split(tmp, `.`)
			if len(split) > 1 {
				y, _ = strconv.Atoi(split[0])
				m, _ = strconv.Atoi(split[1])
				ages = y*12 + m
			} else {
				y, _ = strconv.Atoi(split[0])
				ages = y * 12
			}
			fmt.Println(ages)

		}
		pbh103, _ := sh.Cell(i, 1)
		pbh110, _ := sh.Cell(i, 2)
		pbh125, _ := sh.Cell(i, 3)
		pbh150, _ := sh.Cell(i, 4)
		pbh175, _ := sh.Cell(i, 5)
		pbh190, _ := sh.Cell(i, 6)
		pbh197, _ := sh.Cell(i, 7)
		pbw103, _ := sh.Cell(i, 8)
		pbw110, _ := sh.Cell(i, 9)
		pbw125, _ := sh.Cell(i, 10)
		pbw150, _ := sh.Cell(i, 11)
		pbw175, _ := sh.Cell(i, 12)
		pbw190, _ := sh.Cell(i, 13)
		pbw197, _ := sh.Cell(i, 14)
		//
		tmp_pbh103, _ := strconv.ParseFloat(pbh103.Value, 32)
		tmp_pbh110, _ := strconv.ParseFloat(pbh110.Value, 32)
		tmp_pbh125, _ := strconv.ParseFloat(pbh125.Value, 32)
		tmp_pbh150, _ := strconv.ParseFloat(pbh150.Value, 32)
		tmp_pbh175, _ := strconv.ParseFloat(pbh175.Value, 32)
		tmp_pbh190, _ := strconv.ParseFloat(pbh190.Value, 32)
		tmp_pbh197, _ := strconv.ParseFloat(pbh197.Value, 32)
		tmp_pbw103, _ := strconv.ParseFloat(pbw103.Value, 32)
		tmp_pbw110, _ := strconv.ParseFloat(pbw110.Value, 32)
		tmp_pbw125, _ := strconv.ParseFloat(pbw125.Value, 32)
		tmp_pbw150, _ := strconv.ParseFloat(pbw150.Value, 32)
		tmp_pbw175, _ := strconv.ParseFloat(pbw175.Value, 32)
		tmp_pbw190, _ := strconv.ParseFloat(pbw190.Value, 32)
		tmp_pbw197, _ := strconv.ParseFloat(pbw197.Value, 32)
		//
		b_pbh103, _ := decimal.NewFromFloat(tmp_pbh103 / 100).Round(2).Float64()
		b_pbh110, _ := decimal.NewFromFloat(tmp_pbh110 / 100).Round(2).Float64()
		b_pbh125, _ := decimal.NewFromFloat(tmp_pbh125 / 100).Round(2).Float64()
		b_pbh150, _ := decimal.NewFromFloat(tmp_pbh150 / 100).Round(2).Float64()
		b_pbh175, _ := decimal.NewFromFloat(tmp_pbh175 / 100).Round(2).Float64()
		b_pbh190, _ := decimal.NewFromFloat(tmp_pbh190 / 100).Round(2).Float64()
		b_pbh197, _ := decimal.NewFromFloat(tmp_pbh197 / 100).Round(2).Float64()
		b_pbw103, _ := decimal.NewFromFloat(tmp_pbw103).Round(2).Float64()
		b_pbw110, _ := decimal.NewFromFloat(tmp_pbw110).Round(2).Float64()
		b_pbw125, _ := decimal.NewFromFloat(tmp_pbw125).Round(2).Float64()
		b_pbw150, _ := decimal.NewFromFloat(tmp_pbw150).Round(2).Float64()
		b_pbw175, _ := decimal.NewFromFloat(tmp_pbw175).Round(2).Float64()
		b_pbw190, _ := decimal.NewFromFloat(tmp_pbw190).Round(2).Float64()
		b_pbw197, _ := decimal.NewFromFloat(tmp_pbw197).Round(2).Float64()
		//

		bmi03, _ := decimal.NewFromFloat(b_pbw103 / (b_pbh103 * b_pbh103)).Round(2).Float64()
		bmi10, _ := decimal.NewFromFloat(b_pbw110 / (b_pbh110 * b_pbh110)).Round(2).Float64()
		bmi25, _ := decimal.NewFromFloat(b_pbw125 / (b_pbh125 * b_pbh125)).Round(2).Float64()
		bmi50, _ := decimal.NewFromFloat(b_pbw150 / (b_pbh150 * b_pbh150)).Round(2).Float64()
		bmi75, _ := decimal.NewFromFloat(b_pbw175 / (b_pbh175 * b_pbh175)).Round(2).Float64()
		bmi90, _ := decimal.NewFromFloat(b_pbw190 / (b_pbh190 * b_pbh190)).Round(2).Float64()
		bmi97, _ := decimal.NewFromFloat(b_pbw197 / (b_pbh197 * b_pbh197)).Round(2).Float64()
		boy = append(boy, &height.HeathHeightStandardStature{

			Age:    ages,
			Old:    age.Value,
			Gender: 1,
			HP03:   b_pbh103,
			HP10:   b_pbh110,
			HP25:   b_pbh125,
			HP50:   b_pbh150,
			HP75:   b_pbh175,
			HP90:   b_pbh190,
			HP97:   b_pbh197,
			WP03:   b_pbw103,
			WP10:   b_pbw110,
			WP25:   b_pbw125,
			WP50:   b_pbw150,
			WP75:   b_pbw175,
			WP90:   b_pbw190,
			WP97:   b_pbw197,
			BP03:   bmi03,
			BP10:   bmi10,
			BP25:   bmi25,
			BP50:   bmi50,
			BP75:   bmi75,
			BP90:   bmi90,
			BP97:   bmi97,
			HM:     0,
			HL:     0,
		})
		gbh103, _ := sh.Cell(i, 15)
		gbh110, _ := sh.Cell(i, 16)
		gbh125, _ := sh.Cell(i, 17)
		gbh150, _ := sh.Cell(i, 18)
		gbh175, _ := sh.Cell(i, 19)
		gbh190, _ := sh.Cell(i, 20)
		gbh197, _ := sh.Cell(i, 21)
		gbw103, _ := sh.Cell(i, 22)
		gbw110, _ := sh.Cell(i, 23)
		gbw125, _ := sh.Cell(i, 24)
		gbw150, _ := sh.Cell(i, 25)
		gbw175, _ := sh.Cell(i, 26)
		gbw190, _ := sh.Cell(i, 27)
		gbw197, _ := sh.Cell(i, 28)
		tmp_gbh103, _ := strconv.ParseFloat(gbh103.Value, 32)
		tmp_gbh110, _ := strconv.ParseFloat(gbh110.Value, 32)
		tmp_gbh125, _ := strconv.ParseFloat(gbh125.Value, 32)
		tmp_gbh150, _ := strconv.ParseFloat(gbh150.Value, 32)
		tmp_gbh175, _ := strconv.ParseFloat(gbh175.Value, 32)
		tmp_gbh190, _ := strconv.ParseFloat(gbh190.Value, 32)
		tmp_gbh197, _ := strconv.ParseFloat(gbh197.Value, 32)
		tmp_gbw103, _ := strconv.ParseFloat(gbw103.Value, 32)
		tmp_gbw110, _ := strconv.ParseFloat(gbw110.Value, 32)
		tmp_gbw125, _ := strconv.ParseFloat(gbw125.Value, 32)
		tmp_gbw150, _ := strconv.ParseFloat(gbw150.Value, 32)
		tmp_gbw175, _ := strconv.ParseFloat(gbw175.Value, 32)
		tmp_gbw190, _ := strconv.ParseFloat(gbw190.Value, 32)
		tmp_gbw197, _ := strconv.ParseFloat(gbw197.Value, 32)
		g_gbh103, _ := decimal.NewFromFloat(tmp_gbh103 / 100).Round(2).Float64()
		g_gbh110, _ := decimal.NewFromFloat(tmp_gbh110 / 100).Round(2).Float64()
		g_gbh125, _ := decimal.NewFromFloat(tmp_gbh125 / 100).Round(2).Float64()
		g_gbh150, _ := decimal.NewFromFloat(tmp_gbh150 / 100).Round(2).Float64()
		g_gbh175, _ := decimal.NewFromFloat(tmp_gbh175 / 100).Round(2).Float64()
		g_gbh190, _ := decimal.NewFromFloat(tmp_gbh190 / 100).Round(2).Float64()
		g_gbh197, _ := decimal.NewFromFloat(tmp_gbh197 / 100).Round(2).Float64()
		g_gbw103, _ := decimal.NewFromFloat(tmp_gbw103).Round(2).Float64()
		g_gbw110, _ := decimal.NewFromFloat(tmp_gbw110).Round(2).Float64()
		g_gbw125, _ := decimal.NewFromFloat(tmp_gbw125).Round(2).Float64()
		g_gbw150, _ := decimal.NewFromFloat(tmp_gbw150).Round(2).Float64()
		g_gbw175, _ := decimal.NewFromFloat(tmp_gbw175).Round(2).Float64()
		g_gbw190, _ := decimal.NewFromFloat(tmp_gbw190).Round(2).Float64()
		g_gbw197, _ := decimal.NewFromFloat(tmp_gbw197).Round(2).Float64()
		//
		//
		gbmi03, _ := decimal.NewFromFloat(g_gbw103 / (g_gbh103 * g_gbh103)).Round(2).Float64()
		gbmi10, _ := decimal.NewFromFloat(g_gbw110 / (g_gbh110 * g_gbh110)).Round(2).Float64()
		gbmi25, _ := decimal.NewFromFloat(g_gbw125 / (g_gbh125 * g_gbh125)).Round(2).Float64()
		gbmi50, _ := decimal.NewFromFloat(g_gbw150 / (g_gbh150 * g_gbh150)).Round(2).Float64()
		gbmi75, _ := decimal.NewFromFloat(g_gbw175 / (g_gbh175 * g_gbh175)).Round(2).Float64()
		gbmi90, _ := decimal.NewFromFloat(g_gbw190 / (g_gbh190 * g_gbh190)).Round(2).Float64()
		gbmi97, _ := decimal.NewFromFloat(g_gbw197 / (g_gbh197 * g_gbh197)).Round(2).Float64()
		girl = append(girl, &height.HeathHeightStandardStature{
			Age:    ages,
			Old:    age.Value,
			Gender: 2,
			HP03:   g_gbh103,
			HP10:   g_gbh110,
			HP25:   g_gbh125,
			HP50:   g_gbh150,
			HP75:   g_gbh175,
			HP90:   g_gbh190,
			HP97:   g_gbh197,
			WP03:   g_gbw103,
			WP10:   g_gbw110,
			WP25:   g_gbw125,
			WP50:   g_gbw150,
			WP75:   g_gbw175,
			WP90:   g_gbw190,
			WP97:   g_gbw197,
			BP03:   gbmi03,
			BP10:   gbmi10,
			BP25:   gbmi25,
			BP50:   gbmi50,
			BP75:   gbmi75,
			BP90:   gbmi90,
			BP97:   gbmi97,
			HM:     0,
			HL:     0,
		})

	}
	tab := height.HeathHeightStandardStature{}
	tab.Action(context.TODO(), m.db, boy)
	tab.Action(context.TODO(), m.db, girl)

}
