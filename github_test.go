package gohub

import (
	"github.com/str1ngs/util/file"
	"log"
	"os"
	"testing"
)

var (
	user  = "str1ngs"
	repos []Repo
)

func init() {
	var err error
	repos, err = Repos(user)
	if err != nil {
		log.Fatal(err)
	}
	dir := "./tmp"
	if !file.Exists(dir) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestCloneAll(t *testing.T) {
	err := CloneAll("./tmp", repos)
	if err != nil {
		t.Error(err)
	}
}
