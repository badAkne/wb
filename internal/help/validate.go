package help

import (
	"fmt"
	"wb/internal/model"
)

var (
	errInvalidUID = fmt.Errorf("invalid_uid")

	errInvalidItemsAmount = fmt.Errorf("invalid_items_amount")

	errNoPaymentDetails = fmt.Errorf("no_paymnent_details")

	errNoDeliveryDetails = fmt.Errorf("no_delivery_details")

	errBadTrackNumber = fmt.Errorf("invalid_track_number")

	errInvalidEntry = fmt.Errorf("invalid_entry")

	errInvalidDeliveryName = fmt.Errorf("invalid_delivery_name")

	errinvalidPhoneNumber = fmt.Errorf("invalid_phone")

	errInvalidZip = fmt.Errorf("invalid_zip")

	errInvalidCity = fmt.Errorf("invalid_city")

	errInvalidAddress = fmt.Errorf("invalid_address")

	errInvalidRegion = fmt.Errorf("invalid_region")

	errInvalidEmail = fmt.Errorf("invalid_email")

	errInvalidTransaction = fmt.Errorf("invalid_transaction")

	errInvalidCurrency = fmt.Errorf("invalid_currency")

	errInvalidProvider = fmt.Errorf("invalid_provider")

	errInvalidAmount = fmt.Errorf("invalid_amount")

	errInvalidPaymentDT = fmt.Errorf("invalid_payment_dt")

	errInvalidBank = fmt.Errorf("invalid_bank")

	errInvalidDeliveryCost = fmt.Errorf("invalid_delivery_cost")

	errInvalidGoodsTotal = fmt.Errorf("invalid_goods_total")

	errInvalidChartID = fmt.Errorf("invalid_chrt_id")

	errInvalidPrice = fmt.Errorf("invalid_price")

	errInvalidRID = fmt.Errorf("invalid_rid")

	errInvalidItemName = fmt.Errorf("invalid_item_name")

	errInvalidSaleAmount = fmt.Errorf("invalid sale amount")

	errInvalidTotalPrice = fmt.Errorf("invalid_total_price")

	errInvalidNmID = fmt.Errorf("invalid_nmid")

	errInvalidBrandName = fmt.Errorf("invalid_brand_name")

	errInvalidStatus = fmt.Errorf("invalid_status")

	errInvalidLocale = fmt.Errorf("invalid_locale")

	errInvalidCustomerID = fmt.Errorf("invalid_customer_id")

	errInvalidDeliveryService = fmt.Errorf("invalid_delivery_service")

	errInvalidShardKey = fmt.Errorf("invalid_shard_key")

	errInvalidSmID = fmt.Errorf("invalid_smid")

	errInvalidCreateionDate = fmt.Errorf("invalid_creation_date")

	errInvalidOofShard = fmt.Errorf("invalid_oof_shard")
)

func ValidateOrder(order model.Order) error {

	if order.OrderUID == "" {
		return errInvalidUID
	}

	if order.TrackNumber == "" {
		return errBadTrackNumber
	}

	if order.Entry == "" {
		return errInvalidEntry
	}

	if order.Delivery == (model.Delivery{}) {
		return errNoDeliveryDetails
	}

	if order.Delivery.Name == "" {
		return errInvalidDeliveryName
	}

	if order.Delivery.Phone == "" {
		return errinvalidPhoneNumber
	}

	if order.Delivery.Zip == "" {
		return errInvalidZip
	}

	if order.Delivery.City == "" {
		return errInvalidCity
	}

	if order.Delivery.Address == "" {
		return errInvalidAddress
	}

	if order.Delivery.Region == "" {
		return errInvalidRegion
	}

	if order.Delivery.Email == "" {
		return errInvalidEmail
	}

	if order.Payment == (model.Payment{}) {
		return errNoPaymentDetails
	}

	if order.Payment.Transaction == "" {
		return errInvalidTransaction
	}

	if order.Payment.Currency == "" {
		return errInvalidCurrency
	}

	if order.Payment.Provider == "" {
		return errInvalidProvider
	}

	if order.Payment.Amount < 1 {
		return errInvalidAmount
	}

	if order.Payment.Paymentdt == 0 {
		return errInvalidPaymentDT
	}

	if order.Payment.Bank == "" {
		return errInvalidBank
	}

	if order.Payment.DeliveryCost < 0 {
		return errInvalidDeliveryCost
	}

	if order.Payment.GoodsTotal < 1 {
		return errInvalidGoodsTotal
	}

	if len(order.Items) == 0 {
		return errInvalidItemsAmount
	}

	for i := 0; i < len(order.Items); i++ {
		if order.Items[i].ChrtId == 0 {
			return errInvalidChartID
		}

		if order.Items[i].TrackNumber == "" {
			return errBadTrackNumber
		}

		if order.Items[i].Price < 1 {
			return errInvalidPrice
		}

		if order.Items[i].Rid == "" {
			return errInvalidRID
		}

		if order.Items[i].Name == "" {
			return errInvalidItemName
		}

		if order.Items[i].Sale < 0 {
			return errInvalidSaleAmount
		}

		if order.Items[i].TotalPrice < 1 {
			return errInvalidTotalPrice
		}

		if order.Items[i].NmId < 1 {
			return errInvalidNmID
		}

		if order.Items[i].Brand == "" {
			return errInvalidBrandName
		}

		if order.Items[i].Status < 0 {
			return errInvalidStatus
		}
	}

	if order.Locale == "" {
		return errInvalidLocale
	}

	if order.CustomerId == "" {
		return errInvalidCustomerID
	}

	if order.DeliveryService == "" {
		return errInvalidDeliveryService
	}

	if order.ShardKey == "" {
		return errInvalidShardKey
	}

	if order.SmId < 1 {
		return errInvalidSmID
	}

	if order.DateCreated == "" {
		return errInvalidCreateionDate
	}

	if order.OofShard == "" {
		return errInvalidOofShard
	}

	return nil
}
