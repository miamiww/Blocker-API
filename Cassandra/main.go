package Cassandra
import (
  "github.com/gocql/gocql"
  "fmt"
)

var Session *gocql.Session //making Session accessible outside of the package
func init() {
  var err error
  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "ipdatabase"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}