package user_management

import "context"

type UserManagement interface {
	LoginUser(context.Context, interface{}) (interface{}, error)
	RegisterUser(context.Context, interface{}) error
	GetDetailUserInformation(context.Context, interface{}) (interface{}, error)
}
