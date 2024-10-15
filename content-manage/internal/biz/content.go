package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Content struct {
	ID             int64         `json:"id"`              // 自增ID
	ContentId      string        `json:"content_id"`      // 内容ID
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
type ContentIndex struct {
	ID        int64  `json:"id"`         // 自增ID
	ContentId string `json:"content_id"` // 内容ID
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
	Update(ctx context.Context, idxID int64, content *Content) error
	Delete(ctx context.Context, idxID int64) error
	Find(ctx context.Context, params *FindParams) (*[]*Content, error)
}

type ContentUseCase struct {
	repo ContentRepo
	log  *log.Helper
}

func NewContentUseCase(repo ContentRepo, logger log.Logger) *ContentUseCase {
	return &ContentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *ContentUseCase) CreateContent(ctx context.Context, content *Content) (int64, error) {

	c.log.WithContext(ctx).Infof("[biz domain] create content: [%+v]", content)
	return c.repo.Create(ctx, content)
}

func (c *ContentUseCase) UpdateContent(ctx context.Context, idxID int64, content *Content) error {
	c.log.WithContext(ctx).Infof("[biz domain] update content: id:[%d], [%+v]", idxID, content)
	return c.repo.Update(ctx, idxID, content)
}

func (c *ContentUseCase) DeleteContent(ctx context.Context, idxID int64) error {
	c.log.WithContext(ctx).Infof("[biz domain] del content: id:[%d]", idxID)
	return c.repo.Delete(ctx, idxID)
}

func (c *ContentUseCase) FindContent(ctx context.Context, params *FindParams) (*[]*Content, error) {
	c.log.WithContext(ctx).Infof("[biz domain] find content: FindParams:[%v]", params)
	return c.repo.Find(ctx, params)
}
