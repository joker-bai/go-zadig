package zadig

import (
	"fmt"
	"net/http"
)

// 接口: 项目列表

type ProjectService struct {
	client *Client
}

type ProjectListResponse []struct {
	Name       string   `json:"name"`
	Envs       []string `json:"envs"`
	Alias      string   `json:"alias"`
	Desc       string   `json:"desc"`
	UpdatedAt  int      `json:"updatedAt"`
	UpdatedBy  string   `json:"updatedBy"`
	Onboard    bool     `json:"onboard"`
	Public     bool     `json:"public"`
	DeployType string   `json:"deployType"`
}

// 请求数据
type GetProjectListOptions struct {
	Verbosity string `json:"Verbosity,omitempty"`
}

// GetProjectList  https://xxx.com/api/v1/picket/projects?verbosity=detailed
func (p *ProjectService) GetProjectList(opt *GetProjectListOptions, options ...RequestOptionFunc) (*ProjectListResponse, *Response, error) {

	//path := "/api/aslan/workflow/workflow?projectName=" + opt.PorjectName
	path := fmt.Sprintf("api/v1/picket/projects?verbosity=%s", opt.Verbosity)

	req, err := p.client.NewRequest(http.MethodGet, path, nil, options)

	if err != nil {
		return nil, nil, err
	}

	projectList := new(ProjectListResponse)
	resp, err := p.client.Do(req, &projectList)
	if err != nil {
		return nil, resp, err
	}

	return projectList, resp, err
}
