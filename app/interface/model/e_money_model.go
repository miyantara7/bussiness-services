package model

type BillerRequest struct {
	ID string
}

type Biller struct {
	ID          int
	Category    string
	Product     string
	Description string
	Price       int
	Fee         int
}

type GetBalanceRequest struct {
	UserId   string
	UserName string
	NoKartu  string
}

type CreateEMoneyRequest struct {
	UserId   string
	UserName string
	Saldo    string
}

type GetTrxHistReq struct {
	UserId   string
	UserName string
	NoKartu  string
}
