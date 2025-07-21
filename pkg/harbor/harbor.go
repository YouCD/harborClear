package harbor

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"harborClear/pkg/log"
	"harborClear/pkg/tools"
	"io"
	"net/http"
)

type Harbor struct {
	harborUrl string
	user      string
	password  string
}

func (m *Harbor) getRestClient(method, URL string, body io.Reader) (resp *http.Response, err error) {
	// 创建一个自定义的Transport，并禁用证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: tr}
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(m.user, m.password)
	return client.Do(req)
}

func GetHarbor(harborUrl, user, password string) *Harbor {
	log.Debugf("harborUrl: %s user: %s Password: %s", harborUrl, user, password)
	return &Harbor{
		harborUrl: harborUrl,
		user:      user,
		password:  password,
	}
}

func (m *Harbor) GetStatistics() (statistics *Statistics, err error) {
	statistics = new(Statistics)
	URL := fmt.Sprintf("%s/api/v2.0/statistics", m.harborUrl)
	log.Debugf("API URL: %s", URL)
	resp, err := m.getRestClient(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &statistics)
	if err != nil {
		return nil, err
	}
	return

}

func (m *Harbor) GetProjects(page int) (p []Projects, err error) {
	URL := fmt.Sprintf("%s/api/v2.0/projects?page=%d&page_size=100", m.harborUrl, page)
	log.Debugf("API URL: %s", URL)

	resp, err := m.getRestClient(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &p)
	if err != nil {
		return nil, err
	}
	return
}

func (m *Harbor) GetProjectByName(projectName string) (p *RepoInfo, err error) {

	statistics, err := m.GetStatistics()
	if err != nil {
		return nil, err
	}
	pageSize := tools.GetPageSize(statistics.TotalProjectCount)

	for i := 1; i <= pageSize; i++ {
		Projects, err := m.GetProjects(i)
		if err != nil {
			return nil, err
		}

		for _, Project := range Projects {
			if Project.Name == projectName {
				return m.GetProjectInfoByID(Project.ProjectID)
			}
		}
	}
	return
}

func (m *Harbor) GetProjectInfoByID(projectID int) (r *RepoInfo, err error) {
	URL := fmt.Sprintf("%s/api/v2.0/projects/%d", m.harborUrl, projectID)
	log.Debugf("API URL: %s", URL)

	resp, err := m.getRestClient(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &r)
	if err != nil {
		return nil, err
	}
	return
}

func (m *Harbor) GetReposByProject(projectName string, page int) (r []Repo, err error) {
	URL := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories?page_size=100&page=%d", m.harborUrl, projectName, page)
	log.Debugf("API URL: %s", URL)

	resp, err := m.getRestClient(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(content))
	err = json.Unmarshal(content, &r)
	if err != nil {
		return nil, err
	}
	return
}

func (m *Harbor) GetArtifacts(projectName, repoName string, page int) (a []Artifact, err error) {
	URL := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories/%s/artifacts?with_tag=true&with_scan_overview=true&with_label=true&page_size=100&page=%d", m.harborUrl, projectName, repoName, page)
	log.Debugf("API URL: %s", URL)

	resp, err := m.getRestClient(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &a)
	if err != nil {
		return nil, err
	}
	return
}

func (m *Harbor) DeleteArtifact(projectName, repoName, digest string) (err error) {
	URL := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories/%s/artifacts/%s", m.harborUrl, projectName, repoName, digest)
	log.Debugf("API URL: %s", URL)

	_, err = m.getRestClient(http.MethodDelete, URL, nil)
	if err != nil {
		return err
	}
	return
}
