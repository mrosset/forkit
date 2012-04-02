package main

import (
	"bytes"
	"flag"
	"github.com/str1ngs/forkit"
	"github.com/str1ngs/util/console"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	user = "str1ngs"
	home = "/home/strings/gocode/src/github.com/str1ngs"
)

var commands = []*Command{
	&Command{"clone", clone},
	&Command{"status", status},
}

type Command struct {
	Name string
	Run  func(args []string)
}

func init() {
	log.SetPrefix("gohub: ")
	log.SetFlags(log.Lshortfile)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for _, cmd := range commands {
		if cmd.Name == args[0] {
			cmd.Run(args[1:])
		}
	}
}

func clone(args []string) {
	repos, err := gohub.Repos(user)
	if err != nil {
		log.Fatal(err)
	}
	err = gohub.CloneAll(home, repos)
	if err != nil {
		log.Fatal(err)
	}
}

func status(args []string) {
	glob := filepath.Join(home, "*")
	dirs, err := filepath.Glob(glob)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dirs {
		buf := new(bytes.Buffer)
		git := exec.Command("git", "status", "--porcelain")
		git.Stdout = buf
		git.Stderr = os.Stderr
		git.Dir = d
		err = git.Run()
		if err != nil {
			log.Fatal(err)
		}
		if len(buf.Bytes()) > 0 {
			console.Println(filepath.Base(d), "*")
		}
	}
	console.Flush()
}
