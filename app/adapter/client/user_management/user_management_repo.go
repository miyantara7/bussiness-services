package user_management

type UserManagementRepo interface {
	Login(in interface{}) error
}
