package main


import (
  "net/http"
  "io"
  "log"
  "encoding/json"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/drive/v3"
)

// template of server
func main() {

  mux := http.NewServeMux()



  mux.HandleFunc("/getFiles", getFilesHandler)
  mux.HandleFunc("/login", loginHandler)

  s := &http.Server{
    Addr: ":8080",
    Handler: mux,
  }
  s.ListenAndServe()

}

func loginHandler(res http.ResponseWriter, req *http.Request) {
}

func getFilesHandler(res http.ResponseWriter, req *http.Request) {
  
}

// user 

