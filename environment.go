package zadig

import (
	"fmt"
	"net/http"
)

// 根据项目名称获取项目环境信息

type EnvironmentService struct {
	client *Client
}

type EnvironmentResponse []struct {
	ProjectName     string        `json:"projectName"`
	Status          string        `json:"status"`
	Error           string        `json:"error"`
	Name            string        `json:"name"`
	UpdateBy        string        `json:"updateBy"`
	UpdateTime      int           `json:"updateTime"`
	IsPublic        bool          `json:"isPublic"`
	ClusterName     string        `json:"clusterName"`
	ClusterID       string        `json:"cluster_id"`
	Production      bool          `json:"production"`
	Source          string        `json:"source"`
	RegistryID      string        `json:"registry_id"`
	BaseRefs        []interface{} `json:"base_refs"`
	BaseName        string        `json:"base_name"`
	IsExisted       bool          `json:"is_existed"`
	ShareEnvEnable  bool          `json:"share_env_enable"`
	ShareEnvIsBase  bool          `json:"share_env_is_base"`
	ShareEnvBaseEnv string        `json:"share_env_base_env"`
}

// 请求数据

type GetEvnByProjectNameOptions struct {
	PorjectName string `json:"PorjectName,omitempty"`
}

//http://xxx.com/api/aslan/environment/environments?projectName=java-demo

func (p *EnvironmentService) GetEvnByProjectName(opt *GetEvnByProjectNameOptions, options ...RequestOptionFunc) (*EnvironmentResponse, *Response, error) {

	path := fmt.Sprintf("api/aslan/environment/environments?projectName=%s", opt.PorjectName)

	req, err := p.client.NewRequest(http.MethodGet, path, nil, options)

	if err != nil {
		return nil, nil, err
	}

	projectList := new(EnvironmentResponse)
	resp, err := p.client.Do(req, &projectList)
	if err != nil {
		return nil, resp, err
	}

	return projectList, resp, err
}
