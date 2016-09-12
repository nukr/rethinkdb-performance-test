package main

import (
	"log"

	r "gopkg.in/dancannon/gorethink.v2"
)

func main() {
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

	r.DB("test").Table("insert_test").Insert(insertData).RunWrite(session)
}
