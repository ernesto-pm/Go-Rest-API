package main

import(
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func something(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hi inside something else")
}

func main(){
  http.HandleFunc("/",index)
  http.HandleFunc("/something",something)
  log.Fatal (http.ListenAndServe(":8016",nil))
}

// ******** Annotations ********

/* Creating custom server to change address
s := &http.Server{
  Addr: ":8016",
}
log.Fatal(s.ListenAndServe())
*/
