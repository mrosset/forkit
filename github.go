package gohub

import (
	"encoding/json"
	"fmt"
	"github.com/str1ngs/util/file"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

const (
	ApiRepos = "users/%s/repos"
)

var (
	client = new(http.Client)
	api    = "https://api.github.com/%s"
)

func init() {
	log.SetPrefix("jflect: ")
	log.SetFlags(log.Lshortfile)
}

func CallApi(v interface{}, rest string) (err error) {
	url := fmt.Sprintf(api, rest)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %v %s", api, res.StatusCode,
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
			cmd := exec.Command("git", "clone", r.GitUrl, path.Join(dir, r.Name))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
		}
	}
	return
}
