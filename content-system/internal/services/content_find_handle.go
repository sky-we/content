package services

import (
	"content-system/internal/api/operate"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
	"time"
)

type Content struct {
	ID             int           // 内容ID
	Title          string        // 内容标题
	Description    string        // 内容描述
	Author         string        // 作者
	VideoURL       string        // 视频播放URL
	Thumbnail      string        // 封面图URL
	Category       string        // 内容分类
	Duration       time.Duration // 内容时长
	Resolution     string        // 分辨率 如720p、1080p
	FileSize       int64         // 文件大小
	Format         string        // 文件格式 如MP4、AVI
	Quality        int32         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         // 审核状态 1-审核中 2-审核通过 3-审核不通过
}
type ContentFindReq struct {
	ID       int64  `json:"id"`       // 内容ID
	Author   string `json:"author"`   // 内容作者
	Title    string `json:"title"`    // 内容标题
	Page     int64  `json:"page"`     // 页数
	PageSize int64  `json:"pageSize"` // 页大小
}

type ContentFindRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    gin.H  `json:"data"`
}

func (app *CmsApp) ContentFind(ctx *gin.Context) {
	var contentFindReq ContentFindReq

	// 入参校验
	if err := ctx.ShouldBindJSON(&contentFindReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "参数错误", "error": err.Error()})
		return
	}
	rsp, err := app.operateAppClient.FindContent(context.Background(), &operate.FindContentReq{
		Id:       contentFindReq.ID,
		Title:    contentFindReq.Title,
		Author:   contentFindReq.Author,
		Page:     contentFindReq.Page,
		PageSize: contentFindReq.PageSize,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(errors.Code(err), gin.H{"Message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &ContentFindRsp{
		Code:    0,
		Message: "success",
		Data: gin.H{
			"content": rsp.Content,
			"total":   rsp.Total,
		},
	})

}
