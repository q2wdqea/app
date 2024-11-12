package model

type TransactionType int

const (
	TransactionDeposit TransactionType = iota
	TransactionWithdraw
	TransactionTransfer
)

type Transaction struct {
	Id         int64  `json:"id",db:"id"`
	UserId     int64  `json:"user_id",db:"user_id"`
	BizType    int64  `json:"biz_type",db:"biz_type"`
	BizId      int64  `json:"biz_id",db:"biz_id"`
	CreateTime string `json:"create_time",db:"create_time"`
}

type Transfer struct {
	Id         int64   `json:"id",db:"id"`
	FromId     int64   `json:"from_id",db:"from_id"`
	ToId       int64   `json:"to_id",db:"to_id"`
	Amount     float64 `json:"amount",db:"amount"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type Withdraw struct {
	Id         int64   `json:"id",db:"id"`
	UserId     int64   `json:"user_id",db:"user_id"`
	Amount     float64 `json:"amount",db:"amount"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type Deposit struct {
	Id         int64   `json:"id",db:"id"`
	UserId     int64   `json:"user_id",db:"user_id"`
	Amount     float64 `json:"amount",db:"amount"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type Wallet struct {
	Id         int64   `json:"id",db:"id"`
	UserId     int64   `json:"user_id",db:"user_id"`
	Balance    float64 `json:"balance",db:"balance"`
	CreateTime string  `json:"create_time",db:"create_time"`
}

type WithdrawRequest struct {
	UserId int64   `json:"user_id" binding:"required,gt=0"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type DepositRequest struct {
	UserId int64   `json:"user_id" binding:"required,gt=0"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type TransferRequest struct {
	FromId int64   `json:"from_id" binding:"required,gt=0"`
	ToId   int64   `json:"to_id" binding:"required,gt=0"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type TransactionDTO struct {
	Id         int64   `json:"id"`
	UserId     int64   `json:"user_id"`
	FromId     int64   `json:"from_id"`
	ToId       int64   `json:"to_id"`
	BizType    string  `json:"biz_type"`
	Amount     float64 `json:"amount"`
	CreateTime string  `json:"create_time"`
}

type AsyncTransactions struct {
	Transfers []*TransactionDTO
	Withdraws []*TransactionDTO
	Deposits  []*TransactionDTO
}

type TransactionListDTO struct {
	Transactions []*TransactionDTO `json:"transactions"`
	Total        int64             `json:"total"`
}

func (t TransactionType) String() string {
	switch t {
	case TransactionDeposit:
		return "Deposit"
	case TransactionWithdraw:
		return "Withdraw"
	case TransactionTransfer:
		return "Transfer"
	}
	return "Unknown"
}

func (t TransactionType) Int64() int64 {
	switch t {
	case TransactionDeposit:
		return 0
	case TransactionWithdraw:
		return 1
	case TransactionTransfer:
		return 2
	}
	return -1
}

type TransactionDTOS []*TransactionDTO

func (s TransactionDTOS) Len() int { return len(s) }

func (s TransactionDTOS) Less(i, j int) bool {
	return s[i].CreateTime > s[j].CreateTime

}

func (s TransactionDTOS) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
