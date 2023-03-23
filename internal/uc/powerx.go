package uc

import (
	"PowerX/internal/config"
	"PowerX/internal/uc/powerx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PowerXUseCase struct {
	db          *gorm.DB
	Auth        *powerx.AuthUseCase
	Employee    *powerx.OrganizationUseCase
	Tag         *powerx.TagUseCase
	Contact     *powerx.ContactUseCase
	WeWork      *powerx.WeWorkUseCase
	MetadataCtx *powerx.MetadataCtx
}

func NewPowerXUseCase(conf *config.Config) (uc *PowerXUseCase, clean func()) {
	// 启动数据库并测试连通性
	db, err := gorm.Open(postgres.Open(conf.PowerXDatabase.DSN), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(errors.Wrap(err, "connect database failed"))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(errors.Wrap(err, "get sql db failed"))
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(errors.Wrap(err, "ping database failed"))
	}

	uc = &PowerXUseCase{
		db: db,
	}
	// 加载子UseCase
	uc.MetadataCtx = powerx.NewMetadataCtx()
	uc.Employee = powerx.NewEmployeeUseCase(db)
	uc.Auth = powerx.NewCasbinUseCase(db, uc.MetadataCtx, uc.Employee)
	uc.Department = powerx.NewDepartmentUseCase(db)
	uc.Tag = powerx.NewTagUseCase(db)
	uc.Contact = powerx.NewContactUseCase(db)
	uc.WeWork = powerx.NewWeWorkUseCase(conf, db, uc.Employee, uc.Department, uc.Auth, uc.Tag)

	uc.AutoMigrate(context.Background())
	uc.AutoInit()

	return uc, func() {
		_ = sqlDB.Close()
	}
}

func (p *PowerXUseCase) AutoMigrate(ctx context.Context) {
	p.db.AutoMigrate(&powerx.CasbinPolicy{}, &powerx.AuthRole{}, &powerx.AuthRestAction{}, &powerx.AuthRecourse{})
	p.db.AutoMigrate(&powerx.Department{}, &powerx.Employee{}, &powerx.LiveQRCode{})
	p.db.AutoMigrate(&powerx.WeWorkDepartment{}, &powerx.WeWorkEmployee{})
}

func (p *PowerXUseCase) AutoInit() {
	p.Auth.Init()
	p.Department.Init()
}