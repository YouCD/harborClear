package harbor

import (
	"fmt"

	"testing"
	"time"
)

var (
	HarborUrl          = "https://some.harbor.com"
	HarborUser         = "admin"
	HarborUserPassword = "harbor12345"
	harbor             = GetHarbor(HarborUrl, HarborUser, HarborUserPassword)
)

func TestGetStatistics(t *testing.T) {
	Statistics, err := harbor.GetStatistics()
	if err != nil {
		t.Error(err)
	}

	t.Log(Statistics)
}
func TestGetProjects(t *testing.T) {
	p, err := harbor.GetProjects(1)
	if err != nil {
		t.Error(err)
	}
	for _, v := range p {
		fmt.Println(v.Name, v.RepoCount, v.ProjectID)
	}
}

func TestGetProjectInfoByID(t *testing.T) {
	p, err := harbor.GetProjectInfoByID(3)
	if err != nil {
		t.Error(err)
	}
	t.Log(p.RepoCount)
}

func TestGetReposByProject(t *testing.T) {
	p, err := harbor.GetReposByProject("test", 1)
	if err != nil {
		t.Error(err)
	}
	for _, v := range p {
		fmt.Println(v.Name, v.ArtifactCount, v)
	}
}

func TestGetArtifacts(t *testing.T) {
	p, err := harbor.GetArtifacts("test", "ms-tiktok-test", 1)
	if err != nil {
		t.Error(err)
	}
	nowTime := time.Now()

	startTime := nowTime.AddDate(0, -1, 0)
	fmt.Println(startTime)

	var Artifacts []Artifact
	for _, v := range p {
		if v.PushTime.Unix() < startTime.Unix() {
			fmt.Println(v.PushTime, v.Digest)
			Artifacts = append(Artifacts, v)
		}

	}
	fmt.Println(len(Artifacts))
}

func TestDeleteArtifact(t *testing.T) {
	err := harbor.DeleteArtifact("test", "ms-tiktok-test", "sha256:5776ead3a767c2be73b65fd8c99e01903e369c8207ad5f1a688e6e2816c5f74f")
	if err != nil {
		t.Error(err)
	}

}

func TestGetProjectByName(t *testing.T) {
	p, err := harbor.GetProjectByName("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)

}
