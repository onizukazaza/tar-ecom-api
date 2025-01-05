package exception

type ProductNotPermission struct {}

func (e *ProductNotPermission) Error() string {
    return "You are not allowed to perform this action on this product"  //send to buyer error
}