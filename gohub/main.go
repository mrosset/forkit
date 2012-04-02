package main

import (
	"github.com/str1ngs/gohub"
	"log"
)

var (
	user  = "str1ngs"
	repos []gohub.Repo
)

func init() {
}

func main() {
	var err error
	repos, err = gohub.Repos(user)
	if err != nil {
		log.Fatal(err)
	}
	err = gohub.CloneAll("/home/strings/gocode/src/github.com/", repos)
	if err != nil {
		log.Fatal(err)
	}
}
