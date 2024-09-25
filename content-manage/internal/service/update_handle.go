package service

import (
	"content-manage/api/content/operate"
	"content-manage/internal/biz"
	"context"
	"time"
)

func (app *AppService) UpdateContent(ctx context.Context, req *operate.UpdateContentReq) (*operate.UpdateContentRsp, error) {
	uc := app.uc
	content := req.GetContent()
	err := uc.UpdateContent(ctx, content.GetID(), &biz.Content{
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
		return &operate.UpdateContentRsp{
			Code: 2,
			Msg:  "fail",
			Data: map[string]int64{
				"ID": content.GetID(),
			},
		}, nil
	}

	return &operate.UpdateContentRsp{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"ID": content.GetID(),
		},
	}, nil
}
