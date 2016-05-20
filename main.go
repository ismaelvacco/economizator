package main

import (
//  "fmt"
  "net/http"
  "encoding/json"

  "appengine"
  "appengine/datastore"
)


func init() {
    http.HandleFunc("/", getTransactions)
    http.HandleFunc("/transactions", getTransactions)
    http.HandleFunc("/transaction", createTransaction)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  q := datastore.NewQuery("Transaction")
  var transactions Transactions
  for t:= q.Run(ctx); ; {
    var a Transaction
    _, err := t.Next(&a)

    if err == datastore.Done {
      break
    }

    if err != nil {
      http.Error(w,err.Error(), 500)
      return
    }

      transactions = append(transactions, a)
  }
  
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w,"POST requests only", http.StatusMethodNotAllowed)
    return
  }

  ctx := appengine.NewContext(r)
  k := datastore.NewKey(ctx, "Transaction", "", 0, nil)
  e := new(Transaction)
  e.Name = "Pagar Carro"
  e.Value = 1502.56
  e.IsRecurrent = false

  if _, err := datastore.Put(ctx, k, e); err != nil {
    http.Error(w,err.Error(), 500)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(e)

}
