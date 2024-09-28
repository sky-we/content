package data

import (
	"content-manage/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
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

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) (contentID int64, err *errors.Error) {
	db := c.data.db
	repeat, err := c.IsVideoRepeat(ctx, content.VideoURL)
	if err != nil {
		return 0, errors.New(http.StatusInternalServerError, "SERVER_INTER_ERROR", fmt.Sprintf("创建失败，内部服务器错误：%s", err))
	}

	if repeat {
		return 0, errors.New(http.StatusBadRequest, "VIDEO_URL_REPEAT", fmt.Sprintf("[video_url=%s]内容已存在", content.VideoURL))
	}
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
		return 0, errors.New(http.StatusInternalServerError, "Create Content Failed", err.Error())
	}
	return detail.ID, nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) *errors.Error {
	db := c.data.db
	exists, err := c.IsExist(ctx, id)
	if err != nil {
		return errors.New(http.StatusInternalServerError, "Update Content Failed", err.Error())
	}
	if !exists {

		return errors.New(http.StatusBadRequest, "ID_NOT_EXIST", fmt.Sprintf("content ID %d is not exist", id))
	}
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
		return errors.New(http.StatusInternalServerError, "Update Content Failed", err.Error())
	}
	return nil
}

func (c *contentRepo) UpdateCol(ctx context.Context, id int64, colName string, data any) *errors.Error {
	db := c.data.db
	exists, err := c.IsExist(ctx, id)
	if err != nil {
		return errors.New(http.StatusInternalServerError, "Update Content col Failed", err.Error())
	}
	if !exists {
		return errors.New(http.StatusBadRequest, "ID_NOT_EXIST", fmt.Sprintf("content ID %d is not exist", id))
	}
	if err := db.Where("id = ", id).Update(colName, data).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "Update Content col Failed", err.Error())
	}
	return nil
}

func (c *contentRepo) Delete(ctx context.Context, id int64) *errors.Error {
	db := c.data.db
	exists, err := c.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !exists {

		return errors.New(http.StatusBadRequest, "ID_NOT_EXIST", fmt.Sprintf("content ID %d is not exist", id))
	}
	if err := db.Where("id = ?", id).Delete(&ContentDetail{}).Error; err != nil {

		return errors.New(http.StatusInternalServerError, "DELETE_FAIL", "Delete failed")
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, *errors.Error) {
	db := c.data.db
	if err := db.Where("id = ?", id).First(&ContentDetail{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New(http.StatusInternalServerError, "Query Content By ID Failed", err.Error())

		}
		return false, nil
	}
	return true, nil
}

func (c *contentRepo) IsVideoRepeat(ctx context.Context, videoUrl string) (bool, *errors.Error) {
	db := c.data.db
	var contentDetail ContentDetail
	err := db.Where("video_url = ?", videoUrl).First(&contentDetail).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, errors.New(http.StatusInternalServerError, "Query Content By video_url Failed", err.Error())
	}

	return true, nil
}

func (c *contentRepo) First(ctx context.Context, id int64) (*ContentDetail, *errors.Error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return &detail, nil
	}
	return &detail, nil
}

func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) (*[]*biz.Content, int64, *errors.Error) {
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
		return nil, 0, errors.New(http.StatusInternalServerError, "Find Content Failed", err.Error())
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
