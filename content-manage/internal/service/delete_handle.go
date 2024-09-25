package service

import (
	"content-manage/api/content/operate"
	"context"
)

func (app *AppService) DeleteContent(ctx context.Context, req *operate.DeleteContentReq) (*operate.DeleteContentRsp, error) {
	id := req.GetId()
	uc := app.uc
	err := uc.DeleteContent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &operate.DeleteContentRsp{
		Code: 0,
		Msg:  "delete ok",
		Data: map[string]int64{"content_id": id},
	}, nil
}
