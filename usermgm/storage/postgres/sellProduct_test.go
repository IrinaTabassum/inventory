package postgres

import (
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateSellProduct(t *testing.T) {
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
	sell, err := s.CreateSell(nweSell)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSell() error = %v", err)
	}

	tests := []struct {
		name    string
		in      storage.SellProduct
		want    *storage.SellProduct
		wantErr bool
	}{
		{
			name: "CREATE_SELL_PRODUCT_SUCEESS",
			in: storage.SellProduct{
				SellId:         sell.ID,
				ProductId:      product.ID,
				Quantity:       3,
				UnitPrice:      product.Price,
				TotalUnitPrice: product.Price * 3,
			},
			want: &storage.SellProduct{
				SellId:         sell.ID,
				ProductId:      product.ID,
				Quantity:       3,
				UnitPrice:      product.Price,
				TotalUnitPrice: product.Price * 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateSellProduct(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateSellProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.SellProduct{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateSellProduct() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetSellProductByID(t *testing.T) {
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
	sell, err := s.CreateSell(nweSell)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSell() error = %v", err)
	}
	sp := storage.SellProduct{
		SellId:         sell.ID,
		ProductId:      product.ID,
		Quantity:       3,
		UnitPrice:      product.Price,
		TotalUnitPrice: product.Price * 3,
	}
	_, err = s.CreateSellProduct(sp)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSellProduct() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		want    []storage.SellProduct
		wantErr bool
	}{
		{
			name: "GET_SELL_PRODUCT_BY_ID_SUCEESS",
			in:   sp.SellId,
			want: []storage.SellProduct{
				{
					SellId:         sell.ID,
					ProductId:      product.ID,
					ProductName:    product.Name,
					Quantity:       3,
					UnitPrice:      product.Price,
					TotalUnitPrice: product.Price * 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSellProductByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetSellProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.SellProduct{}, "ID", "CreatedAt", "UpdatedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetSellProductByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
