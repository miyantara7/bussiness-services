package user_management_services

type UserManagementRepo interface {
	Login(in interface{}) error
	Register(in interface{}) error
	UserInformation(in interface{}) error
}
