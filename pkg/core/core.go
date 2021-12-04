package core

import (
	config "harborClear/configs"
	"harborClear/pkg/harbor"
	"harborClear/pkg/log"
	"harborClear/pkg/tools"

	"strings"
	"time"
)

var (
	h = harbor.GetHarbor(config.HarborUrl, config.HarborUser, config.HarborUserPassword)
)

func Core() {
	for _, project := range config.Projects {

		ProjectInfo, err := h.GetProjectByName(project)
		if err != nil {
			log.Error(err)
		}
		p := tools.GetPageSize(ProjectInfo.RepoCount)
		for i := 1; i <= p; i++ {

			repos, err := h.GetReposByProject(project, i)
			if err != nil {
				log.Error(err)
			}

			for _, repo := range repos {
				repoName := strings.Split(repo.Name, "/")[1]
				p := tools.GetPageSize(repo.ArtifactCount)
				log.Infof("%-6s Project: %-10s RepoName: %-50s", "Scan", project, repoName)
				var ArtifactList []harbor.Artifact
				for i := 1; i <= p; i++ {
					Artifacts, err := h.GetArtifacts(project, repoName, i)
					if err != nil {
						log.Error(err)
					}

					for _, v := range Artifacts {
						nowTime := time.Now()
						startTime := nowTime.AddDate(0, -config.Month, 0)
						if v.PushTime.Unix() < startTime.Unix() {
							if config.ClearFlag {
								err= h.DeleteArtifact(project,repoName, v.Digest)
								if err != nil {
									log.Error(err)
								}
								log.Warnf("%-6s Project: %-10s RepoName: %-50s Tag: %s", "Delete", project, repoName, v.Tags[0].Name)
							}
							ArtifactList = append(ArtifactList, v)
						}
					}

				}
				if len(ArtifactList) > 0 {
					log.Infof("%-6s Project: %-10s RepoName: %-50s \033[31mTotal:\033[0m %d", "Delete", project, repoName, len(ArtifactList))
				}
			}

		}
	}

}
