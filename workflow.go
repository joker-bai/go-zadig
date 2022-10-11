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

type Workflow struct {
	WorkflowName  string       `json:"workflow_name,omitempty"`
	EnvName       string       `json:"env_name,omitempty"`
	ReleaseImages []Image      `json:"release_images,omitempty"`
	Targets       []TargetArgs `json:"targets,omitempty"`
	Callback      Callback     `json:"callback,omitempty"`
}

// 执行工作流
type CreateWorkflowTaskOptions struct {
	WorkflowName  string       `json:"workflow_name,omitempty"`
	EnvName       string       `json:"env_name,omitempty"`
	ReleaseImages []Image      `json:"release_images,omitempty"`
	Targets       []TargetArgs `json:"targets,omitempty"`
	Callback      Callback     `json:"callback,omitempty"`
}

type CreateWorkflowTaskResponse struct {
	ProjectName  string `json:"project_name,omitempty"`
	WorkflowName string `json:"workflow_name,omitempty"`
	TaskID       int64  `json:"task_id,omitempty"`
}

func (w *WorkflowService) CreateWorkflowTask(opt *CreateWorkflowTaskOptions, options ...RequestOptionFunc) (*CreateWorkflowTaskResponse, *Response, error) {
	path := "/directory/workflowTask/create"
	req, err := w.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	task := new(CreateWorkflowTaskResponse)
	resp, err := w.client.Do(req, &task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, err
}

// 获取工作流任务状态
type GetWorkflowTaskOptions struct {
	CommitId string `url:"commitId,omitempty" json:"commitId,omitempty"`
}

type WorkflowTask struct {
	TaskID     int64  `json:"task_id,omitempty"`
	Status     string `json:"status,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	StartTime  int64  `json:"start_time,omitempty"`
	EndTime    int64  `json:"end_time,omitempty"`
	Url        string `json:"url,omitempty"`
}

func (w *WorkflowService) GetWorkflowTask(opt *GetWorkflowTaskOptions, options ...RequestOptionFunc) ([]*WorkflowTask, *Response, error) {
	path := "/directory/workflowTask"
	req, err := w.client.NewRequest(http.MethodGet, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var tasks []*WorkflowTask
	resp, err := w.client.Do(req, &tasks)
	if err != nil {
		return nil, resp, err
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
	return tasks, resp, err
}

// 取消工作流
type CanalWorkflowTaskOptions struct {
	ID           int64  `json:"id,omitempty"`
	PipelineName string `json:"pipelineName,omitempty"`
}

func (w *WorkflowService) CanalWorkflowTask(opt *CanalWorkflowTaskOptions, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("/directory/workflowTask/id/%d/pipelines/%s/cancel", opt.ID, opt.PipelineName)

	req, err := w.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return w.client.Do(req, nil)
}

// 工作流任务重试
type RestartWorkflowTaskOptions struct {
	ID           int64  `json:"id,omitempty"`
	PipelineName string `json:"pipelineName,omitempty"`
}

func (w *WorkflowService) RestartWorkflowTask(opt *RestartWorkflowTaskOptions, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("/directory/workflowTask/id/%d/pipelines/%s/restart", opt.ID, opt.PipelineName)

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

func (w *WorkflowService) GetWorkflowTaskDetail(opt *GetWorkflowTaskDetailOptions, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("/directory/workflowTask/id/%d/pipelines/%s", opt.ID, opt.PipelineName)

	req, err := w.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	return w.client.Do(req, nil)
}
