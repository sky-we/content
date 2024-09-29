package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Content struct {
	ID             int64         `json:"id"`              // 内容ID
	Title          string        `json:"title"`           // 内容标题
	Description    string        `json:"description"`     // 内容描述
	Author         string        `json:"author"`          // 作者
	VideoURL       string        `json:"video_url"`       // 视频播放URL
	Thumbnail      string        `json:"thumbnail"`       // 封面图URL
	Category       string        `json:"category"`        // 内容分类
	Duration       time.Duration `json:"duration"`        // 内容时长
	Resolution     string        `json:"resolution"`      // 分辨率 如720p、1080p
	FileSize       int64         `json:"file_size"`       // 文件大小
	Format         string        `json:"format"`          // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `json:"updated_at"`      // 内容更新时间
	CreatedAt      time.Time     `json:"created_at"`      // 内容创建时间
}

type FindParams struct {
	ID       int64
	Author   string
	Title    string
	PageSize int64
	Page     int64
}

type ContentRepo interface {
	Create(ctx context.Context, content *Content) (int64, error)
	Update(ctx context.Context, id int64, content *Content) error
	UpdateCol(ctx context.Context, id int64, colName string, data any) error
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, params *FindParams) (*[]*Content, int64, error)
}

type ContentUseCase struct {
	repo ContentRepo
	log  *log.Helper
}

func NewContentUseCase(repo ContentRepo, logger log.Logger) *ContentUseCase {
	return &ContentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *ContentUseCase) CreateContent(ctx context.Context, content *Content) (int64, error) {
	c.log.WithContext(ctx).Infof("biz domain create content: [%+v]", content)
	return c.repo.Create(ctx, content)
}

func (c *ContentUseCase) UpdateContent(ctx context.Context, id int64, content *Content) error {
	c.log.WithContext(ctx).Infof("biz domain update content: id:[%d], [%+v]", id, content)
	return c.repo.Update(ctx, id, content)
}

func (c *ContentUseCase) UpdateContentCol(ctx context.Context, id int64, colName string, data any) error {
	c.log.WithContext(ctx).Infof("biz domain update content col: id:[%d] colName:[%s], data:[%+v]", id, colName, data)
	return c.repo.UpdateCol(ctx, id, colName, data)
}

func (c *ContentUseCase) DeleteContent(ctx context.Context, id int64) error {
	c.log.WithContext(ctx).Infof("biz domain del content: id:[%d]", id)
	return c.repo.Delete(ctx, id)
}

func (c *ContentUseCase) FindContent(ctx context.Context, params *FindParams) (*[]*Content, int64, error) {
	c.log.WithContext(ctx).Infof("biz domain find content: FindParams:[%v]", params)
	return c.repo.Find(ctx, params)
}
