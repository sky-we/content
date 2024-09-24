package service

import (
	"content-manage/api/content/operate"
	"content-manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer
	uc *biz.ContentUseCase
}

func NewAppService(uc *biz.ContentUseCase) *AppService {
	return &AppService{uc: uc}

}
