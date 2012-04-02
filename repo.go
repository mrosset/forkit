package gohub

type Repo struct {
	Description  string                 `json:"description"`
	Name         string                 `json:"name"`
	Url          string                 `json:"url"`
	Fork         bool                   `json:"fork"`
	GitUrl       string                 `json:"git_url"`
	HtmlUrl      string                 `json:"html_url"`
	Private      bool                   `json:"private"`
	UpdatedAt    string                 `json:"updated_at"`
	Owner        map[string]interface{} `json:"owner"`
	OpenIssues   int                    `json:"open_issues"`
	Homepage     string                 `json:"homepage"`
	SvnUrl       string                 `json:"svn_url"`
	SshUrl       string                 `json:"ssh_url"`
	Size         int                    `json:"size"`
	CloneUrl     string                 `json:"clone_url"`
	Forks        int                    `json:"forks"`
	PushedAt     string                 `json:"pushed_at"`
	CreatedAt    string                 `json:"created_at"`
	Language     string                 `json:"language"`
	HasWiki      bool                   `json:"has_wiki"`
	HasIssues    bool                   `json:"has_issues"`
	Watchers     int                    `json:"watchers"`
	Id           int                    `json:"id"`
	HasDownloads bool                   `json:"has_downloads"`
}
