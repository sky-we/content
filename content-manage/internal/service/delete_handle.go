package service

import (
	"content-manage/api/operate"
	"context"
)

func (app *AppService) DeleteContent(ctx context.Context, req *operate.DeleteContentReq) (*operate.DeleteContentRsp, error) {
	id := req.GetId()
	uc := app.uc
	err := uc.DeleteContent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &operate.DeleteContentRsp{}, nil
}
