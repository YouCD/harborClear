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
			if len(r.Artifacts) > 0 { // 多个tag 要处理
				// 只有 config.RepoNamePrefix 的镜像会进行删除
				var deleteFlag bool
				for _, prefix := range config.RepoNamePrefix {
					if strings.HasPrefix(r.Name, prefix) {
						deleteFlag = true
						break
					}
				}
				if deleteFlag {
					log.Warnf("%-10s Project: %-8s RepoName: %-50s ", "Match", i, r.Name)
					loopDelete(i, r)
				} else {
					log.Infof("%-10s Project: %-8s RepoName: %-50s ", "KeepSave", i, r.Name)
				}
			}
		}
	}
}

func loopDelete(project string, r *Repo) {
	var ArtifactList []Artifact
	var tags []string
	for _, image := range r.Artifacts {
		if config.ClearFlag {
			for _, prefix := range config.DelTagPrefix {
				if strings.HasPrefix(image.Tag, prefix) {
					log.Warnf("%-10s Project: %-8s RepoName: %-50s Tag: %s", "Delete", project, r.Name, image.Tag)
					ArtifactList = append(ArtifactList, image)
					tags = append(tags, image.Tag)
					err := h.DeleteArtifact(project, r.Name, image.Digest)
					if err != nil {
						log.Error(err)
					}
				}
			}

		} else {
			for _, prefix := range config.DelTagPrefix {
				if strings.HasPrefix(image.Tag, prefix) {
					log.Warnf("%-10s Project: %-8s RepoName: %-50s Tag: %s", "Delete.dryrun", project, r.Name, image.Tag)
					ArtifactList = append(ArtifactList, image)
					tags = append(tags, image.Tag)
				}
			}
		}

	}
	if len(ArtifactList) > 0 {
		log.Infof("%-10s Project: %-8s RepoName: %-50s \033[31mTotal:\033[0m %d  \u001B[31mtags\u001B[0m: %s", "Delete", project, r.Name, len(ArtifactList), strings.Join(tags, ","))
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
	log.Infof("%-10s Project: %-8s RepoName: %-50s", "Scan", project, repoName)
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
			if len(v.Tags) > 0 {
				var artifact = Artifact{
					Digest:   v.Digest,
					Tag:      v.Tags[0].Name,
					PushTime: v.PushTime,
				}
				a = append(a, artifact)
			} else {
				log.Debugf("%-10s Project: %-8s RepoName: %-50s is null", "GetArtifacts", project, repoName)
			}

		}

	}

	return
}
