package ordersRepository

import (
	"E-Commerce/models/dto/ordersDto"
	"E-Commerce/src/orders"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type ordersRepository struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) orders.OrdersRepository {
	return ordersRepository{db}
}

func (o ordersRepository) CreateOrder(order ordersDto.OrderRepo) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	orderQuery := `INSERT INTO
	  orders (
		order_id,
		user_id,
		total_amount,
		status
	  )
	VALUES
	  ($1, $2, $3, $4)`
	_, err = tx.Exec(orderQuery, order.OrderID, order.UsersID, order.TotalAmount, order.Status)
	if err != nil {
		return err
	}

	orderItemsQuery := `INSERT INTO
	  order_items (
		order_item_id,
		order_id,
		product_id,
		quantity,
		price
	  )
	VALUES
	  ($1, $2, $3, $4, $5)`

	for _, item := range order.OrderItems {
		orderItemsID, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		_, err = tx.Exec(orderItemsQuery, orderItemsID, order.OrderID, item.ProductsID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	shippingQuery := `INSERT INTO
	  shipments (
		shipment_id,
		order_id,
		tracking_number,
		shipping_carrier,
		shipping_cost,
		shipping_address
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6)`

	shippingAddress := fmt.Sprintf("%s, %s, %s, %s, %s", order.ShippingAddress.Address, order.ShippingAddress.City, order.ShippingAddress.State, order.ShippingAddress.Country, order.ShippingAddress.PostalCode)
	_, err = tx.Exec(shippingQuery, order.ShippingID, order.OrderID, order.TrackingNumber, order.ShippingCarrier, order.ShippingCost, shippingAddress)
	if err != nil {
		return err
	}

	return nil
}
