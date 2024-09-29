package services

import (
	"content-system/internal/api/operate"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type ContentDeleteReq struct {
	ID int64 `json:"id" binding:"required"`
}

type ContentDeleteRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (app *CmsApp) ContentDelete(ctx *gin.Context) {
	var contentDeleteReq ContentDeleteReq
	if err := ctx.ShouldBindJSON(&contentDeleteReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "参数错误", "error": err.Error()})
		return
	}
	_, err := app.operateAppClient.DeleteContent(context.Background(), &operate.DeleteContentReq{Id: contentDeleteReq.ID})
	if err != nil {
		ctx.AbortWithStatusJSON(errors.Code(err), gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ContentDeleteRsp{
		Code: 0,
		Msg:  "success",
		Data: fmt.Sprintf("Content ID %d deleted", contentDeleteReq.ID),
	})

}
