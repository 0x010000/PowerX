package trade

import (
	"PowerX/internal/model/powermodel"
	"time"
)

type PivotOrderToInventoryLog struct {
	*powermodel.PowerModel

	OrderId             int64     `gorm:"comment:订单Id" json:"orderId"`
	OrderItemId         int64     `gorm:"comment:订单项Id" json:"orderItemId"`
	ProductId           int64     `gorm:"comment:商品Id" json:"productId"`
	InventoryId         int64     `gorm:"comment:库存Id" json:"inventoryId"`
	Action              string    `gorm:"comment:操作类型" json:"action"`
	ActionTime          time.Time `gorm:"comment:操作时间" json:"actionTime"`
	StockQuantityBefore int       `gorm:"comment:回滚前的库存数量" json:"stockQuantityBefore"`
	StockQuantityAfter  int       `gorm:"comment:回滚后的库存数量" json:"stockQuantityAfter"`
}

type ActionType int

const (
	ActionCreate   ActionType = 1 // 创建
	ActionUpdate   ActionType = 2 // 更新
	ActionDelete   ActionType = 3 // 删除
	ActionRollback ActionType = 4 // 回滚
	ActionOther    ActionType = 5 // 其他
)
