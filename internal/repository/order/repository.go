package order

import (
	"context"
	"fmt"
	"sync"
	"time"
	"wb/internal/model"
	"wb/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/patrickmn/go-cache"
)

var _ repository.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	cache *cache.Cache
	db    *pgxpool.Pool
	mu    sync.RWMutex
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return &OrderRepository{
		cache: c,
		db:    db,
	}
}

func (r *OrderRepository) LoadCacheFromDB(ctx context.Context) error {
	op := "internal.repository.order.LoadCachdeFromDB"
	rows, err := r.db.Query(ctx, `
		SELECT o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
		       d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
		       p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee
		FROM orders o
		JOIN deliveries d ON o.order_uid = d.order_uid
		JOIN payments p ON o.order_uid = p.order_uid
	`)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var o model.Order
		var dateCreated time.Time
		err := rows.Scan(
			&o.OrderUID, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.Sharkey, &o.SmId, &dateCreated, &o.OofShard,
			&o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email,
			&o.Payment.Transaction, &o.Payment.RequestId, &o.Payment.Currency, &o.Payment.Provider, &o.Payment.Amount, &o.Payment.Paymentdt, &o.Payment.Bank, &o.Payment.DeliveryCost, &o.Payment.GoodsTotal, &o.Payment.CustomFee,
		)
		if err != nil {
			return fmt.Errorf("%s:%v", op, err)
		}

		itemRows, err := r.db.Query(ctx, "SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1", o.OrderUID)
		if err != nil {
			return fmt.Errorf("%s:%v", op, err)
		}

		for itemRows.Next() {
			var item model.Item
			err = itemRows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
			if err != nil {
				return fmt.Errorf("%s:%v", op, err)
			}
			o.Items = append(o.Items, item)
		}
		itemRows.Close()

		r.cache.Set(o.OrderUID, o, cache.DefaultExpiration)
	}

	return nil
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order model.Order) error {
	op := "internal.repository.order.SaveOrder"

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	_, err = tx.Exec(ctx, `
		INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (order_uid) DO NOTHING`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Sharkey, order.SmId, order.DateCreated, order.OofShard,
	)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email,
	)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO payments (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.Payment.Transaction, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.Paymentdt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee,
	)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	for _, item := range order.Items {
		_, err = tx.Exec(ctx, `
			INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			order.OrderUID, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status,
		)
		if err != nil {
			return fmt.Errorf("%s:%v", op, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	r.cache.Set(order.OrderUID, order, cache.DefaultExpiration)

	return nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, id string) (model.Order, bool, error) {
	op := "internal.repository.order.GetOrderbyID"

	order, ok := r.cache.Get(id)
	if ok {
		return order.(model.Order), true, nil
	}

	row := r.db.QueryRow(ctx, `
		SELECT o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
		       d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
		       p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee
		FROM orders o
		JOIN deliveries d ON o.order_uid = d.order_uid
		JOIN payments p ON o.order_uid = p.order_uid
		WHERE o.order_uid = $1
	`, id)

	var dateCreated time.Time
	var o model.Order
	err := row.Scan(
		&o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.Sharkey, &o.SmId, &dateCreated, &o.OofShard,
		&o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email,
		&o.Payment.Transaction, &o.Payment.RequestId, &o.Payment.Currency, &o.Payment.Provider, &o.Payment.Amount, &o.Payment.Paymentdt, &o.Payment.Bank, &o.Payment.DeliveryCost, &o.Payment.GoodsTotal, &o.Payment.CustomFee,
	)
	if err != nil {
		return model.Order{}, false, fmt.Errorf("%s:%v", op, err)
	}
	o.OrderUID = id

	itemRows, err := r.db.Query(ctx, "SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1", id)
	if err != nil {
		return model.Order{}, false, fmt.Errorf("%s:%v", op, err)
	}

	for itemRows.Next() {
		var item model.Item
		err = itemRows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			return model.Order{}, false, fmt.Errorf("%s:%v", op, err)
		}
		o.Items = append(o.Items, item)
	}
	itemRows.Close()

	r.cache.Set(id, o, cache.DefaultExpiration)

	return o, true, nil
}
