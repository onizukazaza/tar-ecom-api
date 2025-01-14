package exception

type FailedToUpdateProduct struct{}

func (e *FailedToUpdateProduct) Error() string {
    return "Failed to update product"
}

type FailedToUpdateProductImage struct {
    ImageID string
}

func (e *FailedToUpdateProductImage) Error() string {
    return "Failed to update product image with ID: " + e.ImageID
}

type FailedToUpdateProductVariation struct {
    VariationID string
}

func (e *FailedToUpdateProductVariation) Error() string {
    return "Failed to update product variation with ID: " + e.VariationID
}
