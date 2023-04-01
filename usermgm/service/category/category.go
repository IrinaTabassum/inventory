package category

import (
	"context"
	"fmt"

	categorypb "codemen.org/inventory/gunk/v1/category"
	"codemen.org/inventory/usermgm/storage"
)

type CoreCategory interface {
	CreateCategory(storage.Category) (*storage.Category, error)
	ListOfCategory(storage.CategoryFilter) ([]storage.Category, error)
	GetCategoryByID(int) (*storage.Category, error)
	UpdateCategory(sp storage.Category) (*storage.Category, error)
	DeleteCategory(id int) error
	GetCategoryByName(string) (*storage.Category, error)
	
}

type CategorySvc struct {
	categorypb.UnimplementedCategoryServiceServer
	core CoreCategory
}

func NewCategorySvc(cc CoreCategory) *CategorySvc {
	return &CategorySvc{
		core: cc,
	}
}

func (cs CategorySvc) CreateCategory(ctx context.Context, r *categorypb.CreateCategoryRequest) (*categorypb.CreateCategoryResponse, error) {
	category := storage.Category{
		Name: r.GetName(),
	}

	if err := category.Validate(); err != nil {
		return nil, err 
	}
	cat,_ := cs.core.GetCategoryByName(category.Name)
	if cat != nil{
		return nil, fmt.Errorf("category name must be unique")
	}

	s, err := cs.core.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return &categorypb.CreateCategoryResponse{
		Category: &categorypb.Category{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	}, nil
}
func (cs CategorySvc) ListCategory(ctx context.Context, r *categorypb.ListCategoryRequest) (*categorypb.ListCategoryResponse, error){
	Categoryfilts := storage.CategoryFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	sl, err := cs.core.ListOfCategory(Categoryfilts)
	if err != nil {
		return nil, err
	}
	
	list := make([]*categorypb.Category, len(sl))
	for i, sup := range sl {
		list[i] = &categorypb.Category{
			ID:   int32(sup.ID),
			Name: sup.Name,
		}
	}
	return &categorypb.ListCategoryResponse{
		Categorys: list,
	}, nil
}

func (cs CategorySvc) GetCategory(ctx context.Context, r *categorypb.GetCategoryRequest) (*categorypb.GetCategoryResponse, error){
	sId := int(r.GetID())
	

	s, err := cs.core.GetCategoryByID(sId)
	if err != nil {
		return nil, err
	}

	return &categorypb.GetCategoryResponse{
		Category: &categorypb.Category{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil
}

func (cs CategorySvc) UpdateCategory(ctx context.Context, r *categorypb.UpdateCategoryRequest) (*categorypb.UpdateCategoryResponse, error){
	category := storage.Category{
		ID:        int(r.GetCategory().ID),
		Name:      r.GetCategory().Name,
	}

	if err := category.Validate(); err != nil {
		fmt.Println(err)
		return nil, err 
	}
	s, err := cs.core.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return &categorypb.UpdateCategoryResponse{
		Category: &categorypb.Category{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil

}

func (cs CategorySvc) DeleteCategory(ctx context.Context, r *categorypb.DeleteCategoryRequest) (*categorypb.DeleteCategoryResponse, error){
	sId := int(r.GetID())

	err := cs.core.DeleteCategory(sId)
	if err != nil {
		return nil,err
	}

	return &categorypb.DeleteCategoryResponse{},nil
}