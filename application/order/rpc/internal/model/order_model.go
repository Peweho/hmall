package model

import (
	"context"
	"gorm.io/gorm"
	"hmall/application/order/rpc/internal/utils"
)

type OrderModel struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) *OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (m *OrderModel) UpdateOrderStatusById(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Model(&OrderPO{}).Where("id = ?", id).Updates(map[string]any{"status": utils.Paied}).Error
}

func (m *OrderModel) FindOrderById(ctx context.Context, id int64) (*OrderPO, error) {
	var order OrderPO
	err := m.db.WithContext(ctx).Where("id = ?", id).Take(&order).Error
	return &order, err
}

func (m *OrderModel) AddOrder(ctx context.Context, order *OrderPO) error {
	return m.db.WithContext(ctx).Omit("Pay_time", "Consign_time", "End_time", "Close_time", "Comment_time").Create(order).Error
}

func (m *OrderModel) AddOrderDetail(ctx context.Context, orderDetail []map[string]any) error {
	return m.db.WithContext(ctx).Model(&OrderDetailPO{}).Create(orderDetail).Error
}

func (m *OrderModel) AddOrderLogistics(ctx context.Context, orderLogistics map[string]any) error {
	return m.db.WithContext(ctx).Model(&OrderLogisticsPO{}).Create(orderLogistics).Error
}

// 查询用户最新创建的订单id
func (m *OrderModel) FindNewOrderIdByUser(ctx context.Context, usr int64) (int64, error) {
	var res OrderPO
	err := m.db.WithContext(ctx).
		Select("id").
		Where("user_id = ?", usr).
		Order("create_time desc").
		First(&res).Error
	return res.Id, err
}

// 删除订单表
func (m *OrderModel) DelOrderById(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&OrderPO{}).Error
}

// 删除订单详情表
func (m *OrderModel) DelOrderDetailById(ctx context.Context, orderId int64, itemIds []int64) error {
	return m.db.WithContext(ctx).Where("order_id = ? and item_id in ?", orderId, itemIds).Delete(&OrderDetailPO{}).Error
}

// 删除订单物流表
func (m *OrderModel) DelOrderLogisticById(ctx context.Context, orderId int64) error {
	return m.db.WithContext(ctx).Where("order_id = ?", orderId).Delete(&OrderLogisticsPO{}).Error
}
