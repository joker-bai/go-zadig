# go-zadig

Zadig API 客户端，Go 程序可以简单的和 Zadig 进行交互。

## 覆盖范围

该 API 客户端只实现了部分功能，和 `https://docs.koderover.com/zadig/v1.15.0/api/usage/` 开发的 API 对应，如下：

- [x] 获取工作流任务状态
- [x] 工作流任务重试
- [x] 取消工作流任务
- [x] 执行工作流
- [x] 获取工作流任务详情
- [x] 获取数据概览
- [x] 获取构建数据统计
- [x] 获取部署数据统计
- [x] 获取测试数据统计
- [ ] 获取交付中心版本列表
- [ ] 获取交付物追踪信息
- [ ] 获取自定义工作流任务详情
- [ ] 执行自定义工作流
- [ ] 取消自定义工作流
- [ ] 自定义工作流人工审核

## 使用方式

```go
import "github.com/joker-bai/go-zadig"
```

例如获取工作流任务详情，样例如下：

```go
package main

import (
  "fmt"
  "log"

  "github.com/joker-bai/go-zadig"
)

var token = "xxxx"

func main() {
  client, err := zadig.NewClient(token, zadig.WithBaseURL("http://www.example.com/"))
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }


  w, r, err := client.Workflow.GetWorkflowTaskDetail(&zadig.GetWorkflowTaskDetailOptions{
    ID:           108,
    PipelineName: "pipelineName",
  })
  if err != nil {
    log.Fatalf("create workflow faield: %v", err)
  }
  fmt.Println(w, r.Body)
}
```
