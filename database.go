package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  m "fgo-test/models"
)

func OpenDatabase() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("Unable to open sqlite db: test.db")
  }

  db.AutoMigrate(&m.User{})

  return db
}
