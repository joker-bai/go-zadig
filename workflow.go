package zadig

import (
	"fmt"
	"net/http"
)

type WorkflowService struct {
	client *Client
}

type Callback struct {
	CallbackUrl  string                 `json:"callback_url,omitempty"`
	CallbackVars map[string]interface{} `json:"callback_vars,omitempty"`
}

type Image struct {
	Image         string `json:"image,omitempty"`
	ServiceName   string `json:"service_name,omitempty"`
	ServiceModule string `json:"service_module,omitempty"`
	RegistryRepo  string `json:"registry_repo,omitempty"`
}

type BuildArgs struct {
	Repos []Repository `json:"repos,omitempty"`
}

type TargetArgs struct {
	Name        string    `json:"name,omitempty"`
	ServiceType string    `json:"service_type,omitempty"`
	Build       BuildArgs `json:"build,omitempty"`
}

type Repository struct {
	RepoName string `json:"repo_name,omitempty" `
	Branch   string `json:"branch,omitempty" `
	PR       int    `json:"pr,omitempty"`
}

type FunctionTestReport struct {
	Tests     int    `json:"tests,omitempty"`
	Successes int    `json:"successes,omitempty"`
	Failures  int    `json:"failures,omitempty"`
	Skips     int    `json:"skips,omitempty"`
	Errors    int    `json:"errors,omitempty"`
	DetailURL string `json:"detail_url,omitempty"`
}

type TestReport struct {
	TestName           string             `json:"test_name,omitempty"`
	FunctionTestReport FunctionTestReport `json:"function_test_report,omitempty"`
}

type Workflow struct {
	WorkflowName  string       `json:"workflow_name,omitempty"`
	EnvName       string       `json:"env_name,omitempty"`
	ReleaseImages []Image      `json:"release_images,omitempty"`
	Targets       []TargetArgs `json:"targets,omitempty"`
	Callback      Callback     `json:"callback,omitempty"`
}

// 执行工作流
type ExecWorkflowTaskOptions struct {
	WorkflowName string        `json:"workflow_name"`
	ProjectName  string        `json:"project_name"`
	Input        WorkflowInput `json:"input"`
}

type WorkflowInput struct {
	TargetEnv string         `json:"target_env,omitempty"`
	Build     ExecBuildArgs  `json:"build"`
	Deploy    ExecDeployArgs `json:"deploy"`
}

type ExecBuildArgs struct {
	Enabled     bool               `json:"enabled"`
	ServiceList []BuildServiceInfo `json:"service_list"`
}

type BuildServiceInfo struct {
	ServiceModule string           `json:"service_module"`
	ServiceName   string           `json:"service_name"`
	RepoInfo      []RepositoryInfo `json:"repo_info"`
	Inputs        []UserInput      `json:"inputs"`
}

type RepositoryInfo struct {
	CodehostName  string `json:"codehost_name"`
	RepoNamespace string `json:"repo_namespace"`
	RepoName      string `json:"repo_name"`
	Branch        string `json:"branch"`
	PR            int    `json:"pr"`
}

type UserInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ExecDeployArgs struct {
	Enabled     bool                `json:"enabled"`
	Source      string              `json:"source"`
	ServiceList []DeployServiceInfo `json:"service_list"`
}

type DeployServiceInfo struct {
	ServiceModule string `json:"service_module"`
	ServiceName   string `json:"service_name"`
	Image         string `json:"image"`
}

type ExecWorkflowTaskResponse struct {
	ProjectName  string `json:"project_name,omitempty"`
	WorkflowName string `json:"workflow_name,omitempty"`
	TaskID       int64  `json:"task_id,omitempty"`
}

func (w *WorkflowService) ExecWorkflowTask(opt *ExecWorkflowTaskOptions, options ...RequestOptionFunc) (*ExecWorkflowTaskResponse, *Response, error) {
	path := "/openapi/workflows/product/task"
	req, err := w.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	task := new(ExecWorkflowTaskResponse)
	resp, err := w.client.Do(req, &task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, err
}

// 获取工作流任务状态
type GetWorkflowTaskOptions struct {
	CommitId string `url:"commitId" json:"commitId"`
}

type WorkflowTask struct {
	TaskID     int64  `json:"task_id,omitempty"`
	Status     string `json:"status,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	StartTime  int64  `json:"start_time,omitempty"`
	EndTime    int64  `json:"end_time,omitempty"`
	Url        string `json:"url,omitempty"`
}

// GetWorkflowTaskStatus 获取工作流任务状态
func (w *WorkflowService) GetWorkflowTaskStatus(opt *GetWorkflowTaskOptions, options ...RequestOptionFunc) ([]*WorkflowTask, *Response, error) {
	path := "/api/directory/workflowTask"
	req, err := w.client.NewRequest(http.MethodGet, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var tasks []*WorkflowTask
	resp, err := w.client.Do(req, &tasks)
	if err != nil {
		return nil, resp, err
	}

	return tasks, resp, err
}

// 取消工作流
type CanalWorkflowTaskOptions struct {
	ID           int64  `json:"id,omitempty"`
	PipelineName string `json:"pipelineName,omitempty"`
}

// CanalWorkflowTask 取消工作流任务
func (w *WorkflowService) CanalWorkflowTask(opt *CanalWorkflowTaskOptions, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("/api/directory/workflowTask/id/%d/pipelines/%s/cancel", opt.ID, opt.PipelineName)

	req, err := w.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return w.client.Do(req, nil)
}

// 工作流任务重试
type RestartWorkflowTaskOptions struct {
	ID           int64  `json:"id"`
	PipelineName string `json:"pipelineName"`
}

// RestartWorkflowTask 重试工作流任务
func (w *WorkflowService) RestartWorkflowTask(opt *RestartWorkflowTaskOptions, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("/api/directory/workflowTask/id/%d/pipelines/%s/restart", opt.ID, opt.PipelineName)

	req, err := w.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return w.client.Do(req, nil)
}

// 获取工作流任务详情
type GetWorkflowTaskDetailOptions struct {
	ID           int64  `json:"id,omitempty"`
	PipelineName string `json:"pipelineName,omitempty"`
}

type GetWorkflowTaskDetailResponse struct {
	WorkflowName string
	EnvName      string
	Status       string
	Targets      []TargetArgs
	Images       []Image
	TestReports  []TestReport
}

// GetWorkflowTaskDetail 获取工作流任务详情
func (w *WorkflowService) GetWorkflowTaskDetail(opt *GetWorkflowTaskDetailOptions, options ...RequestOptionFunc) (*GetWorkflowTaskDetailResponse, *Response, error) {
	u := fmt.Sprintf("/api/directory/workflowTask/id/%d/pipelines/%s", opt.ID, opt.PipelineName)

	req, err := w.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	task := new(GetWorkflowTaskDetailResponse)
	resp, err := w.client.Do(req, &task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, err
}
