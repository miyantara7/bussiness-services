package e_money

import "context"

type EMoney interface {
	ListBiller() (interface{}, error)
	DetailBiller(context.Context, interface{}) (interface{}, error)
	GetBalance(context.Context, interface{}) (interface{}, error)
	CreateEMoney(context.Context, interface{}) (interface{}, error)
	GetTrxHist(context.Context, interface{}) (interface{}, error)
}
