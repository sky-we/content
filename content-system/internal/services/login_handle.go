package services

import (
	"content-system/internal/dao"
	"content-system/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginReq struct {
	UserId   string `json:"user_id" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
}
type LoginRsp struct {
	SessionId string `json:"session_id" binding:"required"`
	UserId    string `json:"user_id" binding:"required"`
	NickName  string `json:"nick_name" binding:"required"`
}

func (app *CmsApp) Login(ctx *gin.Context) {
	// zipkin span的context记录在gin ctx.Request里面（中间件实现）
	span := opentracing.SpanFromContext(ctx.Request.Context())
	var loginReq LoginReq

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if app.IsLogin(context.Background(), loginReq.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "用户已登录"})
		return
	}
	span.SetTag("reqInfo", loginReq)

	accountDao := dao.NewAccountDao(app.db)
	account, err := accountDao.FindByUserId(ctx.Request.Context(), loginReq.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "请输入正确的用户ID"})
		return

	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginReq.PassWord)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "用户密码错误"})
		return

	}
	sessionId, err := app.GenSessionId(ctx.Request.Context(), account.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "服务器内部错误", "err": err.Error()})

	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0,
		"msg": "login ok",
		"data": &LoginRsp{
			SessionId: sessionId,
			UserId:    account.UserId,
			NickName:  account.NickName,
		}})

}

func (app *CmsApp) GenSessionId(ctx context.Context, userId string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "[Login] GenSessionID")
	defer span.Finish()
	sessionId := uuid.New().String()
	sessionKey := utils.GenSessionKey(sessionId) // e.g. "session:{sessionId}"
	// 将 userId、生成时间、过期时间等存在一个 Hash 里
	data := map[string]interface{}{
		"userId":    userId,
		"createdAt": time.Now().Unix(),
	}
	if err := app.rdb.HSet(ctx, sessionKey, data).Err(); err != nil {
		return "", err
	}
	// 设置过期
	if err := app.rdb.Expire(ctx, sessionKey, 8*time.Hour).Err(); err != nil {
		return "", err
	}
	return sessionId, nil
}

func (app *CmsApp) IsLogin(ctx context.Context, userId string) bool {
	exists, err := app.rdb.Exists(ctx, utils.GenSessionKey(userId)).Result()
	if err != nil {
		panic(err)
	}
	return exists > 0

}
