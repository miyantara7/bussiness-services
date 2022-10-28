package top_up

import "context"

type TopUp interface {
	TopUp(context.Context, interface{}) error
	Payment(context.Context, interface{}) error
}
