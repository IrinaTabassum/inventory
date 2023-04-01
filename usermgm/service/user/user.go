package user

import (
	"context"
	"fmt"

	userpb "codemen.org/inventory/gunk/v1/user"
	"codemen.org/inventory/usermgm/storage"
)

type CoreUser interface {
	Register(storage.User) (*storage.User, error)
	Login(storage.Login) (*storage.User, error)
	ListUser(storage.UserFilter) ([]storage.User, error)
	UpdateUser(storage.User) (*storage.User, error)
	DeleteUserByID(int) (storage.User, error)
	GetUserByID(int) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	GetUserByEmail(string) (*storage.User, error)
}

type UserSvc struct {
	userpb.UnimplementedUserServiceServer
	core CoreUser
}

func NewUserSvc(cu CoreUser) *UserSvc {
	return &UserSvc{
		core: cu,
	}
}

func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
	}
	if err := user.Validate(); err != nil {
		fmt.Println(err)
		return nil, err
	}
    
	usename,_ := us.core.GetUserByUsername(user.Username)
	if usename != nil{
		return nil, fmt.Errorf("username is exist")
	}
	useremail,_ := us.core.GetUserByEmail(user.Email)
	if useremail != nil{
		return nil, fmt.Errorf("this email already is exist")
	}


	u, err := us.core.Register(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.Login(login)
	if err != nil {
		return nil, err
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) Userlist(ctx context.Context, r *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	userFilts := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}

	userList, err := us.core.ListUser(userFilts)
	if err != nil {
		return nil, err
	}

	list := make([]*userpb.User, len(userList))
	for i, ul := range userList {
		list[i] = &userpb.User{
			ID:        int32(ul.ID),
			FirstName: ul.FirstName,
			LastName:  ul.LastName,
			Username:  ul.Username,
			Email:     ul.Email,
			IsActive:  ul.IsActive,
		}
	}
	return &userpb.ListUserResponse{
		Users: list,
	}, nil
}
func (us UserSvc) UpdateUser(ctx context.Context, r *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := storage.User{
		ID:        int(r.GetID()),
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		IsActive:  r.GetIsActive(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) DelateUser(ctx context.Context, r *userpb.DelateUserRequest) (*userpb.DelateUserResponse, error) {
	sId := int(r.GetID())

	_, err := us.core.DeleteUserByID(sId)
	if err != nil {
		return nil, err
	}

	return &userpb.DelateUserResponse{}, nil
}

func (us UserSvc) GetUser(ctx context.Context, r *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	sId := int(r.GetID())

	u, err := us.core.GetUserByID(sId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
		},
	}, nil
}
