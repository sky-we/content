package process

import (
	"content-flow/internal/api/operate"
	"content-flow/internal/api/utils"
	"content-flow/internal/middleware"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	flow "github.com/s8sg/goflow/flow/v1"
)

type ContentFlow struct {
	client operate.AppClient
}

var Logger = middleware.GetLogger()

func NewContentFlow(client operate.AppClient) *ContentFlow {
	return &ContentFlow{client: client}
}

func (c *ContentFlow) ContentFlowHandle(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	dag.Node("input", c.input)
	dag.Node("verify", c.verify)
	dag.Node("finish", c.finish)
	// 定义所有的分支类型
	branches := dag.ConditionalBranch("branches", []string{"category", "thumbnail", "pass", "format", "fail"},
		// 根据审核状态，返回分支类型
		func(bytes []byte) []string {
			var data map[string]interface{}
			if err := json.Unmarshal(bytes, &data); err != nil {
				return nil
			}
			if data["ApprovalStatus"].(float64) == 2 {
				return []string{"category", "thumbnail", "pass", "format"}
			}
			return []string{"fail"}
			// 分支结果聚合
		}, flow.Aggregator(func(m map[string][]byte) ([]byte, error) {
			return []byte("ok"), nil
		}),
	)
	branches["category"].Node("category", c.category)
	branches["thumbnail"].Node("thumbnail", c.thumbnail)
	branches["pass"].Node("category", c.pass)
	branches["format"].Node("format", c.format)
	branches["fail"].Node("fail", c.fail)

	dag.Edge("input", "verify")
	dag.Edge("verify", "branches")
	dag.Edge("branches", "finish")

	return nil

}

func (c *ContentFlow) input(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec input node...")

	var d map[string]int
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	id := int64(d["input"])
	detail, err := c.client.FindContent(context.Background(), &operate.FindContentReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	if detail.Content == nil {
		return nil, errors.New(fmt.Sprintf("ID [%d] content detail not found", id))
	}
	result, err := json.Marshal(map[string]interface{}{
		"title":     detail.Content[0].Title,
		"video_url": detail.Content[0].VideoURL,
		"id":        detail.Content[0].ID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *ContentFlow) verify(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec verify node...")
	var detail map[string]interface{}

	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}
	var (
		title    = detail["title"]
		videoUrl = detail["video_url"]
		id       = detail["id"]
	)
	if int(id.(float64))%2 == 0 {
		detail["ApprovalStatus"] = 3
	} else {
		detail["ApprovalStatus"] = 2
	}
	Logger.Info(id, title, videoUrl)
	return json.Marshal(detail)
}

func (c *ContentFlow) category(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec category node...")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int64(input["id"].(float64))
	err := c.updateColById(contentId, "Category", "category-workflow")
	if err != nil {
		return nil, err
	}
	return []byte("category"), nil
}
func (c *ContentFlow) thumbnail(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec thumbnail node...")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int64(input["id"].(float64))
	err := c.updateColById(contentId, "Thumbnail", "thumbnail-workflow")
	if err != nil {
		return nil, err
	}
	return []byte("thumbnail"), nil
}
func (c *ContentFlow) format(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec format node...")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int64(input["id"].(float64))
	err := c.updateColById(contentId, "Format", "format-workflow")
	if err != nil {
		return nil, err
	}
	return []byte("format"), nil
}
func (c *ContentFlow) pass(data []byte, option map[string][]string) ([]byte, error) {
	Logger.Info("exec pass node...")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int64(input["id"].(float64))
	// 审核成功
	if err := c.updateColById(contentID, "ApprovalStatus", int32(2)); err != nil {
		return nil, err
	}
	return data, nil
}
func (c *ContentFlow) fail(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec fail node...")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int64(input["id"].(float64))
	// 审核失败

	if err := c.updateColById(contentId, "ApprovalStatus", int32(3)); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *ContentFlow) finish(data []byte, options map[string][]string) ([]byte, error) {
	Logger.Info("exec finish node...")
	Logger.Info(string(data))
	return data, nil
}

func (c *ContentFlow) updateColById(contentId int64, colName string, data any) error {

	content := &operate.Content{
		ID:             contentId,
		Title:          "",
		VideoURL:       "",
		Author:         "",
		Description:    "",
		Thumbnail:      "",
		Category:       "",
		Duration:       0,
		Resolution:     "",
		FileSize:       0,
		Format:         "",
		Quality:        0,
		ApprovalStatus: 0,
	}

	if err := utils.UpdateStructField(content, colName, data); err != nil {
		return err
	}

	_, err := c.client.UpdateContent(context.Background(), &operate.UpdateContentReq{
		Content: content,
	}, nil)

	if err != nil {
		return err
	}
	return nil
}
