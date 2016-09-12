package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "localhost:28015", "-url url:port")
}
func main() {
	flag.Parse()
	session := connect(url)
	start := time.Now()
	for i := 0; i < 1000; i++ {
		insert(session)
	}
	elapsed := time.Since(start)
	sec := elapsed / time.Second
	fmt.Printf("writes/sec %d", 1000/sec)
}

func connect(url string) *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func insert(session *r.Session) {
	insertData := struct {
		Name  string
		Value string
	}{
		Name:  "hihi",
		Value: "test",
	}

	_, err := r.DB("test").Table("insert_test").Insert(insertData).RunWrite(session)
	if err != nil {
		log.Fatal(err)
	}
}
