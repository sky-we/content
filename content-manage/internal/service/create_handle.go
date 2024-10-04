package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
	"github.com/google/uuid"
	"time"
)

func (app *AppService) CreateContent(ctx context.Context, req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {
	content := req.GetContent()
	uc := app.uc
	ContentID, err := uc.CreateContent(ctx, &biz.Content{
		ContentId:      uuid.New().String(),
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
		Id: ContentID,
	}, nil
}
