package zadig

import "net/http"

// 交付中心
type DeliveryCenterService struct {
	client *Client
}

// 获取交付中心版本列表

type VersionInfo struct {
	ID             string         `json:"id,omitempty"`
	Version        string         `json:"version,omitempty"`
	ProductName    string         `json:"productName,omitempty"`
	WorkflowName   string         `json:"workflowName,omitempty"`
	TaskID         int            `json:"task_id,omitempty"`
	Desc           string         `json:"desc,omitempty"`
	Labels         []string       `json:"labels,omitempty"`
	ProductEnvInfo ProductEnvInfo `json:"productEnvInfo,omitempty"`
	CreatedBy      string         `json:"created_by,omitempty"`
	CreatedAt      int64          `json:"created_at,omitempty"`
	DeletedAt      int64          `json:"deleted_at,omitempty"`
}

type ProductEnvInfo struct {
	ID           string        `json:"id,omitempty"`
	ProductName  string        `json:"product_name,omitempty"`
	CreateTime   int64         `json:"create_time,omitempty"`
	UpdateTime   int64         `json:"update_time,omitempty"`
	Namespace    string        `json:"namespace,omitempty"`
	Status       string        `json:"status,omitempty"`
	Revision     int64         `json:"revision,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	EnvName      string        `json:"env_name,omitempty"`
	UpdateBy     string        `json:"update_by,omitempty"`
	Auth         []interface{} `json:"auth,omitempty"`
	Visibility   string        `json:"visibility,omitempty"`
	Services     []Services    `json:"services,omitempty"`
	Render       Render        `json:"render,omitempty"`
	Error        string        `json:"error,omitempty"`
	Vars         []Vars        `json:"vars,omitempty"`
	IsPublic     bool          `json:"is_public,omitempty"`
	RoleIds      []int         `json:"role_ids,omitempty"`
	RecycleDay   int           `json:"recycle_day,omitempty"`
	Source       string        `json:"source,omitempty"`
	IsOpensource bool          `json:"is_opensource,omitempty"`
}

type Services struct {
	ServiceName string       `json:"service_name,omitempty"`
	ProductName string       `json:"product_name,omitempty"`
	Type        string       `json:"type,omitempty"`
	Revision    int64        `json:"revision,omitempty"`
	Containers  []Containers `json:"containers,omitempty"`
	Render      Render       `json:"render,omitempty"`
}

type Containers struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

type Render struct {
	Name        string `json:"name,omitempty"`
	Revision    int64  `json:"revision,omitempty"`
	ProductTmpl string `json:"product_tmpl,omitempty"`
	Description string `json:"description,omitempty"`
}

type Vars struct {
	Key      string   `json:"key,omitempty"`
	Value    string   `json:"value,omitempty"`
	Alias    string   `json:"alias,omitempty"`
	State    string   `json:"state,omitempty"`
	Services []string `json:"services,omitempty"`
}

type DeployInfo struct {
	ID              string        `json:"id,omitempty"`
	ReleaseID       string        `json:"releaseId,omitempty"`
	ServiceName     string        `json:"serviceName,omitempty"`
	ContainerName   string        `json:"containerName,omitempty"`
	Image           string        `json:"image,omitempty"`
	RegistryID      string        `json:"registry_id,omitempty"`
	YamlContents    []string      `json:"yamlContents,omitempty"`
	Envs            []interface{} `json:"envs,omitempty"`
	OrderedServices [][]string    `json:"orderedServices,omitempty"`
	StartTime       int64         `json:"start_time,omitempty"`
	EndTime         int64         `json:"end_time,omitempty"`
	CreatedAt       int64         `json:"created_at,omitempty"`
	DeletedAt       int64         `json:"deleted_at,omitempty"`
}

type ListDeliveryCenterVersions struct {
	ProjectName  string `url:"projectName" json:"projectName,omitempty"`   // 项目名称
	WorkflowName string `url:"workflowName" json:"workflowName,omitempty"` // 工作流名称
	TaskID       int    `url:"taskId" json:"taskId,omitempty"`             // 工作流任务ID
	ServiceName  string `url:"serviceName" json:"serviceName,omitempty"`   // 服务名称
}

type ListDeliveryCenterVersionsResponse struct {
	VersionInfo       VersionInfo   `json:"versionInfo,omitempty"`
	BuildInfo         []interface{} `json:"buildInfo,omitempty"`
	DeployInfo        DeployInfo    `json:"deployInfo,omitempty"`
	TestInfo          []interface{} `json:"testInfo,omitempty"`
	DistributeInfo    []interface{} `json:"distributeInfo,omitempty"`
	SecurityStatsInfo []interface{} `json:"securityStatsInfo,omitempty"`
}

func (d *DeliveryCenterService) ListDeliveryCenterVersions(opt ListDeliveryCenterVersions, options ...RequestOptionFunc) (*ListDeliveryCenterVersionsResponse, *Response, error) {
	path := "/api/aslan/delivery/releases"
	req, err := d.client.NewRequest(http.MethodGet, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListDeliveryCenterVersionsResponse)
	resp, err := d.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 获取交付物追踪信息

type Commits struct {
	Address       string `json:"address,omitempty"`
	Source        string `json:"source,omitempty"`
	RepoOwner     string `json:"repo_owner,omitempty"`
	RepoName      string `json:"repo_name,omitempty"`
	Branch        string `json:"branch,omitempty"`
	PR            int    `json:"pr,omitempty"`
	CommitId      string `json:"commit_id,omitempty"`
	CommitMessage string `json:"commit_message,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
}

type Activities struct {
	ArtifactId  string    `json:"artifact_id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Url         string    `json:"url,omitempty"`
	Commits     []Commits `json:"commits,omitempty"`
	StartTime   int64     `json:"start_time,omitempty"`
	EndTime     int64     `json:"end_time,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	CreatedTime int64     `json:"created_time,omitempty"`
}

type SortedActivities struct {
	Build []Activities `json:"build,omitempty"`
}

type ListDeliveryCenterArtifact struct {
	Image string `json:"image"`
}

type ListDeliveryCenterArtifactResponse struct {
	ID               string           `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Type             string           `json:"type,omitempty"`
	Source           string           `json:"source,omitempty"`
	Image            string           `json:"image,omitempty"`
	ImageTag         string           `json:"image_tag,omitempty"`
	ImageDigest      string           `json:"image_digest,omitempty"`
	ImageSize        int64            `json:"image_size,omitempty"`
	Architecture     string           `json:"architecture,omitempty"`
	Os               string           `json:"os,omitempty"`
	CreatedBy        string           `json:"created_by,omitempty"`
	CreatedTime      int64            `json:"created_time,omitempty"`
	Activities       []Activities     `json:"activities,omitempty"`
	SortedActivities SortedActivities `json:"sortedActivities,omitempty"`
}

func (d *DeliveryCenterService) ListDeliveryCenterArtifact(opt ListDeliveryCenterArtifact, options ...RequestOptionFunc) (*ListDeliveryCenterArtifactResponse, *Response, error) {
	path := "/api/directory/dc/artifact"
	req, err := d.client.NewRequest(http.MethodGet, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListDeliveryCenterArtifactResponse)
	resp, err := d.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
