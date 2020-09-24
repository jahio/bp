package main

import(
  "fmt"
  "log"
  "time"
  "net/http"
  "encoding/json"
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

type RequestStatus struct {
  Status  string
}

func HTTPError(w http.ResponseWriter, err error) {
  log.Println("Error in application:", err)
  resp, err := json.Marshal(RequestStatus{Status: "Error"})
  if err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusInternalServerError)
  w.Write(resp)
  return
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  resp, err := json.Marshal(RequestStatus{Status: "OK"})
  if err != nil {
    HTTPError(w, err)
    return
  }
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(resp)
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