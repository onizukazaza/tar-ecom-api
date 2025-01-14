package exception

type  UnCreateProduct struct {}

func (e *UnCreateProduct) Error() string {
    return "Failed to create product"
}