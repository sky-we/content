package service

import (
	"content-manage/api/operate"
	"context"
)

func (app *AppService) DeleteContent(ctx context.Context, req *operate.DeleteContentReq) (*operate.DeleteContentRsp, error) {
	idxID := req.GetIdxID()
	uc := app.uc
	err := uc.DeleteContent(ctx, idxID)
	if err != nil {
		return nil, err
	}
	return &operate.DeleteContentRsp{}, nil
}
