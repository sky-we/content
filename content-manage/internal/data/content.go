package data

import (
	"content-manage/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  // 内容ID
	Title          string        `gorm:"column:title"`           // 内容标题
	Description    string        `gorm:"column:description"`     // 内容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 视频播放URL
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面图URL
	Category       string        `gorm:"column:category"`        // 内容分类
	Duration       time.Duration `gorm:"column:duration"`        // 内容时长
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 如720p、1080p
	FileSize       int64         `gorm:"column:file_size"`       // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式 如MP4、AVI
	Quality        int32         `gorm:"column:quality"`         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `gorm:"column:approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      // 内容更新时间
	CreatedAt      time.Time     `gorm:"column:created_at"`      // 内容创建时间
}

func (c *ContentDetail) TableName() string {
	return "cms_content.content_details"

}

func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) (contentID int64, err error) {
	db := c.data.db
	detail := &ContentDetail{
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
		UpdatedAt:      time.Time{},
		CreatedAt:      time.Time{},
	}
	if err := db.Create(&detail).Error; err != nil {
		return 0, err
	}
	return detail.ID, nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	db := c.data.db
	detail := &ContentDetail{
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
		UpdatedAt:      time.Time{},
		CreatedAt:      time.Time{},
	}
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		return err
	}
	return nil
}

func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	exists, err := c.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(400, "ID_NOT_EXIST", fmt.Sprintf("id %d is not exist", id))
	}

	if err := db.Where("id = ?", id).Delete(&ContentDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	if err := db.Where("id = ?", id).First(&ContentDetail{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) (*[]*biz.Content, int64, error) {
	db := c.data.db
	query := db.Model(&ContentDetail{})
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Author != "" {
		query = query.Where("author = ?", params.Author)
	}
	if params.Title != "" {
		query = query.Where("title = ?", params.Title)
	}

	var cnt int64
	query.Count(&cnt)

	var page, pageSize = 1, 10

	if params.Page != 0 {
		page = int(params.Page)
	}
	if params.PageSize != 0 {
		pageSize = int(params.PageSize)
	}
	offset := (page - 1) * pageSize
	var contentDetails []*ContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&contentDetails).Error; err != nil {
		return nil, 0, err
	}
	var bizContent []*biz.Content
	for _, detail := range contentDetails {
		bizContent = append(bizContent, &biz.Content{
			ID:             detail.ID,
			Title:          detail.Title,
			Description:    detail.Description,
			Author:         detail.Author,
			VideoURL:       detail.VideoURL,
			Thumbnail:      detail.Thumbnail,
			Category:       detail.Category,
			Duration:       detail.Duration,
			Resolution:     detail.Resolution,
			FileSize:       detail.FileSize,
			Format:         detail.Format,
			Quality:        detail.Quality,
			ApprovalStatus: detail.ApprovalStatus,
		})
	}
	return &bizContent, cnt, nil
}
