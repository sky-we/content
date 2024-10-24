package services

import (
	"content-system/internal/api/operate"
	"content-system/internal/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ContentCreateReq struct {
	Title          string        `json:"title" binding:"required"`     // 内容标题
	VideoURL       string        `json:"video_url" binding:"required"` // 视频播放URL
	Author         string        `json:"author" binding:"required"`    // 作者
	Description    string        `json:"description"`                  // 内容描述
	Thumbnail      string        `json:"thumbnail"`                    // 封面图URL
	Category       string        `json:"category"`                     // 内容分类
	Duration       time.Duration `json:"duration"`                     // 内容时长
	Resolution     string        `json:"resolution"`                   // 分辨率 如720p、1080p
	FileSize       int64         `json:"file_size"`                    // 文件大小
	Format         string        `json:"format"`                       // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`                      // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"`              // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `json:"updated_at"`                   // 内容更新时间
	CreatedAt      time.Time     `json:"created_at"`                   // 内容创建时间
}

type ContentCreateRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (app *CmsApp) ContentCreate(ctx *gin.Context) {
	Logger.Info("[content-system|Gin service] create content start")
	var contentCreateReq ContentCreateReq

	// 入参校验
	if err := ctx.ShouldBindJSON(&contentCreateReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "参数错误", "error": err.Error()})
		return
	}
	Logger.Info("[content-system|Invoke Grpc Service] operateAppClient.CreateContent start")
	rsp, err := app.operateAppClient.CreateContent(context.Background(), &operate.CreateContentReq{Content: &operate.Content{
		ContentID:      uuid.New().String(),
		Title:          contentCreateReq.Title,
		VideoURL:       contentCreateReq.VideoURL,
		Author:         contentCreateReq.Author,
		Description:    contentCreateReq.Description,
		Thumbnail:      contentCreateReq.Thumbnail,
		Category:       contentCreateReq.Category,
		Duration:       contentCreateReq.Duration.Nanoseconds(),
		Resolution:     contentCreateReq.Resolution,
		FileSize:       contentCreateReq.FileSize,
		Format:         contentCreateReq.Format,
		Quality:        contentCreateReq.Quality,
		ApprovalStatus: contentCreateReq.ApprovalStatus,
	}})
	Logger.Info("[content-system|Invoke Grpc Service] operateAppClient.CreateContent end")

	if err != nil {
		ctx.AbortWithStatusJSON(errors.Code(err), gin.H{"Message": err.Error()})
		return
	}
	Logger.Info("[content-system|Invoke Go-Flow server] startContentFlow start")

	// 数据加工开始
	go func() {
		if err := app.startContentFlow(rsp.IdxID, config.ClientCfg.FlowServiceClient); err != nil {
			Logger.Errorf("start content flow error %v", err)
		}
	}()
	type ContentID struct {
		ID int64
	}
	ctx.JSON(http.StatusOK, &ContentCreateRsp{
		Code:    0,
		Message: "success",
		Data:    ContentID{ID: rsp.IdxID},
	})
	Logger.Info("[content-system|Gin service] create content end")
}
