package service

import _productModel "github.com/onizukazaza/tar-ecom-api/pkg/product/model"


type ProductService interface {
    CreateProduct(req *_productModel.ProductCreatingReq) error
    EditProduct(req *_productModel.ProductEditingReq) error
    DeleteProduct(productID string) error 
    DeleteProductWithSeller(productID string, sellerID string) error
    
}
