package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	url   string
	count int
	dur   string
)

func init() {
	flag.StringVar(&dur, "dur", "hard", "-dur soft")
	flag.IntVar(&count, "count", 1000, "-count 1000")
}
func main() {
	flag.Parse()
	urls := flag.Args()
	session := connect(urls)
	r.Branch(
		r.DB("test").TableList().Contains("insert_test"),
		r.DB("test").Table("insert_test").Delete(),
		r.DB("test").TableCreate("insert_test"),
	).Run(session)
	start := time.Now()
	for i := 0; i < count; i++ {
		insert(session, dur)
	}
	elapsed := time.Since(start)
	sec := int(elapsed / time.Second)
	fmt.Printf("writes/sec %d\n", count/sec)
}

func connect(urls []string) *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Addresses: urls,
	})
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func insert(session *r.Session, dur string) {
	insertData := struct {
		Name  string
		Value string
	}{
		Name:  "hihi",
		Value: "test",
	}

	_, err := r.
		DB("test").
		Table("insert_test").
		Insert(insertData, r.InsertOpts{Durability: dur}).
		RunWrite(session)
	if err != nil {
		log.Fatal(err)
	}
}
