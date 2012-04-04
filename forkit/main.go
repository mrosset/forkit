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
		&Command{"clone", clone, "clone user github repos"},
		&Command{"status", status, "check dirty statuf of all repos"},
	}
	jconfig = &Config{
		User: "",
		Home: "$HOME/github.com",
	}
	chome = os.ExpandEnv("$HOME/.config/forkit")
	jfile = filepath.Join(chome, "config.json")
)

type Config struct {
	User string
	Home file.Path
}

type Command struct {
	Name  string
	Run   func()
	Usage string
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
	if jconfig.User == "" {
		fmt.Printf("edit %s and add your github username\n", jfile)
		os.Exit(2)
	}
}

func usage() {
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for _, cmd := range commands {
		if cmd.Name == args[0] {
			cmd.Run()
		}
	}
}

func clone() {
	repos, err := gohub.Repos(jconfig.User)
	if err != nil {
		log.Fatal(err)
	}
	err = gohub.CloneAll(jconfig.Home.Expand(), repos)
	if err != nil {
		log.Fatal(err)
	}
}

func status() {
	glob := filepath.Join(jconfig.Home.Expand(), "*")
	dirs, err := filepath.Glob(glob)
	if err != nil {
		log.Fatal(err)
	}
	if len(dirs) == 0 {
		log.Printf("no repos in %s", jconfig.Home)
	}
	clean := true
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
		switch len(buf.Bytes()) {
		case 0:
		default:
			clean = false
			console.Println(filepath.Base(d), "*")
		}
	}
	if clean {
		console.Println(jconfig.Home, "clean")
	}
	console.Flush()
}
