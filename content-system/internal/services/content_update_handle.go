package services

import (
	"content-system/internal/api/operate"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
	"time"
)

type ContentUpdateReq struct {
	ID             int64         `json:"id" binding:"required"` // 内容ID
	Title          string        `json:"title"`                 // 内容标题
	VideoURL       string        `json:"video_url"`             // 视频播放URL
	Author         string        `json:"author" `               // 作者
	Description    string        `json:"description"`           // 内容描述
	Thumbnail      string        `json:"thumbnail"`             // 封面图URL
	Category       string        `json:"category"`              // 内容分类
	Duration       time.Duration `json:"duration"`              // 内容时长
	Resolution     string        `json:"resolution"`            // 分辨率 如720p、1080p
	FileSize       int64         `json:"file_size"`             // 文件大小
	Format         string        `json:"format"`                // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`               // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"`       // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `json:"updated_at"`            // 内容更新时间
	CreatedAt      time.Time     `json:"created_at"`            // 内容创建时间
}

type ContentUpdateRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (app *CmsApp) ContentUpdate(ctx *gin.Context) {
	var contentUpdateReq ContentUpdateReq
	if err := ctx.ShouldBindJSON(&contentUpdateReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "参数错误", "err": err.Error()})
		return
	}

	_, err := app.operateAppClient.UpdateContent(context.Background(), &operate.UpdateContentReq{Content: &operate.Content{
		ID:             contentUpdateReq.ID,
		Title:          contentUpdateReq.Title,
		VideoURL:       contentUpdateReq.VideoURL,
		Author:         contentUpdateReq.Author,
		Description:    contentUpdateReq.Description,
		Thumbnail:      contentUpdateReq.Thumbnail,
		Category:       contentUpdateReq.Category,
		Duration:       contentUpdateReq.Duration.Milliseconds(),
		Resolution:     contentUpdateReq.Resolution,
		FileSize:       contentUpdateReq.FileSize,
		Format:         contentUpdateReq.Format,
		Quality:        contentUpdateReq.Quality,
		ApprovalStatus: contentUpdateReq.ApprovalStatus,
	}})

	if err != nil {
		ctx.AbortWithStatusJSON(errors.Code(err), gin.H{"Message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &ContentUpdateRsp{
		Code:    0,
		Message: "success",
		Data:    fmt.Sprintf("content id [%d] updated", contentUpdateReq.ID),
	})

}
