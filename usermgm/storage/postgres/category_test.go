package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.Category
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCEESS",
			in: storage.Category{
				Name: "Phone",
			},
			want: &storage.Category{
				Name: "Phone",
			},
		},
		{
			name: "CREATE_CATEGORY_NAME_UNIQUE_FAILED",
			in: storage.Category{
				Name: "Phone",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListOfCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	categorys := []storage.Category{
		{
			Name: "Phone",
		},
		{
			Name: "Laptop",
		},
		{
			Name: "Tab",
		},
	}

	for _, category := range categorys {
		_, err := s.CreateCategory(category)
		if err != nil {
			t.Fatalf("unable to create category for list category testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.CategoryFilter
		want    []storage.Category
		wantErr bool
	}{
		{
			name: "LIST_ALL_CATEGORY_SUCCESS",
			in:   storage.CategoryFilter{},
			want: []storage.Category{
				{
					Name: "Phone",
				},
				{
					Name: "Laptop",
				},
				{
					Name: "Tab",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListOfCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCategory := storage.Category{
		Name: "Phone",
	}

	tests := []struct {
		name    string
		in      storage.Category
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "UPDATE_Categorty_SUCEESS",
			in: storage.Category{
				Name: "PhoneUpdate",
			},
			want: &storage.Category{
				Name: "PhoneUpdate",
			},
		},
	}
	category, err := s.CreateCategory(newCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = category.ID
			got, err := s.UpdateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
func TestDeleteCategoryByID(t *testing.T) {
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
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_CATEGORY_BY_ID_SUCEESS",
			in:   category.ID,
		},
		{
			name: "DELETE_CATEGORY_BY_ID_FAILED",
			in:      category.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteCategoryByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetCategoryByID(t *testing.T) {
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
		in      int
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_BY_ID_SUCEESS",
			in:   category.ID,
			want: &storage.Category{
				Name: "Phone",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCategoryByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCategoryByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}

		})
	}
}

func TestGetCategoryByName(t *testing.T) {
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
		in      string
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_BY_NAME_SUCEESS",
			in:   category.Name,
			want: &storage.Category{
				Name: "Phone",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCategoryByName(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCategoryByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCategoryByName() diff = %v", cmp.Diff(got, tt.want, opts...))
			}

		})
	}
}
