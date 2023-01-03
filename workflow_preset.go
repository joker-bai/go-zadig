package zadig

import (
	"fmt"
	"net/http"
)

// 接口: 获取执行工作流的参数接口

type WorkflowPresetService struct {
	client *Client
}

// response数据
type WorkflowPresetResponse struct {
	WorkflowName        string    `json:"workflow_name"`
	ProductTmplName     string    `json:"product_tmpl_name"`
	Namespace           string    `json:"namespace"`
	Targets             []Targets `json:"targets"`
	ReqID               string    `json:"req_id"`
	DistributeEnabled   bool      `json:"distribute_enabled"`
	WorkflowTaskCreator string    `json:"workflow_task_creator"`
	IgnoreCache         bool      `json:"ignore_cache"`
	ResetCache          bool      `json:"reset_cache"`
	NotificationID      string    `json:"notification_id"`
	MergeRequestID      string    `json:"merge_request_id"`
	CommitID            string    `json:"commit_id"`
	Source              string    `json:"source"`
	RepoOwner           string    `json:"repo_owner"`
	RepoNamespace       string    `json:"repo_namespace"`
	RepoName            string    `json:"repo_name"`
	EnvName             string    `json:"env_name"`
}
type Repos2 struct {
	Source        string `json:"source"`
	RepoOwner     string `json:"repo_owner"`
	RepoNamespace string `json:"repo_namespace"`
	RepoName      string `json:"repo_name"`
	RemoteName    string `json:"remote_name"`
	Branch        string `json:"branch"`
	IsPrimary     bool   `json:"is_primary"`
	CodehostID    int    `json:"codehost_id"`
	OauthToken    string `json:"oauth_token"`
	Address       string `json:"address"`
	AuthType      string `json:"auth_type"`
}
type Build struct {
	Repos []Repos2 `json:"repos"`
}
type Deploy struct {
	Env  string `json:"env"`
	Type string `json:"type"`
}
type Targets struct {
	Name        string   `json:"name"`
	ImageName   string   `json:"image_name"`
	ServiceName string   `json:"service_name"`
	ProductName string   `json:"product_name"`
	Build       Build    `json:"build"`
	Deploy      []Deploy `json:"deploy"`
	BinFile     string   `json:"bin_file"`
	HasBuild    bool     `json:"has_build"`
	BuildName   string   `json:"build_name"`
}

// 请求数据
//1.env 默认是约定   dev / test
//2.workflow_name  传递
//3.project_name   项目名称

type GetWorkflowPresetNameOptions struct {
	Env          string `json:"Env,omitempty"`
	WorkflowName string `json:"WorkflowName,omitempty"`
	ProjectName  string `json:"PorjectName,omitempty"`
}

// GetWorkflowByPorectName https://xxx.com/api/aslan/workflow/workflowtask/preset/dev/show-demo?projectName=nancal-demo
func (w *WorkflowPresetService) PresetWorkflow(opt *GetWorkflowPresetNameOptions, options ...RequestOptionFunc) (*WorkflowPresetResponse, *Response, error) {

	//path := "/api/aslan/workflow/workflow?projectName=" + opt.PorjectName
	path := fmt.Sprintf("api/aslan/workflow/workflowtask/preset/%s/%s?projectName=%s", opt.Env, opt.WorkflowName, opt.ProjectName)

	req, err := w.client.NewRequest(http.MethodGet, path, nil, options)

	if err != nil {
		return nil, nil, err
	}

	workflow := new(WorkflowPresetResponse)
	resp, err := w.client.Do(req, &workflow)
	if err != nil {
		fmt.Println("请求或者转换数据出错")
		fmt.Println(err)
		return nil, resp, err
	}

	return workflow, resp, err
}
