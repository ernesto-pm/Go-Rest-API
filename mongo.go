package main
import(
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Phone string
}

func main(){
  session,err := mgo.Dial("localhost:27017")
  if err != nil{
    panic(err)
  }
  defer session.Close()

  c := session.DB("yourtime").C("vehicles")
  //err = c.Insert(&Person{"Ernesto", "+55 53 8116 9639"},&Person{"Keki", "+55 53 8402 8510"})
  /*
  if err != nil{
    log.Fatal(err)
  }
  */

  var result []Person
  err = c.Find(bson.M{}).All(&result)
  if err != nil{
    log.Fatal("error", err)
  }

  for _,person := range result{
    fmt.Println("Name: ",person.Name)
    fmt.Println("Phone number: ", person.Phone)
  }

  //fmt.Println("Results:", result)
}
