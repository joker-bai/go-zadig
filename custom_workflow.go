package zadig

import "net/http"

// 自定义工作流

type CustomWorkflowService struct {
	client *Client
}

// 获取自定义工作流任务详情

type ListCustomWorkflowTaskDetail struct {
	TaskId       int    `json:"task_id"`       // 自定义工作流任务ID
	WorkflowName string `json:"workflow_name"` // 自定义工作流名称
}

type Params struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Type         string `json:"type,omitempty"`
	Value        string `json:"value,omitempty"`
	Default      string `json:"default,omitempty"`
	IsCredential bool   `json:"is_credential,omitempty"`
}

type Repos struct {
	Source        string `json:"source,omitempty"`
	RepoOwner     string `json:"repo_owner,omitempty"`
	RepoNamespace string `json:"repo_namespace,omitempty"`
	RepoName      string `json:"repo_name,omitempty"`
	RemoteName    string `json:"remote_name,omitempty"`
	Branch        string `json:"branch,omitempty"`
	CommitId      string `json:"commit_id,omitempty"`
	CommitMessage string `json:"commit_message,omitempty"`
	Address       string `json:"address,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
}

type Envs struct {
	Key          string `json:"key,omitempty"`
	Value        string `json:"value,omitempty"`
	Type         string `json:"type,omitempty"`
	IsCredential bool   `json:"is_credential,omitempty"`
	Name         string `json:"name,omitempty"` // 自定义任务使用
}

type ServiceAddImage struct {
	ServiceName   string `json:"service_name,omitempty"`
	ServiceModule string `json:"service_module,omitempty"`
	Image         string `json:"image,omitempty"`
}

type Inputs struct {
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	Type         string   `json:"type,omitempty"`
	Value        string   `json:"value,omitempty"`
	ChoiceOption []string `json:"choice_option,omitempty"`
	Default      string   `json:"default,omitempty"`
	IsCredential bool     `json:"is_credential,omitempty"`
}

type Spec struct {
	Name               string            `json:"name,omitempty"`
	IsOffical          bool              `json:"is_offical,omitempty"`
	Description        string            `json:"description,omitempty"`
	RepoUrl            string            `json:"repo_url,omitempty"`
	Version            string            `json:"version,omitempty"`
	Image              string            `json:"image,omitempty"`
	Args               []interface{}     `json:"args,omitempty"`
	Cms                []interface{}     `json:"cms,omitempty"`
	Envs               []Envs            `json:"envs,omitempty"` // 构建变量信息
	Inputs             []Inputs          `json:"inputs,omitempty"`
	Outputs            []interface{}     `json:"outputs,omitempty"`
	ServiceName        string            `json:"service_name,omitempty"`   // 服务名称
	ServiceModule      string            `json:"service_module,omitempty"` // 服务组件名称
	Env                string            `json:"env,omitempty"`
	SkipCheckRunStatus bool              `json:"skip_check_run_status,omitempty"` // 是否关闭服务状态检测
	ServiceAddImage    []ServiceAddImage `json:"service_add_image,omitempty"`     //部署的服务、服务组件、镜像信息
	Repos              []Repos           `json:"repos,omitempty"`                 // 代码信息
}

type Jobs struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`       // 构建任务 freestyle：通用任务 zadig-deploy：内置部署任务 zadig-build：内置构建任务
	Status    string `json:"status,omitempty"`     // 任务状态
	StartTime int64  `json:"start_time,omitempty"` // 任务执行开始时间，Unix 时间戳格式
	EndTime   int64  `json:"end_time,omitempty"`   // # 任务执行结束时间，Unix 时间戳格式
	Error     string `json:"error,omitempty"`
	Spec      Spec   `json:"spec,omitempty"` // # 构建任务执行详细信息（包括代码信息、镜像信息、服务信息、服务组件信息、构建变量信息）
}

type ApproveUsers struct {
	UserName        string `json:"user_name,omitempty"`
	RejectOrApprove string `json:"reject_or_approve,omitempty"`
	Comment         string `json:"comment,omitempty"`
	OperationTime   int64  `json:"operation_time,omitempty"`
}

// Approval 审核信息
type Approval struct {
	Enabled         bool           `json:"enabled,omitempty"`           // 是否需要审核
	ApproveUsers    []ApproveUsers `json:"approve_users,omitempty"`     // 审核人列表
	Timeout         int            `json:"timeout,omitempty"`           // 审核超时时间，单位：分钟
	NeededApprovers int            `json:"needed_approvers,omitempty"`  // 需要满足的审核通过人数
	Description     string         `json:"description,omitempty"`       // 审核描述
	RejectOrApprove string         `json:"reject_or_approve,omitempty"` // approve：通过 reject：拒绝
}

type Stages struct {
	Name      string
	Status    string
	StartTime int64
	EndTime   int64
	Approval  Approval
	Jobs      []Jobs
}

type ListCustomWorkflowTaskDetailReponse struct {
	TaskId        int      `json:"task_id,omitempty"`
	Workflowoname string   `json:"workflowoname,omitempty"`
	Parms         []Params `json:"parms,omitempty"`
	Status        string   `json:"status,omitempty"`
	TaskCreator   string   `json:"task_creator,omitempty"`
	CreateTime    int64    `json:"create_time,omitempty"`
	StartTime     int64    `json:"start_time,omitempty"`
	EndTime       int64    `json:"end_time,omitempty"`
	Stages        []Stages `json:"stages,omitempty"`
	ProjectName   string   `json:"project_name,omitempty"`
}

func (c *CustomWorkflowService) ListCustomWorkflowTaskDetail(opt *ListCustomWorkflowTaskDetail, options ...RequestOptionFunc) (*ListCustomWorkflowTaskDetailReponse, *Response, error) {
	path := "/openapi/workflows/custom/task"
	req, err := c.client.NewRequest(http.MethodGet, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListCustomWorkflowTaskDetailReponse)
	resp, err := c.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 执行自定义工作流

type ServiceList struct {
	ServiceName   string     `json:"service_name,omitempty"`   // 待部署服务名称
	ServiceModule string     `json:"service_module,omitempty"` // 待部署服务组件名称
	ImageName     string     `json:"image_name,omitempty"`     // 待部署服务组件的镜像
	RepoInfo      []RepoInfo `json:"repo_info,omitempty"`
	Inputs        []KV       `json:"inputs,omitempty"`
}

type TargetList struct {
	WorkloadType  string `json:"workload_type,omitempty"`  // 待部署容器应用的类型，支持 Deployment 以及 StatefulSet
	WorkloadName  string `json:"workload_name,omitempty"`  // 待部署容器应用的名称
	ContainerName string `json:"container_name,omitempty"` // 待部署容器应用的 container 名称
	ImageName     string `json:"image_name,omitempty"`     // 待部署容器的镜像信息
}

type KV struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type RepoInfo struct {
	CodehostName  string `json:"codehost_name,omitempty"`  // 代码库标识
	RepoNamespace string `json:"repo_namespace,omitempty"` // 代码库组织名/用户名
	RepoName      string `json:"repo_name,omitempty"`      // 代码库名称
	Branch        string `json:"branch,omitempty"`         // 代码分支
}

type CreateCustomWorkflowTaskParameters struct {
	Register    string        `json:"register,omitempty"`
	EnvName     string        `json:"env_name,omitempty"` // 待部署环境信息，若设置为固定值或全局变量，则无需配置该字段
	ServiceList []ServiceList `json:"service_list,omitempty"`
	TargetList  []TargetList  `json:"target_list,omitempty"` // 待部署容器信息，若设置为固定值，则无需指定
	KV          []KV          `json:"kv,omitempty"`
	RepoInfo    []RepoInfo    `json:"repo_info,omitempty"`
}

type CreateCustomWorkflowTaskInput struct {
	JobName    string                             `json:"job_name,omitempty"` // 构建任务名称
	JobType    string                             `json:"job_type,omitempty"` // 任务类型 构建任务：zadig-build 部署任务：zadig-deploy K8s任务：custom-deploy 通用任务：freestyle
	Parameters CreateCustomWorkflowTaskParameters `json:"parameters,omitempty"`
}

type CreateCustomWorkflowTask struct {
	ProjectName  string                          `json:"project_name"`  // 项目名称
	WorkflowName string                          `json:"workflow_name"` // 自定义工作流名称
	Inputs       []CreateCustomWorkflowTaskInput `json:"inputs"`        // 执行工作流具体参数
}

type CreateCustomWorkflowTaskResponse struct {
	ProjectName  string `json:"project_name,omitempty"`
	WorkflowName string `json:"workflow_name,omitempty"`
	TaskID       int    `json:"task_id,omitempty"`
}

func (c *CustomWorkflowService) CreateCustomWorkflowTask(opt *CreateCustomWorkflowTask, options ...RequestOptionFunc) (*CreateCustomWorkflowTaskResponse, *Response, error) {
	path := "openapi/workflows/custom/task"
	req, err := c.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(CreateCustomWorkflowTaskResponse)
	resp, err := c.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 取消自定义工作流任务
type DeleteCustomWorkflowTask struct {
	TaskId       int    `json:"task_id"`
	WorkflowName string `json:"workflow_name"`
}

type DeleteCustomWorkflowTaskResponse struct {
	Code        int
	Description string
	Extra       interface{}
	Message     string
	Type        string
}

func (c *CustomWorkflowService) DeleteCustomWorkflowTask(opt *DeleteCustomWorkflowTask, options ...RequestOptionFunc) (*DeleteCustomWorkflowTaskResponse, *Response, error) {
	path := "openapi/workflows/custom/task"
	req, err := c.client.NewRequest(http.MethodDelete, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(DeleteCustomWorkflowTaskResponse)
	resp, err := c.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 自定义工作流审核
type AuditCustomWorkflowTask struct {
	TaskID       string `json:"task_id"`
	WorkflowName string `json:"workflow_name"`
	StageName    string `json:"stage_name"`
	Approve      bool   `json:"approve,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

type AuditCustomWorkflowTaskResponse struct {
	Code        int
	Description string
	Extra       interface{}
	Message     string
	Type        string
}

func (c *CustomWorkflowService) AuditCustomWorkflowTask(opt *AuditCustomWorkflowTask, options ...RequestOptionFunc) (*AuditCustomWorkflowTaskResponse, *Response, error) {
	path := "openapi/workflows/custom/task/approve"
	req, err := c.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(AuditCustomWorkflowTaskResponse)
	resp, err := c.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
