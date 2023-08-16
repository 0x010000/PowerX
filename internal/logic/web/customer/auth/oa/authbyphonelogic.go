package oa

import (
	"PowerX/internal/model"
	"PowerX/internal/model/customerdomain"
	"PowerX/internal/types/errorx"
	customerdomain2 "PowerX/internal/uc/powerx/customerdomain"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByPhoneLogic {
	return &AuthByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthByPhoneLogic) AuthByPhone(req *types.OACustomerAuthRequest) (resp *types.OACustomerLoginAuthReply, err error) {

	user, err := l.svcCtx.PowerX.WechatOA.App.OAuth.UserFromCode(req.Code)
	if err != nil {
		return nil, err
	}
	if user.GetOpenID() == `` {
		return nil, errorx.ErrBadRequest
	}
	// 解码手机授权信息
	msgData, errEncrypt := l.svcCtx.PowerX.WechatOA.App.Encryptor.DecryptContent(req.EncryptedData)
	fmt.Println(msgData)
	if errEncrypt != nil {
		return nil, errors.New(errEncrypt.ErrMsg)
	}

	oaCustomer := &model.WechatOACustomer{
		SessionKey: user.GetAccessToken(),
		OpenId:     user.GetOpenID(),
		UnionId:    user.Get(`unionid`, ``).(string),
	}

	// upsert 公众号客户记录
	oaCustomer, err = l.svcCtx.PowerX.WechatOA.UpsertOACustomer(l.ctx, oaCustomer)
	if err != nil {
		return nil, err
	}

	source := l.svcCtx.PowerX.DataDictionary.GetCachedDDId(l.ctx, model.TypeSourceChannel, model.ChannelWechat)
	personType := l.svcCtx.PowerX.DataDictionary.GetCachedDDId(l.ctx, customerdomain.TypeCustomerType, customerdomain.CustomerPersonal)

	// upsert 线索
	lead := &customerdomain.Lead{
		Name:        user.GetName(),
		Mobile:      user.GetMobile(),
		Source:      source,
		Type:        personType,
		IsActivated: true,
		ExternalId: customerdomain.ExternalId{
			OpenIdInWeChatOfficialAccount: oaCustomer.OpenId,
		},
	}
	lead, err = l.svcCtx.PowerX.Lead.UpsertLead(l.ctx, lead)
	if err != nil {
		return nil, err
	}

	// upsert 客户
	customer := &customerdomain.Customer{
		Name:        user.GetName(),
		Mobile:      user.GetMobile(),
		Source:      source,
		Type:        personType,
		IsActivated: true,
		ExternalId: customerdomain.ExternalId{
			OpenIdInWeChatOfficialAccount: oaCustomer.OpenId,
		},
	}
	customer, err = l.svcCtx.PowerX.Customer.UpsertCustomer(l.ctx, customer)
	if err != nil {
		return nil, err
	}

	token := l.svcCtx.PowerX.CustomerAuthorization.SignWebToken(customer, l.svcCtx.Config.JWT.WebJWTSecret)

	return &types.OACustomerLoginAuthReply{
		OpenId:      oaCustomer.OpenId,
		UnionId:     oaCustomer.UnionId,
		PhoneNumber: user.GetMobile(),
		NickName:    user.GetNickname(),
		AvatarURL:   user.GetAvatar(),
		Gender:      ``,
		Token: types.OAToken{
			TokenType:    token.TokenType,
			ExpiresIn:    fmt.Sprintf("%d", customerdomain2.CustomerTokenExpiredDuration),
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		},
	}, nil
}
