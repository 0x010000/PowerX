package active

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"testing"
)

func TestActionSceneActivitiesDetailLogic_activitiesSceneQrcode(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		sceneQrcode []*scene.SceneActivitiesQrcode
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantQrcode []*types.ActivitiesQrcodeWeb
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			active := &ActionSceneActivitiesDetailLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			if gotQrcode := active.activitiesSceneQrcode(tt.args.sceneQrcode); !reflect.DeepEqual(gotQrcode, tt.wantQrcode) {
				t.Errorf("activitiesSceneQrcode() = %v, want %v", gotQrcode, tt.wantQrcode)
			}
		})
	}
}
