package service

import (
	"content-manage/api/content/operate"
	"content-manage/internal/biz"
	"context"
	"time"
)

func (app *AppService) CreateContent(ctx context.Context, req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {

	content := req.GetContent()

	uc := app.uc

	newContentId, err := uc.CreateContent(ctx, &biz.Content{
		Title:          content.GetTitle(),
		VideoURL:       content.GetVideoURL(),
		Author:         content.GetAuthor(),
		Description:    content.GetDescription(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentRsp{
		Code: 0,
		Msg:  "create ok",
		Data: map[string]int64{"content_id": newContentId},
	}, nil
}
