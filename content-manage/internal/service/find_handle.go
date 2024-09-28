package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
)

func (app *AppService) FindContent(ctx context.Context, req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	params := biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}
	uc := app.uc
	contentDetails, total, err := uc.FindContent(ctx, &params)
	if err != nil {
		return nil, err
	}

	var contents []*operate.Content
	for _, d := range *contentDetails {
		contents = append(contents, &operate.Content{
			ID:             d.ID,
			Title:          d.Title,
			Description:    d.Description,
			Author:         d.Author,
			VideoURL:       d.VideoURL,
			Thumbnail:      d.Thumbnail,
			Category:       d.Category,
			Duration:       d.Duration.Milliseconds(),
			Resolution:     d.Resolution,
			FileSize:       d.FileSize,
			Format:         d.Format,
			Quality:        d.Quality,
			ApprovalStatus: d.ApprovalStatus,
		})
	}
	return &operate.FindContentRsp{
		Content: contents,
		Total:   total,
	}, nil
}
