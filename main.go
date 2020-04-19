package main

import ("fmt"
        "webScraper_GO/banking"
        )

func main() {
  account := banking.Account{Owner:"jaewon", Balance:5000}
  print(account)
}