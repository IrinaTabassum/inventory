package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateOrUpdateStock(t *testing.T) {
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
		in      storage.Stock
		want    *storage.Stock
		wantErr bool
	}{
		{
			name: "CREATE_STOCK_SUCEESS",
			in: storage.Stock{
				ProductId: product.ID,
				Quantity:  3,
			},
			want: &storage.Stock{
				ProductId: product.ID,
				Quantity:  3,
			},
		},
		{
			name: "CREATE_STOCK_USING_SAMEPRODUCT_ID_SUCEESS",
			in: storage.Stock{
				ProductId: product.ID,
				Quantity:  7,
			},
			want: &storage.Stock{
				ProductId: product.ID,
				Quantity:  7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateOrUpdateStock(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateOrUpdateStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Stock{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateOrUpdateStock() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOStock(t *testing.T) {
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
	newProduct2 := storage.Product{
		CategoryId:  category.ID,
		Name:        "product2",
		Description: "this is test product2",
		Price:       100,
	}
	product2, err := s.CreateProduct(newProduct2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}

	stocks := []storage.Stock{
		{
			ProductId: product.ID,
			Quantity:  3,
		},
		{
			ProductId: product.ID,
			Quantity:  6,
		},
		{
			ProductId: product2.ID,
			Quantity:  7,
		},
	}

	for _, stock := range stocks {
		_, err := s.CreateOrUpdateStock(stock)
		if err != nil {
			t.Fatalf("unable to create Stock for list sell testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.StockFilter
		want    []storage.Stock
		wantErr bool
	}{
		{
			name: "LIST_ALL_Stock_SUCCESS",
			in:   storage.StockFilter{},
			want: []storage.Stock{
				{
					ProductId:   product.ID,
					Quantity:    6,
					ProductNane: product.Name,
				},
				{
					ProductId:   product2.ID,
					Quantity:    7,
					ProductNane: product2.Name,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfStock(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListOfStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Stock{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListOfSell() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetProductAndStockByID(t *testing.T) {
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
	newStock := storage.Stock{
		ProductId: product.ID,
		Quantity:  3,
	}
	stock, err := s.CreateOrUpdateStock(newStock)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOrUpdateStock() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.ProductStock
		wantErr bool
	}{
		{
			name: "GET_STOCK_BY_PRODUCT_ID_SUCEESS",
			in:   stock.ID,
			want: &storage.ProductStock{
				ProductId:     stock.ID,
				ProductName:   product.Name,
				UnitPrice:     product.Price,
				StockQuantity: stock.Quantity,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetProductAndStockByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetProductAndStockByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.ProductStock{},"ID"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetProductAndStockByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

