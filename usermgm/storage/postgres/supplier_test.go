package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateSupplier(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.Supplier
		want    *storage.Supplier
		wantErr bool
	}{
		{
			name: "CREATE_SUPPLIER_SUCEESS",
			in: storage.Supplier{
				Name: "Supplier1",
			},
			want: &storage.Supplier{
				Name: "Supplier1",
			},
		},
		{
			name: "CREATE_SUPPLIER_NAME_UNIQUE_FAILED",
			in: storage.Supplier{
				Name: "Supplier1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateSupplier(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Supplier{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateSupplier() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOfSupplier(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	suppliers := []storage.Supplier{
		{
			Name: "Supplier1",
		},
		{
			Name: "Supplier2",
		},
		{
			Name: "Supplier3",
		},
	}

	for _, supplier := range suppliers {
		_, err := s.CreateSupplier(supplier)
		if err != nil {
			t.Fatalf("unable to create supplier for list supplier testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.SupplierFilter
		want    []storage.Supplier
		wantErr bool
	}{
		{
			name: "LIST_ALL_SUPPLIER_SUCCESS",
			in:   storage.SupplierFilter{},
			want: []storage.Supplier{
				{
					Name: "Supplier1",
				},
				{
					Name: "Supplier2",
				},
				{
					Name: "Supplier3",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfSupplier(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListSupplier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Supplier{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListSupplier() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateSupplier(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newSupplier := storage.Supplier{
		Name: "Supplier1",
	}

	tests := []struct {
		name    string
		in      storage.Supplier
		want    *storage.Supplier
		wantErr bool
	}{
		{
			name: "UPDATE_SUPPLIER_SUCEESS",
			in: storage.Supplier{
				Name: "SupplierOne",
			},
			want: &storage.Supplier{
				Name: "SupplierOne",
			},
		},
	}
	supplier, err := s.CreateSupplier(newSupplier)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateSupplier() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = supplier.ID
			got, err := s.UpdateSupplier(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateSupplier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Supplier{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateSupplier() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestDeleteSupplierByID(t *testing.T) {
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

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_SUPPLIER_BY_ID_SUCEESS",
			in:   supplier.ID,
		},
		{
			name:    "DELETE_SUPPLIER_BY_ID_FAILED",
			in:      supplier.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteSupplierByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteSupplierByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetSupplierByID(t *testing.T) {
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

	tests := []struct {
		name    string
		in      int
		want    *storage.Supplier
		wantErr bool
	}{
		{
			name: "GET_SUPPLIER_BY_ID_SUCEESS",
			in:   supplier.ID,
			want: &storage.Supplier{
				Name: "Supplier1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSupplierByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetSupplierBySuppliername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Supplier{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateSupplier() diff = %v", cmp.Diff(got, tt.want, opts...))
			}

		})
	}
}

func TestGetSupplierByName(t *testing.T) {
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
	tests := []struct {
		name    string
		in      string
		want    *storage.Supplier
		wantErr bool
	}{
		{
			name: "GET_SUPPLIER_BY_NAME_SUCEESS",
			in:   supplier.Name,
			want: &storage.Supplier{
				Name: "Supplier1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSupplierByName(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetSupplierByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Supplier{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetSupplierByName() diff = %v", cmp.Diff(got, tt.want, opts...))
			}

		})
	}
}
