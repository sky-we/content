package data

import (
	"content-manage/internal/biz"
	"content-manage/utils"
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
	ID             int64         `gorm:"column:id;primary_key"`  // 自增ID
	ContentId      string        `gorm:"column:content_id"`      // 内容ID
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

type IdxContentDetail struct {
	ID        int64     `gorm:"column:id;primary_key"` // 自增ID
	ContentID string    `gorm:"column:content_id"`     // 内容ID
	Title     string    `gorm:"column:title"`          // 内容标题
	Author    string    `gorm:"column:author"`         // 作者
	UpdatedAt time.Time `gorm:"column:updated_at"`     // 内容更新时间
	CreatedAt time.Time `gorm:"column:created_at"`     // 内容创建时间
}

func (c IdxContentDetail) TableName() string {
	return "cms_content.idx_content_details"
}

func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func getShardTableName(contentID string) string {
	idx := utils.GenIdx(contentID, 4)
	log.Infof("content_id = %s, tableIdx = %d", contentID, idx)
	return fmt.Sprintf("cms_content.content_details_%d", idx)
}
func (c *contentRepo) Create(ctx context.Context, content *biz.Content) (int64, error) {
	db := c.data.db
	tableName := getShardTableName(content.ContentId)
	idx := IdxContentDetail{
		ContentID: content.ContentId,
		Title:     content.Title,
		Author:    content.Author,
	}

	if err := db.Create(&idx).Error; err != nil {
		return 0, err
	}

	repeat, err := c.IsVideoRepeat(ctx, tableName, content.VideoURL)
	if err != nil {
		return 0, errors.New(http.StatusInternalServerError, "SERVER_INTER_ERROR", fmt.Sprintf("创建失败，内部服务器错误：%s", err))
	}

	if repeat {
		return 0, errors.New(http.StatusBadRequest, "VIDEO_URL_REPEAT", fmt.Sprintf("[video_url=%s]内容已存在", content.VideoURL))
	}
	detail := &ContentDetail{
		ContentId:      content.ContentId,
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
	//if err := db.Create(&detail).Error; err != nil {
	//	return 0, errors.New(http.StatusInternalServerError, "Create Content Failed", err.Error())
	//}
	if err := db.Table(tableName).Create(&detail).Error; err != nil {
		return 0, errors.New(http.StatusInternalServerError, "Create Content Failed", err.Error())
	}
	return detail.ID, nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	db := c.data.db

	var idx IdxContentDetail
	if err := db.Where("id = ?", id).First(&idx).Error; err != nil {
		log.Infof("get content id by ID %d", idx.ID)
		return errors.New(http.StatusInternalServerError, "Get Content ID Failed", err.Error())
	}
	tableName := getShardTableName(idx.ContentID)

	exists, err := c.IsContentExist(ctx, tableName, idx.ContentID)

	if err != nil {
		return errors.New(http.StatusInternalServerError, "Update Content Failed", err.Error())
	}
	if !exists {

		return errors.New(http.StatusBadRequest, "ID_NOT_EXIST", fmt.Sprintf("content ID %d is not exist", id))
	}
	detail := &ContentDetail{
		ContentId:      content.ContentId,
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
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}
	if err := db.Table(tableName).Where("id = ?", id).Updates(&detail).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "Update Content Failed", err.Error())
	}
	return nil
}

func (c *contentRepo) UpdateCol(ctx context.Context, id int64, colName string, data any) error {
	db := c.data.db
	var idx IdxContentDetail
	if err := db.Where("id = ?", id).First(&idx).Error; err != nil {
		log.Infof("get content id by ID %d", idx.ID)
		return errors.New(http.StatusInternalServerError, "Get Content ID Failed", err.Error())
	}
	tableName := getShardTableName(idx.ContentID)

	exists, err := c.IsContentExist(ctx, tableName, idx.ContentID)
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

func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	var idx IdxContentDetail
	if err := db.Where("id = ?", id).First(&idx).Error; err != nil {
		log.Infof("get content id by ID %d", idx.ID)
		return errors.New(http.StatusInternalServerError, "Get Content ID Failed", err.Error())
	}
	tableName := getShardTableName(idx.ContentID)
	exists, err := c.IsContentExist(ctx, tableName, idx.ContentID)
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

func (c *contentRepo) IsContentExist(ctx context.Context, tableName string, ContentID string) (bool, error) {
	db := c.data.db
	if err := db.Table(tableName).Where("content_id = ?", ContentID).First(&ContentDetail{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New(http.StatusInternalServerError, "Query Content By ID Failed", err.Error())

		}
		return false, nil
	}
	return true, nil
}

func (c *contentRepo) IsVideoRepeat(ctx context.Context, tableName string, videoUrl string) (bool, error) {
	db := c.data.db
	var contentDetail ContentDetail
	err := db.Table(tableName).Where("video_url = ?", videoUrl).First(&contentDetail).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, errors.New(http.StatusInternalServerError, "Query Content By video_url Failed", err.Error())
	}

	return true, nil
}

func (c *contentRepo) First(ctx context.Context, contentId string) (*ContentDetail, error) {
	db := c.data.db
	var detail ContentDetail
	tableName := getShardTableName(contentId)

	err := db.Table(tableName).Where("content_id = ?", contentId).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return &detail, nil
	}
	return &detail, nil
}
func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) (*[]*biz.Content, int64, error) {
	contentIdxMap, total, err := c.FindContentId(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	var contents []*biz.Content
	for _, idxMap := range contentIdxMap {
		contentDetail, getContentDetailErr := c.First(ctx, idxMap.ContentId)
		if getContentDetailErr != nil {
			return nil, 0, getContentDetailErr
		}
		contents = append(contents, &biz.Content{
			Title:          contentDetail.Title,
			Description:    contentDetail.Description,
			Author:         contentDetail.Author,
			VideoURL:       contentDetail.VideoURL,
			Thumbnail:      contentDetail.Thumbnail,
			Category:       contentDetail.Category,
			Duration:       contentDetail.Duration,
			Resolution:     contentDetail.Resolution,
			FileSize:       contentDetail.FileSize,
			Format:         contentDetail.Format,
			Quality:        contentDetail.Quality,
			ApprovalStatus: contentDetail.ApprovalStatus,
			UpdatedAt:      contentDetail.UpdatedAt,
			CreatedAt:      contentDetail.CreatedAt,
		})
	}
	return &contents, total, nil
}

func (c *contentRepo) FindContentId(ctx context.Context, params *biz.FindParams) ([]*biz.ContentIndex, int64, error) {
	db := c.data.db
	query := db.Model(&IdxContentDetail{})
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
	var IdxContentDetails []*IdxContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&IdxContentDetails).Error; err != nil {
		return nil, 0, errors.New(http.StatusInternalServerError, "Find Content Failed", err.Error())
	}
	var contentIdx []*biz.ContentIndex
	for _, detail := range IdxContentDetails {
		contentIdx = append(contentIdx, &biz.ContentIndex{
			ContentId: detail.ContentID,
			ID:        detail.ID,
		})
	}
	return contentIdx, cnt, nil
}
