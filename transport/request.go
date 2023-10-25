package transport

type OrderRequest struct {
	MenuID      int     `json:"menu_id"  validate:"required"`
	ToppingName string  `json:"topping_name"`
	FillingName string  `json:"filling_name"`
	TotalPrice  float64 `json:"-"`
}
