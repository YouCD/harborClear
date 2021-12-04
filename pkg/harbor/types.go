package harbor

import "time"

type Statistics struct {
	PrivateProjectCount int `json:"private_project_count"`
	PrivateRepoCount    int `json:"private_repo_count"`
	PublicProjectCount  int `json:"public_project_count"`
	PublicRepoCount     int `json:"public_repo_count"`
	TotalProjectCount   int `json:"total_project_count"`
	TotalRepoCount      int `json:"total_repo_count"`
}

type Projects struct {
	UpdateTime         time.Time `json:"update_time"`
	OwnerName          string    `json:"owner_name"`
	Name               string    `json:"name"`
	Deleted            bool      `json:"deleted"`
	OwnerID            int       `json:"owner_id"`
	RepoCount          int       `json:"repo_count"`
	ChartCount         int       `json:"chart_count"`
	CreationTime       time.Time `json:"creation_time"`
	Togglable          bool      `json:"togglable"`
	CurrentUserRoleID  int       `json:"current_user_role_id"`
	CurrentUserRoleIds []int     `json:"current_user_role_ids"`
	CveAllowlist       struct {
		Items []struct {
			CveID string `json:"cve_id"`
		} `json:"items"`
		ProjectID    int       `json:"project_id"`
		ID           int       `json:"id"`
		ExpiresAt    int       `json:"expires_at"`
		UpdateTime   time.Time `json:"update_time"`
		CreationTime time.Time `json:"creation_time"`
	} `json:"cve_allowlist"`
	ProjectID  int `json:"project_id"`
	RegistryID int `json:"registry_id"`
	Metadata   struct {
		EnableContentTrust   string `json:"enable_content_trust"`
		AutoScan             string `json:"auto_scan"`
		Severity             string `json:"severity"`
		Public               string `json:"public"`
		ReuseSysCveAllowlist string `json:"reuse_sys_cve_allowlist"`
		PreventVul           string `json:"prevent_vul"`
		RetentionID          string `json:"retention_id"`
	} `json:"metadata"`
}

type RepoInfo struct {
	CreationTime       time.Time `json:"creation_time"`
	CurrentUserRoleID  int       `json:"current_user_role_id"`
	CurrentUserRoleIds []int     `json:"current_user_role_ids"`
	CveAllowlist       struct {
		CreationTime time.Time     `json:"creation_time"`
		ID           int           `json:"id"`
		Items        []interface{} `json:"items"`
		ProjectID    int           `json:"project_id"`
		UpdateTime   time.Time     `json:"update_time"`
	} `json:"cve_allowlist"`
	Metadata struct {
		Public string `json:"public"`
	} `json:"metadata"`
	Name       string    `json:"name"`
	OwnerID    int       `json:"owner_id"`
	OwnerName  string    `json:"owner_name"`
	ProjectID  int       `json:"project_id"`
	RepoCount  int       `json:"repo_count"`
	UpdateTime time.Time `json:"update_time"`
}

type Repo struct {
	ArtifactCount int       `json:"artifact_count"`
	CreationTime  time.Time `json:"creation_time"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int       `json:"project_id"`
	PullCount     int       `json:"pull_count,omitempty"`
	UpdateTime    time.Time `json:"update_time"`
}

type Artifact struct {
	AdditionLinks struct {
		BuildHistory struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"build_history"`
	} `json:"addition_links"`
	Digest     string `json:"digest"`
	ExtraAttrs struct {
		Architecture string    `json:"architecture"`
		Author       string    `json:"author"`
		Created      time.Time `json:"created"`
		Os           string    `json:"os"`
	} `json:"extra_attrs"`
	Icon              string      `json:"icon"`
	ID                int         `json:"id"`
	Labels            interface{} `json:"labels"`
	ManifestMediaType string      `json:"manifest_media_type"`
	MediaType         string      `json:"media_type"`
	ProjectID         int         `json:"project_id"`
	PullTime          time.Time   `json:"pull_time"`
	PushTime          time.Time   `json:"push_time"`
	References        interface{} `json:"references"`
	RepositoryID      int         `json:"repository_id"`
	Size              int         `json:"size"`
	Tags              []struct {
		ArtifactID   int       `json:"artifact_id"`
		ID           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		Name         string    `json:"name"`
		PullTime     time.Time `json:"pull_time"`
		PushTime     time.Time `json:"push_time"`
		RepositoryID int       `json:"repository_id"`
		Signed       bool      `json:"signed"`
	} `json:"tags"`
	Type string `json:"type"`
}

type DeleteArtifactErr struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
