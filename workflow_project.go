package zadig

import (
	"fmt"
	"net/http"
)

// 接口: 根据projectname 获取工作流列表

type WorkflowProjectService struct {
	client *Client
}

// response数据
type WorkflowProjectResponse []struct {
	Name                 string               `json:"name"`
	ProjectName          string               `json:"projectName"`
	UpdateTime           int                  `json:"updateTime"`
	CreateTime           int                  `json:"createTime"`
	UpdateBy             string               `json:"updateBy"`
	Schedules            Schedules            `json:"schedules,omitempty"`
	SchedulerEnabled     bool                 `json:"schedulerEnabled"`
	EnabledStages        []string             `json:"enabledStages"`
	IsFavorite           bool                 `json:"isFavorite"`
	WorkflowType         string               `json:"workflow_type"`
	RecentTask           RecentTask           `json:"recentTask"`
	RecentSuccessfulTask RecentSuccessfulTask `json:"recentSuccessfulTask"`
	RecentFailedTask     RecentFailedTask     `json:"recentFailedTask"`
	AverageExecutionTime float64              `json:"averageExecutionTime"`
	SuccessRate          float64              `json:"successRate"`
	BaseName             string               `json:"base_name"`
	BaseRefs             []interface{}        `json:"base_refs"`
}
type Schedules struct {
	Enabled bool          `json:"enabled"`
	Items   []interface{} `json:"items"`
}
type RecentTask struct {
	TaskID       int    `json:"taskID"`
	PipelineName string `json:"pipelineName"`
	Status       string `json:"status"`
}
type RecentSuccessfulTask struct {
	TaskID       int    `json:"taskID"`
	PipelineName string `json:"pipelineName"`
	Status       string `json:"status"`
}
type RecentFailedTask struct {
	TaskID       int    `json:"taskID"`
	PipelineName string `json:"pipelineName"`
	Status       string `json:"status"`
}

// 请求数据
type GetWorkflowProjectNameOptions struct {
	PorjectName string `json:"PorjectName,omitempty"`
}

// GetWorkflowByPorectName https://leyancd.nancalcloud.com/api/aslan/workflow/workflow?projectName=leyan-devops
func (w *WorkflowProjectService) GetWorkflowByPorectName(opt *GetWorkflowProjectNameOptions, options ...RequestOptionFunc) (*WorkflowProjectResponse, *Response, error) {

	//path := "/api/aslan/workflow/workflow?projectName=" + opt.PorjectName
	path := fmt.Sprintf("api/aslan/workflow/workflow?projectName=%s", opt.PorjectName)

	req, err := w.client.NewRequest(http.MethodGet, path, nil, options)

	if err != nil {
		return nil, nil, err
	}

	workflow := new(WorkflowProjectResponse)
	resp, err := w.client.Do(req, &workflow)
	if err != nil {
		return nil, resp, err
	}

	return workflow, resp, err
}
