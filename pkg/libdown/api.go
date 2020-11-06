package libdown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GithubApiUrl = "https://api.github.com"
)

func FetchRepos(user, token string, page, per_page int) ([]Repo, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/users/%s/repos?per_page=%d&page=%d", GithubApiUrl, user, per_page, page),
		nil,
	)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := ConsumeResponse(res)
	if err != nil {
		return nil, err
	}
	repos, err := ParseRepos(content)
	if err != nil {
		return nil, err
	}
	return repos, err
}

func FetchRelease(url, token string) (*Release, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := ConsumeResponse(response)
	if err != nil {
		return nil, err
	}
	release, err := ParseRelease(content)
	if err != nil {
		return nil, err
	}
	return &release, err
}

func ConsumeResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ParseRepos(payload []byte) (res []Repo, err error) {
	err = json.Unmarshal(payload, &res)
	return
}

func ParseRelease(payload []byte) (res Release, err error) {
	err = json.Unmarshal(payload, &res)
	return
}
