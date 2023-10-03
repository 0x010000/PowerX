package common

import (
	"PowerX/internal/model/health/height"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHealthHeightStandardOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHealthHeightStandardOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHealthHeightStandardOptionLogic {
	return &ListHealthHeightStandardOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListHealthHeightStandardOption
//  @Description:
//  @receiver l
//  @param opt
//  @return resp
//  @return err
//
func (l *ListHealthHeightStandardOptionLogic) ListHealthHeightStandardOption(opt *types.HealthHeightStandardListRequest) (resp *types.HealthHeightStandardOptionReply, err error) {
	// todo: add your logic here and delete this line
	option, err := l.svcCtx.Custom.Health.FindManyHeightStandardOption(l.ctx, opt)
	return &types.HealthHeightStandardOptionReply{
		List: l.OPT(option),
	}, err
}

//
// OPT
//  @Description:
//  @receiver l
//  @param stature
//  @return standard
//
func (l *ListHealthHeightStandardOptionLogic) OPT(stature []*height.HeathHeightStandardStature) (standard []*types.HealthHeightStandardListReply) {

	if stature != nil {
		for _, val := range stature {
			standard = append(standard, &types.HealthHeightStandardListReply{
				Age:    val.Age,
				Old:    val.Old,
				Gender: val.Gender,
				HP03:   val.HP03,
				HP10:   val.HP10,
				HP25:   val.HP25,
				HP50:   val.HP50,
				HP75:   val.HP75,
				HP90:   val.HP90,
				HP97:   val.HP97,
				WP03:   val.WP03,
				WP10:   val.WP10,
				WP25:   val.WP25,
				WP50:   val.WP50,
				WP75:   val.WP75,
				WP90:   val.WP90,
				WP97:   val.WP97,
				BP03:   val.BP03,
				BP10:   val.BP10,
				BP25:   val.BP25,
				BP50:   val.BP50,
				BP75:   val.BP75,
				BP90:   val.BP90,
				BP97:   val.BP97,
				HM:     val.HM,
				HL:     val.HL,
			})
		}
	}
	return standard
}
