package zadig

import "net/http"

// 效能洞察
type EfficiencyService struct {
	client *Client
}

type ListDataOverviewResponse struct {
	ProjectCount  int64 `json:"project_count,omitempty"`  // 项目数量
	ClusterCount  int64 `json:"cluster_count,omitempty"`  // 集群数量
	ServiceCount  int64 `json:"service_count,omitempty"`  // 服务数量
	WorkflowCount int64 `json:"workflow_count,omitempty"` // 工作流数量
	EnvCount      int64 `json:"env_count,omitempty"`      // 环境数量
	ArtifactCount int64 `json:"artifact_count,omitempty"` // 交付物数量
}

// 数据概览
func (e *EfficiencyService) ListDataOverview(options ...RequestOptionFunc) (*ListDataOverviewResponse, *Response, error) {
	path := "/openapi/statistics/overview"
	req, err := e.client.NewRequest(http.MethodPost, path, nil, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListDataOverviewResponse)
	resp, err := e.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

type ListBuildDataStatistics struct {
	StartDate int64
	EndDate   int64
}

type ListBuildDataStatisticsResponse struct {
	Total   int64  `json:"total,omitempty"`
	Success int64  `json:"success,omitempty"`
	Data    []Data `json:"data,omitempty"`
}

type Data struct {
	Date    string `json:"date,omitempty"`
	Success int64  `json:"success,omitempty"`
	Failure int64  `json:"failure,omitempty"`
	Total   int64  `json:"total,omitempty"`
}

// 构建数据统计
func (e *EfficiencyService) ListBuildDataStatistics(opt ListBuildDataStatistics, options ...RequestOptionFunc) (*ListBuildDataStatisticsResponse, *Response, error) {
	path := "/openapi/statistics/build"
	req, err := e.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListBuildDataStatisticsResponse)
	resp, err := e.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 部署数据统计
type ListDeployDataStatistics struct {
	StartDate int64
	EndDate   int64
}

type ListDeployDataStatisticsResponse struct {
	Total   int64  `json:"total,omitempty"`
	Success int64  `json:"success,omitempty"`
	Data    []Data `json:"data,omitempty"`
}

func (e *EfficiencyService) ListDeployDataStatistics(opt ListDeployDataStatistics, options ...RequestOptionFunc) (*ListDeployDataStatisticsResponse, *Response, error) {
	path := "/openapi/statistics/deploy"
	req, err := e.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListDeployDataStatisticsResponse)
	resp, err := e.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 测试数据统计
type ListTestDataStatistics struct {
	StartDate int64
	EndDate   int64
}

type ListTestDataStatisticsResponse struct {
	CaseCount      int64      `json:"case_count,omitempty"`
	ExecCount      int64      `json:"exec_count,omitempty"`
	SuccessCount   int64      `json:"success_count,omitempty"`
	AverageRuntime int64      `json:"average_runtime,omitempty"`
	Data           []TestData `json:"data,omitempty"`
}

type TestData struct {
	Date         string `json:"date,omitempty"`
	SuccessCount int64  `json:"success_count,omitempty"`
	TimeoutCount int64  `json:"timeout_count,omitempty"`
	FailedCount  int64  `json:"failed_count,omitempty"`
}

func (e *EfficiencyService) ListTestDataStatistics(opt ListTestDataStatistics, options ...RequestOptionFunc) (*ListTestDataStatisticsResponse, *Response, error) {
	path := "/openapi/statistics/test"
	req, err := e.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	data := new(ListTestDataStatisticsResponse)
	resp, err := e.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
