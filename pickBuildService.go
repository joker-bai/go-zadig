package zadig

import (
	"fmt"
	"net/http"
)

//pickbuildservice 获取执行工作流参数的接口

type pickBuildService struct {
	client *Client
}

type PickBuildResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Enabled         bool   `json:"enabled"`
	ProductTmplName string `json:"product_tmpl_name"`
	UpdateBy        string `json:"update_by"`
	CreateBy        string `json:"create_by"`
	UpdateTime      int    `json:"update_time"`
	CreateTime      int    `json:"create_time"`
	Schedules       struct {
		Enabled bool          `json:"enabled"`
		Items   []interface{} `json:"items"`
	} `json:"schedules"`
	ScheduleEnabled bool `json:"schedule_enabled"`
	BuildStage      struct {
		Enabled bool `json:"enabled"`
		Modules []struct {
			Target struct {
				ProductName   string        `json:"product_name"`
				ServiceName   string        `json:"service_name"`
				ServiceModule string        `json:"service_module"`
				BuildName     string        `json:"build_name"`
				Envs          []interface{} `json:"envs"`
			} `json:"target"`
			HideServiceModule bool          `json:"hide_service_module"`
			BuildModuleVer    string        `json:"build_module_ver"`
			BranchFilter      []interface{} `json:"branch_filter"`
		} `json:"modules"`
	} `json:"build_stage"`
	ArtifactStage struct {
		Enabled bool `json:"enabled"`
		Modules []struct {
			HideServiceModule bool `json:"hide_service_module"`
			Target            struct {
				ProductName   string        `json:"product_name"`
				ServiceName   string        `json:"service_name"`
				ServiceModule string        `json:"service_module"`
				BuildName     string        `json:"build_name"`
				Envs          []interface{} `json:"envs"`
			} `json:"target"`
		} `json:"modules"`
	} `json:"artifact_stage"`
	TestStage struct {
		Enabled bool `json:"enabled"`
	} `json:"test_stage"`
	SecurityStage   interface{} `json:"security_stage"`
	DistributeStage struct {
		Enabled     bool          `json:"enabled"`
		S3StorageID string        `json:"s3_storage_id"`
		ImageRepo   string        `json:"image_repo"`
		JumpBoxHost string        `json:"jump_box_host"`
		Distributes []interface{} `json:"distributes"`
		Releases    []interface{} `json:"releases"`
	} `json:"distribute_stage"`
	ExtensionStage struct {
		Enabled    bool          `json:"enabled"`
		URL        string        `json:"url"`
		Path       string        `json:"path"`
		IsCallback bool          `json:"is_callback"`
		Timeout    int           `json:"timeout"`
		Headers    []interface{} `json:"headers"`
	} `json:"extension_stage"`
	NotifyCtls []struct {
		Enabled       bool     `json:"enabled"`
		WebhookType   string   `json:"webhook_type"`
		WeChatWebHook string   `json:"weChat_webHook"`
		NotifyType    []string `json:"notify_type"`
	} `json:"notify_ctls"`
	HookCtl struct {
		Enabled bool `json:"enabled"`
		Items   []struct {
			AutoCancel          bool `json:"auto_cancel"`
			CheckPatchSetChange bool `json:"check_patch_set_change"`
			MainRepo            struct {
				Name          string   `json:"name"`
				Source        string   `json:"source"`
				RepoOwner     string   `json:"repo_owner"`
				RepoNamespace string   `json:"repo_namespace"`
				RepoName      string   `json:"repo_name"`
				Branch        string   `json:"branch"`
				Tag           string   `json:"tag"`
				Committer     string   `json:"committer"`
				MatchFolders  []string `json:"match_folders"`
				CodehostID    int      `json:"codehost_id"`
				Events        []string `json:"events"`
				Label         string   `json:"label"`
				Revision      string   `json:"revision"`
				IsRegular     bool     `json:"is_regular"`
			} `json:"main_repo"`
			WorkflowArgs struct {
				WorkflowName     string `json:"workflow_name"`
				ProductTmplName  string `json:"product_tmpl_name"`
				Namespace        string `json:"namespace"`
				EnvRecyclePolicy string `json:"env_recycle_policy"`
				EnvUpdatePolicy  string `json:"env_update_policy"`
				Targets          []struct {
					Name        string        `json:"name"`
					ImageName   string        `json:"image_name"`
					ServiceName string        `json:"service_name"`
					ProductName string        `json:"product_name"`
					Build       interface{}   `json:"build"`
					Deploy      []interface{} `json:"deploy"`
					BinFile     string        `json:"bin_file"`
					Envs        []interface{} `json:"envs"`
					HasBuild    bool          `json:"has_build"`
					BuildName   string        `json:"build_name"`
				} `json:"targets"`
				ArtifactArgs        []interface{} `json:"artifact_args"`
				Tests               []interface{} `json:"tests"`
				ReqID               string        `json:"req_id"`
				DistributeEnabled   bool          `json:"distribute_enabled"`
				WorkflowTaskCreator string        `json:"workflow_task_creator"`
				IgnoreCache         bool          `json:"ignore_cache"`
				ResetCache          bool          `json:"reset_cache"`
				NotificationID      string        `json:"notification_id"`
				MergeRequestID      string        `json:"merge_request_id"`
				CommitID            string        `json:"commit_id"`
				Source              string        `json:"source"`
				CodehostID          int           `json:"codehost_id"`
				RepoOwner           string        `json:"repo_owner"`
				RepoNamespace       string        `json:"repo_namespace"`
				RepoName            string        `json:"repo_name"`
				IsParallel          bool          `json:"is_parallel"`
				EnvName             string        `json:"env_name"`
				Callback            interface{}   `json:"callback"`
			} `json:"workflow_args"`
		} `json:"items"`
	} `json:"hook_ctl"`
	BaseName   string `json:"base_name"`
	ResetImage bool   `json:"reset_image"`
	IsParallel bool   `json:"is_parallel"`
}

//请求参数

type PickBuildOptions struct {
	WorkflowName string `json:"WorkflowName,omitempty"`
	PorjectName  string `json:"PorjectName,omitempty"`
}

//http://xxx.com/api/aslan/workflow/workflow/find/first-workflow?projectName=java-demo

func (p *pickBuildService) PickBuildInfo(opt *PickBuildOptions, options ...RequestOptionFunc) (*PickBuildResponse, *Response, error) {

	path := fmt.Sprintf("api/aslan/workflow/workflow/find/%s?projectName=%s", opt.WorkflowName, opt.PorjectName)

	req, err := p.client.NewRequest(http.MethodGet, path, nil, options)

	if err != nil {
		return nil, nil, err
	}

	pickBuild := new(PickBuildResponse)
	resp, err := p.client.Do(req, &pickBuild)
	if err != nil {
		return nil, resp, err
	}

	return pickBuild, resp, err
}
