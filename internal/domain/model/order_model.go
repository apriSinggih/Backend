package model

type OrderTicketRequest struct {
	EventID    uint   `json:"event_id" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required,min=1"`
	SeatNumber string `json:"seat_number" validate:"required"`
}

type OrderResponse struct {
	ID         uint             `json:"id"`
	EventID    uint             `json:"event_id"`
	UserID     string           `json:"user_id"`
	Quantity   int              `json:"quantity"`
	TotalPrice float64          `json:"total_price"`
	Tickets    []TicketResponse `json:"tickets"`
}

type UpdateOrderRequest struct {
	ID uint `param:"id" validate:"required"`
}

type GetOrderRequest struct {
	ID uint `param:"id" validate:"required"`
}

type OrdersRequest struct {
	Page  int    `query:"page" validate:"numeric"`
	Size  int    `query:"size" validate:"numeric"`
	Sort  string `query:"sort" validate:"omitempty,oneof=date total_price"`
	Order string `query:"order" validate:"omitempty"`
}
