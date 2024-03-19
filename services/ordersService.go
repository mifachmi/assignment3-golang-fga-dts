package services

import (
	"asssignment2/db"
	"asssignment2/models"
	"asssignment2/requests"
)

type OrderService struct{}

func (os *OrderService) CreateOrder(orderWithItems requests.OrderWithItemsBody) (*models.Order, error) {
	tx := db.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order := models.Order{
		CustomerName: orderWithItems.CustomerName,
		OrderedAt:    orderWithItems.OrderedAt,
	}

	if err := db.GetDB().Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var items []models.Item
	for _, itemData := range orderWithItems.Items {
		items = append(items, models.Item{
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
			OrderID:     order.OrderID,
		})
	}

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &order, nil
}

func (os *OrderService) UpdateOrder(orderID uint, orderWithItems requests.OrderWithItemsBody) (*models.Order, error) {
	tx := db.GetDB().Begin()

	var existOrder models.Order

	if err := tx.Table("orders").Where("order_id = ?", orderID).First(&existOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	existOrder.CustomerName = orderWithItems.CustomerName
	existOrder.OrderedAt = orderWithItems.OrderedAt

	if err := tx.Save(&existOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, itemData := range orderWithItems.Items {
		var existingItem models.Item

		if err := tx.Table("items").Where("item_id = ?", itemData.ItemID).First(&existingItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		existingItem.ItemCode = itemData.ItemCode
		existingItem.Description = itemData.Description
		existingItem.Quantity = itemData.Quantity
		existingItem.OrderID = int64(orderID)

		if err := tx.Save(&existingItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	db.GetDB().Table("orders").Preload("Items").Where("order_id = ?", orderID).First(&existOrder)

	return &existOrder, nil
}

func (os *OrderService) DeleteOrder(orderID uint) error {
	var order models.Order
	if err := db.GetDB().First(&order, orderID).Error; err != nil {
		return err
	}

	if err := db.GetDB().Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		return err
	}

	if err := db.GetDB().Delete(&order).Error; err != nil {
		return err
	}

	return nil
}
