package postgres

import (
	"sort"
	"testing"

	"codemen.org/inventory/usermgm/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	
)

func TestUserRegister(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_USER_SUCEESS",
			in: storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Password:  "123",
			},
			want: &storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Role:      "user",
				IsActive:  true,
			},
		},
		{
			name: "CREATE_USER_EMIAL_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina2",
				Password:  "123",
			},
			wantErr: true,
		},
		{
			name: "CREATE_USER_USERNAME_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina2@gmail.com",
				Username:  "irina",
				Password:  "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Register(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Irina",
		LastName:  "Tabassum",
		Email:     "irina@gmail.com",
		Username:  "irina",
		Password:  "123",
	}

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "UPDATE_USER_SUCEESS",
			in: storage.User{
				FirstName: "Irinaupdate",
				LastName:  "Tabassumupdate",
				IsActive:  false,
			},
			want: &storage.User{
				FirstName: "Irinaupdate",
				LastName:  "Tabassumupdate",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Role:      "user",
				IsActive:  false,
			},
		},
	}
	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.Register() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = user.ID
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	users := []storage.User{
		{
			FirstName: "Irian",
			LastName:  "Tabassum",
			Email:     "irina@gmail.com",
			Username:  "irina",
			Password:  "12345678",
		},
		{
			FirstName: "Humira",
			LastName:  "Rahaman",
			Email:     "rahaman@gmail.com",
			Username:  "rahaman",
			Password:  "12345678",
		},
		{
			FirstName: "Rijoanul",
			LastName:  "Hasan",
			Email:     "rijoanul@gmail.com",
			Username:  "rijoanul",
			Password:  "12345678",
		},
	}

	for _, user := range users {
		_, err := s.Register(user)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.User
		wantErr bool
	}{
		{
			name: "LIST_ALL_USER_SUCCESS",
			in:   storage.UserFilter{},
			want: []storage.User{
				{
					FirstName: "Irian",
					LastName:  "Tabassum",
					Email:     "irina@gmail.com",
					Username:  "irina",
					Role:      "user",
					IsActive:  true,
				},
				{
					FirstName: "Humira",
					LastName:  "Rahaman",
					Email:     "rahaman@gmail.com",
					Username:  "rahaman",
					Role:      "user",
					IsActive:  true,
				},
				{
					FirstName: "Rijoanul",
					LastName:  "Hasan",
					Email:     "rijoanul@gmail.com",
					Username:  "rijoanul",
					Role:      "user",
					IsActive:  true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Humira",
		LastName:  "Rahaman",
		Email:     "rahaman@gmail.com",
		Username:  "rahaman",
		Password:  "12345678",
	}

	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name: "DELETE_USER_BY_ID_SUCEESS",
			in:   user.ID,
		},
		{
			name:    "DELETE_USER_BY_ID_FAILED",
			in:      user.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_,err := s.DeleteUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Irina",
		LastName:  "Tabassum",
		Email:     "irina@gmail.com",
		Username:  "irina",
		Password:  "12345678",
	}
	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_USERNAME_SUCEESS",
			in:   user.Username,
			want: &storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Role:      "user",
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByUsername(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
			
		})
	}
}

func TestGetUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Irina",
		LastName:  "Tabassum",
		Email:     "irina@gmail.com",
		Username:  "irina",
		Password:  "12345678",
	}
	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	tests := []struct {
		name    string
		in      int
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_ID_SUCEESS",
			in:   user.ID,
			want: &storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Role:      "user",
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
			
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Irina",
		LastName:  "Tabassum",
		Email:     "irina@gmail.com",
		Username:  "irina",
		Password:  "12345678",
	}
	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}

	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_Email_SUCEESS",
			in:   user.Email,
			want: &storage.User{
				FirstName: "Irina",
				LastName:  "Tabassum",
				Email:     "irina@gmail.com",
				Username:  "irina",
				Role:      "user",
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByEmail(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt", "Total"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetUserByEmail() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
			
		})
	}
}