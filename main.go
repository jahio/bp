package main

import(
  "fmt"
  "log"
  "time"
  "net/http"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "github.com/gorilla/mux"
)

type Entry struct {
  gorm.Model
  Diastolic  uint
  Systolic   uint
  Heartrate  uint
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "{ \"status\":\"OK\" }")
}

func main() {
  fmt.Println("Beginning database automigrate...")
  db, err := gorm.Open(sqlite.Open("bp.db"), &gorm.Config{})
  if err != nil {
    panic("Cannot open/create SQLite database bp.db!")
  }

  db.AutoMigrate(&Entry{})

  // Set up HTTP routing for requests/responses
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  srv := &http.Server{
    Handler: r,
    Addr:    "0.0.0.0:9000",
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }
  
  log.Fatal(srv.ListenAndServe())
}