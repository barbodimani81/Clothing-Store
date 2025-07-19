package requests

type AddToCartInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartQuantityInput struct {
	ItemID   uint `json:"item_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required,min=1"`
}
