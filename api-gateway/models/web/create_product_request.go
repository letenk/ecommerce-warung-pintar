package web

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required,numeric"`
	Quantity    int64  `json:"quantity" binding:"required,numeric"`
}
