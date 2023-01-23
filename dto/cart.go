package dto

import "final-project-backend/entity"

type GetCartResponse struct {
	CartItems  []*entity.Cart `json:"cartItems"`
	TotalPrice float64        `json:"totalPrice"`
}

func (r *GetCartResponse) FromCart(cart []*entity.Cart) {
	var totalPrice float64
	for _, item := range cart {
		totalPrice += item.Course.Price
	}

	r.CartItems = cart
	r.TotalPrice = totalPrice
}
