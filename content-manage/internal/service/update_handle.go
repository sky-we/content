package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
	"time"
)

func (app *AppService) UpdateContent(ctx context.Context, req *operate.UpdateContentReq) (*operate.UpdateContentRsp, error) {
	uc := app.uc
	idxID := req.GetIdxID()
	content := req.GetContent()
	err := uc.UpdateContent(ctx, idxID, &biz.Content{
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

	return &operate.UpdateContentRsp{}, nil
}
