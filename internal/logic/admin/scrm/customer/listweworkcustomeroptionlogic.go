package customer

import (
	"PowerX/internal/model/scrm/customer"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeWorkCustomerOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListWeWorkCustomerOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeWorkCustomerOptionLogic {
	return &ListWeWorkCustomerOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListWeWorkCustomerOption
//  @Description:
//  @receiver l
//  @return resp
//  @return err
//
func (l *ListWeWorkCustomerOptionLogic) ListWeWorkCustomerOption() (resp *types.WechatListCustomersOptionReply, err error) {

	option, err := l.svcCtx.PowerX.SCRM.Wechat.FindManyWeWorkCustomerOption(l.ctx)

	return &types.WechatListCustomersOptionReply{
		List: l.DTO(option),
	}, err
}

//
// DTO
//  @Description:
//  @receiver l
//  @param contacts
//  @return option
//
func (l *ListWeWorkCustomerOptionLogic) DTO(contacts []*customer.WeWorkExternalContacts) (option []*types.OptionCustomer) {
	if contacts != nil {
		for _, val := range contacts {
			option = append(option, &types.OptionCustomer{
				Id:    val.ExternalUserId,
				Names: val.Name,
			})
		}
	}
	return option
}
