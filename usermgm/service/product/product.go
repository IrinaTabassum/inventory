package product

import (
	"context"
	"fmt"

	productpb "codemen.org/inventory/gunk/v1/product"
	"codemen.org/inventory/usermgm/storage"
)

type CoreProduct interface {
	CreateProduct(storage.Product) (*storage.Product, error)
	ListOfProduct(storage.ProductFilter) ([]storage.Product, error)
	GetProductByID(int) (*storage.Product, error)
	UpdateProduct(storage.Product) (*storage.Product, error)
	DeleteProduct(id int) error	
}

type ProductSvc struct {
	productpb.UnimplementedProductServiceServer
	core CoreProduct
}

func NewProductSvc(cp CoreProduct) *ProductSvc {
	return &ProductSvc{
		core: cp,
	}
}

func (ps ProductSvc) CreateProduct(ctx context.Context, r *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	product := storage.Product{
		CategoryId:  int(r.GetCategoryId()),
		Name:        r.GetName(),
		Description: r.GetDescription(),
		Price:       r.GetPrice(),
		
	}

	if err := product.Validate(); err != nil {
		fmt.Println(err)
		return nil, err 
	}

	p, err := ps.core.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			ID:          int32(p.ID),
			CategoryId:  int32(p.CategoryId),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
	}, nil
}
func (ps ProductSvc) ListProduct(ctx context.Context, r *productpb.ListProductRequest) (*productpb.ListProductResponse, error){
	productfilts := storage.ProductFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	pl, err := ps.core.ListOfProduct(productfilts)
	if err != nil {
		return nil, err
	}
	
	list := make([]*productpb.Product, len(pl))
	for i, p := range pl {
		list[i] = &productpb.Product{
			ID:          int32(p.ID),
			CategoryId:  int32(p.CategoryId),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}
	}
	return &productpb.ListProductResponse{
		Products: list,
	}, nil
}

func (ps ProductSvc) GetProduct(ctx context.Context, r *productpb.GetProductRequest) (*productpb.GetProductResponse, error){
	pId := int(r.GetID())
	

	p, err := ps.core.GetProductByID(pId)
	if err != nil {
		return nil, err
	}

	return &productpb.GetProductResponse{
		Product: &productpb.Product{
			ID:          int32(p.ID),
			CategoryId:  int32(p.CategoryId),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
	},nil
}

func (ps ProductSvc) UpdateProduct(ctx context.Context, r *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error){
	product := storage.Product{
		ID:          int(r.GetProduct().ID),
		CategoryId:  int(r.GetProduct().CategoryId),
		Name:        r.GetProduct().Name,
		Description: r.GetProduct().Description,
		Price:       r.GetProduct().Price,
	}
	p, err := ps.core.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return &productpb.UpdateProductResponse{
		Product: &productpb.Product{
			ID:          int32(p.ID),
			CategoryId:  int32(p.CategoryId),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
	},nil

}

func (ps ProductSvc) DeleteProduct(ctx context.Context, r *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error){
	pId := int(r.GetID())

	err := ps.core.DeleteProduct(pId)
	if err != nil {
		return nil,err
	}

	return &productpb.DeleteProductResponse{},nil
}
