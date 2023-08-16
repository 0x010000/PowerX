package active

import (
	"PowerX/internal/model/scene"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSceneQrcodeOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSceneQrcodeOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSceneQrcodeOptionLogic {
	return &ListSceneQrcodeOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListSceneQrcodeOption
//  @Description:
//  @receiver scene
//  @param opt
//  @return resp
//  @return err
//
func (scene *ListSceneQrcodeOptionLogic) ListSceneQrcodeOption(opt *types.SceneQrcodeRequest) (resp *types.SceneQrcodeOptionReply, err error) {

	option := scene.svcCtx.PowerX.Scene.Svc.FindSceneQrcodeOption(opt)
	return &types.SceneQrcodeOptionReply{
		List: scene.DTO(option),
	}, err

}

//
// DTO
//  @Description:
//  @receiver scene
//  @param data
//  @return qrcode
//
func (scene *ListSceneQrcodeOptionLogic) DTO(data []*scene.SceneQrcode) (qrcode []*types.SceneQrcode) {
	if data != nil {
		for _, val := range data {
			qrcode = append(qrcode, &types.SceneQrcode{
				Qid:              val.QId,
				Name:             val.Name,
				RealQrcodeLink:   val.RealQrcodeLink,
				Platform:         val.Platform,
				ActiveQrcodeLink: val.ActiveQrcodeLink,
				State:            val.State,
			})
		}
	}
	return qrcode
}
