package service

import (
	"content-manage/api/operate"
	"content-manage/internal/biz"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (app *AppService) FindContent(ctx context.Context, req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	params := biz.FindParams{
		ID:       req.GetIdxID(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}
	uc := app.uc
	contentDetails, err := uc.FindContent(ctx, &params)
	if err != nil {
		return nil, err
	}

	var contents []*operate.Content
	for _, d := range *contentDetails {
		contents = append(contents, &operate.Content{
			ID:             d.ID,
			ContentID:      d.ContentId,
			Title:          d.Title,
			Description:    d.Description,
			Author:         d.Author,
			VideoURL:       d.VideoURL,
			Thumbnail:      d.Thumbnail,
			Category:       d.Category,
			Duration:       d.Duration.Nanoseconds(),
			Resolution:     d.Resolution,
			FileSize:       d.FileSize,
			Format:         d.Format,
			Quality:        d.Quality,
			ApprovalStatus: d.ApprovalStatus,
			CreatedAt:      timestamppb.New(d.CreatedAt),
			UpdatedAt:      timestamppb.New(d.UpdatedAt),
		})
	}
	return &operate.FindContentRsp{
		Content: contents,
	}, nil
}
