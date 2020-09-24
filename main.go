package main

import(
  "fmt"
  "log"
  "time"
  "errors"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "github.com/gorilla/mux"
  "github.com/golang/gddo/httputil/header"
)

type Entry struct {
  gorm.Model
  Diastolic  uint
  Systolic   uint
  Heartrate  uint
}

type RequestStatus struct {
  Status  string  `json:"status"`
}

func checkHeader(r *http.Request) error {
  if r.Header.Get("Content-Type") != "" {
    ctype, _ := header.ParseValueAndParams(r.Header, "Content-Type")
    if ctype != "application/json" {
      return errors.New("Content-Type must be application/json")
    }
  } else {
    return errors.New("No Content-Type header provided in request")
  }
  return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  err := checkHeader(r)
  if err != nil {
    msg, _ := json.Marshal(RequestStatus{Status: "Error: " + err.Error()})
    http.Error(w, string(msg), http.StatusBadRequest)
    return
  }
  resp, _ := json.Marshal(RequestStatus{Status: "OK"})
  w.Header().Add("Content-Type", "application/json")
  w.Write(resp)
}

// func NewEntryHandler(w http.ResponseWriter, r *http.Request) {
//
// }

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
