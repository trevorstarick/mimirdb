package main

import "github.com/trevorstarick/mimirdb"

func main() {
	s := mimirdb.Server{}

	err := s.ListenAndServe(":4444")
	if err != nil {
		panic(err)
	}
}