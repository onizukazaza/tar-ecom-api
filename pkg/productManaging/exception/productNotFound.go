package exception

type ProductNotFoundError struct{}

func (e *ProductNotFoundError) Error() string {
	return "Product was not found"
}
