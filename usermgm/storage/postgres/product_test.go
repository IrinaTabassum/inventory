package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateProduct(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newCategory := storage.Category{
		Name: "Phone",
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.Product
		want    *storage.Product
		wantErr bool
	}{
		{
			name: "CREATE_PRODUCT_SUCEESS",
			in: storage.Product{
				CategoryId:  category.ID,
				Name:        "product1",
				Description: "this is test product",
				Price:       200,
			},
			want: &storage.Product{
				CategoryId:  category.ID,
				Name:        "product1",
				Description: "this is test product",
				Price:       200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateProduct(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Product{}, "ID", "CategoryName", "Quantity", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateProduct() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOfProduct(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCategory := storage.Category{
		Name: "Phone",
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	products := []storage.Product{
		{
			CategoryId:  category.ID,
			Name:        "Product1",
			Description: "this is product1",
			Price:       100,
		},
		{
			CategoryId:  category.ID,
			Name:        "Product2",
			Description: "this is product2",
			Price:       100,
		},
		{
			CategoryId:  category.ID,
			Name:        "Product3",
			Description: "this is product3",
			Price:       100,
		},
	}

	for _, product := range products {
		_, err := s.CreateProduct(product)
		if err != nil {
			t.Fatalf("unable to create product for list product testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.ProductFilter
		want    []storage.Product
		wantErr bool
	}{
		{
			name: "LIST_ALL_PRODUCT_SUCCESS",
			in:   storage.ProductFilter{},
			want: []storage.Product{
				{
					CategoryId:  category.ID,
					Name:        "Product1",
					Description: "this is product1",
					Price:       100,
				},
				{
					CategoryId:  category.ID,
					Name:        "Product2",
					Description: "this is product2",
					Price:       100,
				},
				{
					CategoryId:  category.ID,
					Name:        "Product3",
					Description: "this is product3",
					Price:       100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfProduct(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Product{}, "ID", "Quantity", "CategoryName", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListProduct() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCategory := storage.Category{
		Name: "Phone",
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	newCategory2 := storage.Category{
		Name: "Laptop",
	}
	category2, err := s.CreateCategory(newCategory2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	newProduct := storage.Product{
		CategoryId:  category.ID,
		Name:        "product1",
		Description: "this is test product",
		Price:       200,
	}

	tests := []struct {
		name    string
		in      storage.Product
		want    *storage.Product
		wantErr bool
	}{
		{
			name: "UPDATE_Product_SUCEESS",
			in: storage.Product{
				CategoryId:  category2.ID,
				Name:        "product1Update",
				Description: "this is test product Update",
				Price:       300,
			},
			want: &storage.Product{
				CategoryId:  category2.ID,
				Name:        "product1Update",
				Description: "this is test product Update",
				Price:       300,
			},
		},
	}
	product, err := s.CreateProduct(newProduct)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = product.ID
			got, err := s.UpdateProduct(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Product{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateProduct() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteProductByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCategory := storage.Category{
		Name: "Phone",
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	newProduct := storage.Product{
		CategoryId:  category.ID,
		Name:        "product1",
		Description: "this is test product",
		Price:       200,
	}
	product, err := s.CreateProduct(newProduct)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_PRODUCT_BY_ID_SUCEESS",
			in:   product.ID,
		},
		{
			name: "DELETE_PRODUCT_BY_ID_FAILED",
			in:      product.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteProductByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCategory := storage.Category{
		Name: "Phone",
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	newProduct := storage.Product{
		CategoryId:  category.ID,
		Name:        "product1",
		Description: "this is test product",
		Price:       200,
	}
	product, err := s.CreateProduct(newProduct)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		want    *storage.Product
		wantErr bool
	}{
		{
			name: "DELETE_PRODUCT_BY_ID_SUCEESS",
			in:   product.ID,
			want: &storage.Product{
				CategoryId:  category.ID,
				Name:        "product1",
				Description: "this is test product",
				Price:       200,
			},
			
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetProductByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Product{}, "ID", "CategoryName", "Quantity", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetProductByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
