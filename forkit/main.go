package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/str1ngs/forkit"
	"github.com/str1ngs/util"
	"github.com/str1ngs/util/console"
	"github.com/str1ngs/util/file"
	"github.com/str1ngs/util/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	commands = []*Command{
		&Command{"clone", clone},
		&Command{"status", status},
	}
	jconfig = &Config{
		GithubUser: "",
		GithubHome: "$HOME/github.com",
	}
	chome = os.ExpandEnv("$HOME/.config/forkit")
	jfile = filepath.Join(chome, "config.json")
)

type Config struct {
	GithubUser string
	GithubHome file.Path
}

type Command struct {
	Name string
	Run  func(args []string)
}

func init() {
	log.SetPrefix("forkit: ")
	log.SetFlags(log.Lshortfile)
	if !file.Exists(chome) {
		console.Println("creating config dir", chome)
		err := os.MkdirAll(chome, 0755)
		util.CheckFatal(err)
	}
	if !file.Exists(jfile) {
		console.Println("writing default config", jfile)
		err := json.Write(&jconfig, jfile)
		util.CheckFatal(err)
	}
	console.Flush()
	err := json.Read(&jconfig, jfile)
	util.CheckFatal(err)
	if jconfig.GithubUser == "" {
		fmt.Printf("edit %s and add your github username\n", jfile)
		os.Exit(2)
	}
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

func config(args []string) {
}

func clone(args []string) {
	repos, err := gohub.Repos(jconfig.GithubUser)
	if err != nil {
		log.Fatal(err)
	}
	err = gohub.CloneAll(jconfig.GithubHome.Expand(), repos)
	if err != nil {
		log.Fatal(err)
	}
}

func status(args []string) {
	glob := filepath.Join(jconfig.GithubHome.Expand(), "*")
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
