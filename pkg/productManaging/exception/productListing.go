package exception

type ProductFetchingError struct{}

func (e *ProductFetchingError) Error() string {
	return "product fetching failed"
}

