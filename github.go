package gohub

import (
	"encoding/json"
	"fmt"
	"github.com/str1ngs/util/file"
	"log"
	"net/http"
	"os/exec"
	"path"
)

const (
	ApiRepos = "users/%s/repos"
	Api      = "https://api.github.com/%s"
)

var (
	client = new(http.Client)
)

func init() {
	log.SetPrefix("gohub: ")
	log.SetFlags(log.Lshortfile)
}

func CallApi(v interface{}, rest string) (err error) {
	url := fmt.Sprintf(Api, rest)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %v %s", url, res.StatusCode,
			http.StatusText(res.StatusCode))
	}
	err = json.NewDecoder(res.Body).Decode(v)
	switch err.(type) {
	case *json.UnmarshalTypeError:
	default:
		return err
	}
	return nil

}

func Repos(user string) (repos []Repo, err error) {
	rest := fmt.Sprintf(ApiRepos, user)
	err = CallApi(&repos, rest)
	if err != nil {
		return nil, err
	}
	return
}

func CloneAll(dir string, repos []Repo) (err error) {
	for _, r := range repos {
		if !file.Exists(path.Join(dir, r.Name)) {
			fmt.Printf("cloneing %-20.20s %s\n", r.Name, r.SshUrl)
			cmd := exec.Command("git", "clone", r.SshUrl, path.Join(dir, r.Name))
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
	}
	return
}
