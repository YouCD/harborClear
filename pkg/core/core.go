package core

import (
	config "harborClear/configs"
	"harborClear/pkg/harbor"
	"harborClear/pkg/log"
	"harborClear/pkg/tools"
	"strings"
	"sync"
	"time"
)

var (
	h         = harbor.GetHarbor(config.HarborUrl, config.HarborUser, config.HarborUserPassword)
	ImageList = make(map[string][]*Repo)

	repoList []*Repo
)

type Artifact struct {
	Digest   string
	Tag      string
	PushTime time.Time
}

type Repo struct {
	Name      string
	Artifacts []Artifact
}

func Core() {

	for _, project := range config.Projects {

		ProjectInfo, err := h.GetProjectByName(project)
		if err != nil {
			log.Error(err)
		}
		pageSize := tools.GetPageSize(ProjectInfo.RepoCount)
		getRepos(project, pageSize)
		ImageList[project] = repoList
	}

	deleteHandler()

}

func deleteHandler() {
	for i, v := range ImageList {
		for _, r := range v {
			switch {
			case config.KeepSave:
				if len(r.Artifacts) == 1 {
					log.Infof("%-8s Project: %-10s RepoName: %-50s Tag: %s", "KeepSave", i, r.Name, r.Artifacts[0].Tag)
				} else if len(r.Artifacts) > 1 {
					loopDelete(i, r)
				}

			case !config.KeepSave:
				loopDelete(i, r)
			}
		}
	}
}

func loopDelete(project string, r *Repo) {
	for _, image := range r.Artifacts {
		nowTime := time.Now()
		startTime := nowTime.AddDate(0, -config.Month, 0)
		if image.PushTime.Unix() < startTime.Unix() {
			if config.ClearFlag {
				err := h.DeleteArtifact(project, r.Name, image.Digest)
				if err != nil {
					log.Error(err)
				}
			}
			log.Warnf("%-8s Project: %-10s RepoName: %-50s Tag: %s", "Delete", project, r.Name, image.Tag)
		}
	}

}

func getRepos(project string, pageSize int) {
	for i := 1; i <= pageSize; i++ {

		repos, err := h.GetReposByProject(project, i)
		if err != nil {
			log.Error(err)
		}
		wg := &sync.WaitGroup{}

		wg.Add(len(repos))
		for _, repo := range repos {
			repoName := strings.Split(repo.Name, "/")[1]
			go getRepo(project, repoName, repo.ArtifactCount, wg)

		}
		wg.Wait()
	}
	return
}

func getRepo(project, repoName string, ArtifactCount int, wg *sync.WaitGroup) {
	var tempRepo Repo
	pSize := tools.GetPageSize(ArtifactCount)
	log.Infof("%-8s Project: %-10s RepoName: %-50s", "Scan", project, repoName)
	tempRepo.Name = repoName
	tempRepo.Artifacts = getArtifacts(pSize, project, repoName)
	//RepoChan <- &tempRepo
	repoList = append(repoList, &tempRepo)
	wg.Done()
	return
}

func getArtifacts(pageSize int, project, repoName string) (a []Artifact) {
	for i := 1; i <= pageSize; i++ {
		Artifacts, err := h.GetArtifacts(project, repoName, i)
		if err != nil {
			log.Error(err)
		}
		for _, v := range Artifacts {

			var artifact = Artifact{
				Digest:   v.Digest,
				Tag:      v.Tags[0].Name,
				PushTime: time.Time{},
			}
			a = append(a, artifact)
		}

	}

	return
}
