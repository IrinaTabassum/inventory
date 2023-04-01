package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateCustomer(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.Customer
		want    *storage.Customer
		wantErr bool
	}{
		{
			name: "CREATE_CUSTOMER_SUCEESS",
			in: storage.Customer{
				Name: "Customer1",
			},
			want: &storage.Customer{
				Name: "Customer1",
			},
		},
		{
			name: "CREATE_CUSTOMER_NAME_UNIQUE_FAILED",
			in: storage.Customer{
				Name: "Customer1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCustomer(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Customer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateCustomer() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOfCustomer(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	customers := []storage.Customer{
		{
			Name: "Customer1",
		},
		{
			Name: "Customer2",
		},
		{
			Name: "Customer3",
		},
	}

	for _, customer := range customers {
		_, err := s.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("unable to create customer for list customer testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.CustomerFilter
		want    []storage.Customer
		wantErr bool
	}{
		{
			name: "LIST_ALL_CUSTOMER_SUCCESS",
			in:   storage.CustomerFilter{},
			want: []storage.Customer{
				{
					Name: "Customer1",
				},
				{
					Name: "Customer2",
				},
				{
					Name: "Customer3",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfCustomer(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Customer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListCustomer() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "Customer1",
	}

	tests := []struct {
		name    string
		in      storage.Customer
		want    *storage.Customer
		wantErr bool
	}{
		{
			name: "UPDATE_CUSTOMER_SUCEESS",
			in: storage.Customer{
				Name: "CustomerOne",
			},
			want: &storage.Customer{
				Name: "CustomerOne",
			},
		},
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCustomer() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = customer.ID
			got, err := s.UpdateCustomer(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Customer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateCustomer() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestDeleteCustomerByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "Customer1",
	}

	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCustomer() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_CUSTOMER_BY_ID_SUCEESS",
			in:   customer.ID,
		},
		{
			name:    "DELETE_CUSTOMER_BY_ID_FAILED",
			in:      customer.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteCustomerByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetCustomerByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "Customer1",
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCustomer() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		want    *storage.Customer
		wantErr bool
	}{
		{
			name: "GET_CUSTOMER_BY_ID_SUCEESS",
			in:   customer.ID,
			want: &storage.Customer{
				Name: "Customer1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCustomerByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCustomerByCustomername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Customer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateCustomer() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
			
		})
	}
}

func TestGetCustomerByName(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCustomer := storage.Customer{
		Name: "Customer1",
	}
	customer, err := s.CreateCustomer(newCustomer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCustomer() error = %v", err)
	}

	tests := []struct {
		name    string
		in      string
		want    *storage.Customer
		wantErr bool
	}{
		{
			name: "GET_CUSTOMER_BY_NAME_SUCEESS",
			in:   customer.Name,
			want: &storage.Customer{
				Name: "Customer1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCustomerByName(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCustomerByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Customer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCustomerByName() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
			
		})
	}
}
