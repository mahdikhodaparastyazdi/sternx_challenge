package model

type Trade struct {
	Symbol                 string
	Open, High, Low, Close int
	LastTransactionDate    string
}
