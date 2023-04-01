package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateSell(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "customer1",
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSupplier() error = %v", err)
	}

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
		Description: "this is test product",
		Price:       100,
	}
	product2, err := s.CreateProduct(newProduct2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}
	
	tests := []struct {
		name    string
		in      storage.Sell
		want    storage.Sell
		wantErr bool
	}{
		{
			name: "CREATE_SELL_SUCEESS",
			in: storage.Sell{
				CustomerId: customer.ID,
				TotalPrice: (product.Price * 2) + (product2.Price * 3),
			},
			want: storage.Sell{
				CustomerId: customer.ID,
				TotalPrice: (product.Price * 2) + (product2.Price * 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateSell(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateSell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Sell{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateSell() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestListOfSell(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "customer1",
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSupplier() error = %v", err)
	}

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
		Description: "this is test product",
		Price:       100,
	}
	product2, err := s.CreateProduct(newProduct2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}

	sells := []storage.Sell{
		{
			CustomerId: customer.ID,
			TotalPrice: (product.Price * 2) + (product2.Price * 3),
		},
		{
			CustomerId: customer.ID,
			TotalPrice: (product.Price * 4) + (product2.Price * 2),
		},
	}

	for _, sell := range sells {
		_, err := s.CreateSell(sell)
		if err != nil {
			t.Fatalf("unable to create sell for list sell testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.SellFilter
		want    []storage.Sell
		wantErr bool
	}{
		{
			name: "LIST_ALL_SELL_SUCCESS",
			in:   storage.SellFilter{},
			want: []storage.Sell{
				{
					CustomerId: customer.ID,
					CustomerName: customer.Name,
					TotalPrice:   (product.Price * 2) + (product2.Price * 3),
				},
				{
					CustomerId: customer.ID,
					CustomerName: customer.Name,
					TotalPrice: (product.Price * 4) + (product2.Price * 2),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfSell(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListOfSell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Sell{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
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
func TestGetSellByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "customer1",
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSupplier() error = %v", err)
	}

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
		Description: "this is test product",
		Price:       100,
	}
	product2, err := s.CreateProduct(newProduct2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}
	nweSell := storage.Sell{
		CustomerId: customer.ID,
		TotalPrice: (product.Price * 2) + (product2.Price * 3),
	}
	sell,err := s.CreateSell(nweSell)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSell() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.Sell
		wantErr bool
	}{
		{
			name: "GET_SELL_BY_ID_SUCEESS",
			in: sell.ID,
			want: &storage.Sell{
				CustomerId: customer.ID,
				CustomerName: customer.Name,
				TotalPrice: (product.Price * 2) + (product2.Price * 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSellByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Sell{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCategoryByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
