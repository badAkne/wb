package help

import (
	"testing"

	"wb/internal/model"
	testdata "wb/internal/testdata"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	testCases := []struct {
		name   string
		in     model.Order
		expErr error
	}{
		{
			name:   "bad_UID",
			in:     testdata.BadUID,
			expErr: errInvalidUID,
		},
		{
			name:   "no_items",
			in:     testdata.NoItems,
			expErr: errInvalidItemsAmount,
		},
		{
			name:   "no_payment",
			in:     testdata.NoPayment,
			expErr: errNoPaymentDetails,
		},
		{
			name:   "no_delivery",
			in:     testdata.NoDelivery,
			expErr: errNoDeliveryDetails,
		},
		{
			name:   "bad_track_number",
			in:     testdata.BadTrackNumber,
			expErr: errBadTrackNumber,
		},
		{
			name:   "invalid_entry",
			in:     testdata.InvalidEntry,
			expErr: errInvalidEntry,
		},
		{
			name:   "invalid_delivery_name",
			in:     testdata.InvalidDeliveryName,
			expErr: errInvalidDeliveryName,
		},
		{
			name:   "invalid_phone_number",
			in:     testdata.InvalidPhoneNumber,
			expErr: errinvalidPhoneNumber,
		},
		{
			name:   "invalid_zip",
			in:     testdata.InvalidZip,
			expErr: errInvalidZip,
		},
		{
			name:   "invalid_city",
			in:     testdata.InvalidCity,
			expErr: errInvalidCity,
		},
		{
			name:   "invalid_adress",
			in:     testdata.InvalidAddress,
			expErr: errInvalidAddress,
		},
		{
			name:   "invalid_region",
			in:     testdata.InvalidRegion,
			expErr: errInvalidRegion,
		},
		{
			name:   "invalid_email",
			in:     testdata.InvalidEmail,
			expErr: errInvalidEmail,
		},
		{
			name:   "invalid_transaction",
			in:     testdata.InvalidTransaction,
			expErr: errInvalidTransaction,
		},
		{
			name:   "invalid_currency",
			in:     testdata.InvalidCurrency,
			expErr: errInvalidCurrency,
		},
		{
			name:   "invalid_provider",
			in:     testdata.InvalidProvider,
			expErr: errInvalidProvider,
		},
		{
			name:   "invalid_amount",
			in:     testdata.InvalidAmount,
			expErr: errInvalidAmount,
		},
		{
			name:   "invalid_payment_dt",
			in:     testdata.InvalidPaymentDT,
			expErr: errInvalidPaymentDT,
		},
		{
			name:   "invalid_bank",
			in:     testdata.InvalidBank,
			expErr: errInvalidBank,
		},
		{
			name:   "invalid_delivery_cost",
			in:     testdata.InvalidDelvieryCost,
			expErr: errInvalidDeliveryCost,
		},
		{
			name:   "invalid_goods_total",
			in:     testdata.InvalidGoodsTotal,
			expErr: errInvalidGoodsTotal,
		},
		{
			name:   "invalid_chrt_id",
			in:     testdata.InvalidChrtID,
			expErr: errInvalidChartID,
		},
		{
			name:   "invalid_price",
			in:     testdata.InvalidPrice,
			expErr: errInvalidPrice,
		},
		{
			name:   "invalid_RID",
			in:     testdata.InvalidRID,
			expErr: errInvalidRID,
		},
		{
			name:   "invalid_item_name",
			in:     testdata.InvalidItemName,
			expErr: errInvalidItemName,
		},
		{
			name:   "invalid_sale_amont",
			in:     testdata.InvalidSaleAmount,
			expErr: errInvalidSaleAmount,
		},
		{
			name:   "invalid_total_price",
			in:     testdata.InvalidTotalPrice,
			expErr: errInvalidTotalPrice,
		},
		{
			name:   "invalid_nm_id",
			in:     testdata.InvalidNmID,
			expErr: errInvalidNmID,
		},
		{
			name:   "invalid_brand_name",
			in:     testdata.InvalidBrandName,
			expErr: errInvalidBrandName,
		},
		{
			name:   "invalid_status",
			in:     testdata.InvalidStatus,
			expErr: errInvalidStatus,
		},
		{
			name:   "invalid_locale",
			in:     testdata.InvalidLocale,
			expErr: errInvalidLocale,
		},
		{
			name:   "invalid_customer_id",
			in:     testdata.InvalidCustomerID,
			expErr: errInvalidCustomerID,
		},
		{
			name:   "invalid_delivery_service",
			in:     testdata.InvalidDeliveryService,
			expErr: errInvalidDeliveryService,
		},
		{
			name:   "invalid_shard_key",
			in:     testdata.InvalidShardKey,
			expErr: errInvalidShardKey,
		},
		{
			name:   "invalid_smid",
			in:     testdata.InvalidSmID,
			expErr: errInvalidSmID,
		},
		{
			name:   "invalid_creation_date",
			in:     testdata.InvalidCreationDate,
			expErr: errInvalidCreateionDate,
		},
		{
			name:   "invalid_oof_key",
			in:     testdata.InvalidOofKey,
			expErr: errInvalidOofShard,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			//err := ValidateOrder(tCase.in)
			err := ValidateOrder(tCase.in)
			require.Error(t, err)
			require.EqualError(t, err, tCase.expErr.Error())
		})
	}
}
