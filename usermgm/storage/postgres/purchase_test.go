package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreatePurchase(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newSupplier := storage.Supplier{
		Name: "Supplier1",
	}
	supplier, err := s.CreateSupplier(newSupplier)
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

	tests := []struct {
		name    string
		in      storage.Purchase
		want    *storage.Purchase
		wantErr bool
	}{
		{
			name: "CREATE_PURCHASE_SUCEESS",
			in: storage.Purchase{
				SupplierId: supplier.ID,
				ProductId:  product.ID,
				Quantity:   3,
				UnitPrice:  product.Price,
				TotalPrice: product.Price * 3,
			},
			want: &storage.Purchase{
				SupplierId: supplier.ID,
				ProductId:  product.ID,
				Quantity:   3,
				UnitPrice:  product.Price,
				TotalPrice: product.Price * 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreatePurchase(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreatePurchase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Purchase{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreatePurchase() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOfPurchase(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newSupplier := storage.Supplier{
		Name: "Supplier1",
	}
	supplier, err := s.CreateSupplier(newSupplier)
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

	purchases := []storage.Purchase{
		{
			SupplierId: supplier.ID,
			ProductId:  product.ID,
			Quantity:   4,
			UnitPrice:  product.Price,
			TotalPrice: product.Price * 4,
		},
		{
			SupplierId: supplier.ID,
			ProductId:  product.ID,
			Quantity:   3,
			UnitPrice:  product.Price,
			TotalPrice: product.Price * 4,
		},
	}

	for _, purchase := range purchases {
		_, err := s.CreatePurchase(purchase)
		if err != nil {
			t.Fatalf("unable to create purchase for list purchase testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.PurchaseFilter
		want    []storage.Purchase
		wantErr bool
	}{
		{
			name: "LIST_ALL_PURCHASE_SUCCESS",
			in:   storage.PurchaseFilter{},
			want: []storage.Purchase{
				{
					SupplierId: supplier.ID,
					SupplierName: supplier.Name,
					ProductId:  product.ID,
					Quantity:   4,
					UnitPrice:  product.Price,
					TotalPrice: product.Price * 4,
				},
				{
					SupplierId: supplier.ID,
					SupplierName: supplier.Name,
					ProductId:  product.ID,
					Quantity:   3,
					UnitPrice:  product.Price,
					TotalPrice: product.Price * 4,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfPurchase(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListPurchase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Purchase{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListPurchase() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestUpdatePurchase(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newSupplier := storage.Supplier{
		Name: "Supplier1",
	}
	supplier, err := s.CreateSupplier(newSupplier)
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
		Description: "this is test product2",
		Price:       100,
	}
	product2, err := s.CreateProduct(newProduct2)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}
	newPurchase := storage.Purchase{
		SupplierId: supplier.ID,
		ProductId:  product.ID,
		Quantity:   3,
		UnitPrice:  product.Price,
		TotalPrice: product.Price * 3,
	}

	tests := []struct {
		name    string
		in      storage.Purchase
		want    *storage.Purchase
		wantErr bool
	}{
		{
			name: "UPDATE_PURCHASE_SUCEESS",
			in: storage.Purchase{
				SupplierId: supplier.ID,
				ProductId:  product2.ID,
				Quantity:   4,
				UnitPrice:  product.Price,
				TotalPrice: product.Price * 4,
			},
			want: &storage.Purchase{
				SupplierId: supplier.ID,
				ProductId:  product2.ID,
				Quantity:   4,
				UnitPrice:  product.Price,
				TotalPrice: product.Price * 4,
			},
		},
	}
	purchase, err := s.CreatePurchase(newPurchase)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = purchase.ID
			got, err := s.UpdatePurchase(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdatePurchase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Purchase{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdatePurchase() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestGetPurchaseByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newSupplier := storage.Supplier{
		Name: "Supplier1",
	}
	supplier, err := s.CreateSupplier(newSupplier)
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

	newPurchase := storage.Purchase{
		SupplierId: supplier.ID,
		ProductId:  product.ID,
		Quantity:   3,
		UnitPrice:  product.Price,
		TotalPrice: product.Price * 3,
	}
	purchase, err := s.CreatePurchase(newPurchase)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateProduct() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		want    *storage.Purchase
		wantErr bool
	}{
		{
			name: "UPDATE_PURCHASE_SUCEESS",
			in:   purchase.ID,
			want: &storage.Purchase{
				SupplierId: supplier.ID,
				ProductId:  product.ID,
				Quantity:   3,
				UnitPrice:  product.Price,
				TotalPrice: product.Price * 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetPurchaseByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetPurchaseByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Purchase{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetPurchaseByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

