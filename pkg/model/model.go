package model

type Transaction struct {
	Id         int64   `json:"id",db:"id"`
	FromId     int64   `json:"from_id",db:"from_id"`
	ToId       int64   `json:"to_id",db:"to_id"`
	Amount     float64 `json:"amount",db:"amount"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type Wallet struct {
	Id         int64   `json:"id",db:"id"`
	UserId     int64   `json:"user_id",db:"user_id"`
	Balance    float64 `json:"balance",db:"balance"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type Withdraw struct {
	UserId int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Deposit struct {
	UserId int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Transfer struct {
	FromId int64   `json:"from_id"`
	ToId   int64   `json:"to_id"`
	Amount float64 `json:"amount"`
}

type TransactionDTO struct {
	Transactions []*Transaction `json:"transactions"`
	Total        int64          `json:"total"`
}
