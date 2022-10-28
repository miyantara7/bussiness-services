package top_up_client

type TopUpRepo interface {
	TopUp(interface{}) error
	Payment(interface{}) error
}
