package model

type TopUpReq struct {
	UserId   string
	UserName string
	Balance  string
	NoKartu  string
}

type PaymentReq struct {
	UserId   string
	UserName string
	BillerId string
	NoKartu  string
}
