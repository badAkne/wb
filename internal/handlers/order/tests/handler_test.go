package order_test

import(
	"net/http"
	"testing"
	"wb/internal/handlers/order"
	"wb/internal/service"
	serviceMock "wb/internal/service"
)

func TestOrderHandler_GetOrder(t *testing.T) {
	
	m := serviceMock.NewMockOrderService(t)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		service service.OrderService
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := order.NewOrderHandlers(tt.service)
			h.GetOrder(tt.w, tt.r)
		})
	}
}
