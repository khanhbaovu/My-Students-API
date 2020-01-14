package main

import (
  "studentProject/driver"
  "fmt"
)

func main() {
  db := driver.Connect("localhost", "5432", "postgres", "Vbk02122000", "postgres")

  err := db.SQL.Ping()

  if err != nil {
    panic(err)
  }

  fmt.Println("Connection!")
}
