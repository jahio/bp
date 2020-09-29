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
  "github.com/golang/gddo/httputil/header"
)

type Entry struct {
  gorm.Model
  Diastolic  uint  `json:diastolic`
  Systolic   uint  `json:systolic`
  Heartrate  uint  `json:heartrate`
}

type RequestStatus struct {
  Status  string  `json:"status"`
  Message string  `json:"message,omitempty"`
}

func jsonParseError(w http.ResponseWriter, r *http.Request) {
  msg, _ := json.Marshal(RequestStatus{Status: "JSON Parse Error"})
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusBadRequest)
  w.Write(msg)
}

func applicationError(w http.ResponseWriter, r *http.Request, errMsg string) {
  msg, _ := json.Marshal(RequestStatus{Status: "Error", Message: errMsg})
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusInternalServerError)
  w.Write(msg)
}

func checkHeaderMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RemoteAddr, r.Method, r.RequestURI) // stdout logging
    if r.Header.Get("Content-Type") != "" {
      ctype, _ := header.ParseValueAndParams(r.Header, "Content-Type")
      if ctype != "application/json" {
        msg, err := json.Marshal(RequestStatus{Status: "Error", Message: "Content-Type must be application/json"})
        if err != nil {
          log.Println(err.Error())
        }
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(msg)
        return
      }
    } else {
      msg, err := json.Marshal(RequestStatus{Status: "Error", Message: "No Content-Type header provided in request"})
      if err != nil {
        log.Println(err.Error())
      }
      w.Header().Add("Content-Type", "application/json")
      w.WriteHeader(http.StatusBadRequest)
      w.Write(msg)
      return
    }
    // Headers look good
    next.ServeHTTP(w, r)
  })
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  resp, _ := json.Marshal(RequestStatus{Status: "OK"})
  w.Header().Add("Content-Type", "application/json")
  w.Write(resp)
}

func NewEntryHandler(w http.ResponseWriter, r *http.Request) {
  var entry Entry
  err := json.NewDecoder(r.Body).Decode(&entry)
  if err != nil {
    jsonParseError(w, r)
    log.Println("JSON parse error: ", r.RequestURI)
    return
  }

  db, err := gorm.Open(sqlite.Open("bp.db"), &gorm.Config{})
  if err != nil {
    applicationError(w, r, "There is a problem with the application database")
    log.Println("Could not open bp.db for request")
    return
  }

  // Try to create the object
  db.Create(&entry)

  msg, _ := json.Marshal(entry)
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(msg)
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
  r.HandleFunc("/entries/new", NewEntryHandler)
  r.Use(checkHeaderMiddleware)
  srv := &http.Server{
    Handler:           r,
    Addr:              "0.0.0.0:9000",
    WriteTimeout:      15 * time.Second,
    ReadTimeout:       15 * time.Second,
    IdleTimeout:       15 * time.Second,
    ReadHeaderTimeout: 15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}
