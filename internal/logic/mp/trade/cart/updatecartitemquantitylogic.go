package cart

import (
	"PowerX/internal/types/errorx"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartItemQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemQuantityLogic {
	return &UpdateCartItemQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartItemQuantityLogic) UpdateCartItemQuantity(req *types.UpdateCartItemQuantityRequest) (resp *types.UpdateCartItemQuantityReply, err error) {

	if req.Quantity <= 0 {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "修改数量必须大于0")
	}

	mdlCartItem, err := l.svcCtx.PowerX.Cart.GetCartItem(l.ctx, req.ItemId)
	if err != nil {
		return nil, err
	}

	// 更新购物项对象
	mdlCartItem, err = l.svcCtx.PowerX.Cart.UpsertCartItem(l.ctx, mdlCartItem, []string{"quantity"})

	return &types.UpdateCartItemQuantityReply{
		CartItem: TransformCartItemToCartItemReplyToMP(mdlCartItem),
	}, err
}