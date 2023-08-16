package scene

import (
	"PowerX/internal/model/scene"
	"PowerX/internal/types"
)

// FindOneSceneQrcodeDetail
//
//	@Description:
//	@receiver sc
//	@param qid
//	@return *qrcode.QrcodeActive
func (sc *sceneUseCase) FindOneSceneQrcodeDetail(qid string) *scene.SceneQrcode {

	qrcode := sc.modelSceneQrcode.qrcode.FindEnableSceneQrcodeByQid(sc.db, qid)

	return qrcode

}

// IncreaseSceneCpaNumber
//
//	@Description:
//	@receiver sc
//	@param qid
func (sc *sceneUseCase) IncreaseSceneCpaNumber(qid string) {

	sc.modelSceneQrcode.qrcode.IncreaseCpa(sc.db, qid)
}

//
// FindSceneQrcodeOption
//  @Description:
//  @receiver sc
//  @param opt
//  @return []*scene.SceneQrcode
//
func (sc *sceneUseCase) FindSceneQrcodeOption(opt *types.SceneQrcodeRequest) []*scene.SceneQrcode {

	qrcode := sc.modelSceneQrcode.qrcode.Options(sc.db, opt)

	return qrcode

}
