package transport

import "time"

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Data: data,
		Meta: meta,
	}

	return jsonResponse
}

type (
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}

	MenuMapping struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	AdditionMapping struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	OrderMapping struct {
		ID          int       `json:"id"`
		MenuName    string    `json:"menu_name"`
		ToppingName string    `json:"topping_name"`
		FillingName string    `json:"filling_name"`
		TotalPrice  float64   `json:"total_price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Response struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}
)
