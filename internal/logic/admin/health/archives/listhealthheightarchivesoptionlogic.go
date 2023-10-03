package archives

import (
	"PowerX/internal/model/health/height"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHealthHeightArchivesOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHealthHeightArchivesOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHealthHeightArchivesOptionLogic {
	return &ListHealthHeightArchivesOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (archive *ListHealthHeightArchivesOptionLogic) ListHealthHeightArchivesOption() (resp *types.HealthHeightArchivesOptionReply, err error) {

	option, err := archive.svcCtx.Custom.Health.FindManyHeightArchivesOption(archive.ctx)

	return &types.HealthHeightArchivesOptionReply{
		List: archive.DTO(option),
	}, err

}

//
// DTO
//  @Description:
//  @receiver archive
//  @param archives
//  @return option
//
func (archive *ListHealthHeightArchivesOptionLogic) DTO(archives []*height.HeathHeightArchives) (option []*types.Option) {

	if archives != nil {
		for _, val := range archives {
			option = append(option, &types.Option{
				Id:   val.Oid,
				Name: val.Name,
			})
		}
	}
	return option

}
