package main

type Transaction struct {
  Name  string `json:"name"`
  Value float64 `json:"value"`
  IsRecurrent bool `json:"is_recurrent"`
}

type Transactions []Transaction
