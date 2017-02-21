ckage main

import(
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main(){
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/",Index)
  router.HandleFunc("/messages",MessageIndex)
  router.HandleFunc("/messages/{messageId}",MessageShow)
  log.Fatal (http.ListenAndServe(":8016",router))
}

func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func MessageIndex(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Messages Index!")
}

func MessageShow(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  messageId := vars["messageId"]
  fmt.Fprintln(w, "Message shown:", messageId)
}

// ******** Annotations ********

/* Creating custom server to change address
s := &http.Server{
  Addr: ":8016",
}
log.Fatal(s.ListenAndServe())
*/

/* Standard way of handling routes
http.HandleFunc("/",index)
http.HandleFunc("/something",something)
*/

