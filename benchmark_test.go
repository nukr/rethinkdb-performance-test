package main

import "testing"
import r "gopkg.in/dancannon/gorethink.v2"

var session = connect("localhost:28015")

func BenchmarkInsertALot(b *testing.B) {
	b.SetParallelism(10)
	for i := 0; i < 5000; i++ {
		insertData := struct {
			Name  string
			Value string
		}{
			Name:  "hihi",
			Value: "test",
		}
		r.DB("test").Table("insert_test").Insert(insertData, r.InsertOpts{Durability: "soft"}).RunWrite(session)
	}
}
