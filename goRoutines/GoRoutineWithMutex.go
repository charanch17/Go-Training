package main

import (
	"fmt"
	"sync"
)

var TransactionTable []Transaction
var wg sync.WaitGroup
var mutex sync.Mutex

type Transaction struct {
	TransactionId         int
	TransactionType       string
	TransactionDate       string
	ExchangeTransactionId int
}

func AddPurchaseTransactions(id int, Date string) {
	defer wg.Done()
	mutex.Lock()
	TransactionTable = append(TransactionTable, Transaction{TransactionId: id, TransactionType: "Purchase", TransactionDate: Date})
	mutex.Unlock()
	fmt.Println("added purchase transaction")
}

func AddSellTransaction(id int, Date string) {
	defer wg.Done()
	mutex.Lock()
	TransactionTable = append(TransactionTable, Transaction{TransactionId: id, TransactionType: "Sell", TransactionDate: Date})
	mutex.Unlock()
	fmt.Println("added sell transaction")

}

func AddExchangeTransaction(ExchangeTransactionId int, id int, Date string) {
	defer wg.Done()
	mutex.Lock()
	TransactionTable = append(TransactionTable, Transaction{TransactionId: id, TransactionType: "Exchange", TransactionDate: Date, ExchangeTransactionId: ExchangeTransactionId})
	mutex.Unlock()
	fmt.Println("added Exchange transaction")

}

func main() {

	wg.Add(5)

	go AddPurchaseTransactions(1, "2022-01-01")
	go AddExchangeTransaction(2, 1, "2022-02-02")
	go AddSellTransaction(1, "2022-12-12")
	go AddExchangeTransaction(3, 2, "2022-02-02")
	go AddSellTransaction(5, "2022-12-12")

	wg.Wait()
	fmt.Println(TransactionTable)

}
