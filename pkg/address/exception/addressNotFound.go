package exception


type AddressNotFound  struct {

}

func (e *AddressNotFound) Error() string {
    return "address not found"
}