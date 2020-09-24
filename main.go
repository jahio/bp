package main

import(
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Entry struct {
  gorm.Model
  Diastolic  uint
  Systolic   uint
  Heartrate  uint
}

func main() {
  fmt.Println("Beginning database automigrate...")
  db, err := gorm.Open(sqlite.Open("bp.db"), &gorm.Config{})
  if err != nil {
    panic("Cannot open/create SQLite database bp.db!")
  }

  db.AutoMigrate(&Entry{})
}