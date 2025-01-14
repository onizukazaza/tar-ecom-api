package service

import ( 
    _productModel "github.com/onizukazaza/tar-ecom-api/pkg/product/model"
    _productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)


type ProductService interface {
    CreateProduct(req *_productModel.ProductCreatingReq) error
    EditProduct(req *_productModel.ProductEditingReq) error
    DeleteProduct(productID string) error 
    DeleteProductWithSeller(productID string, sellerID string) error
    Listing(filter *_productManagingModel.FilterRequestBySeller, sellerID string) ([]*_productManagingModel.ProductDetail, error)
    GetProductByIDAndSeller(productID string, sellerID string) (*_productManagingModel.ProductDetail, error)

    
    
    
}
