package exception

type ProductNotFound struct{}

func (e *ProductNotFound) Error() string {
	return "Product was not found"
}
