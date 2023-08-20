package handler

import (
	"chatplus/core"
	"chatplus/service/oss"
	"chatplus/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadHandler struct {
	BaseHandler
	db              *gorm.DB
	uploaderManager *oss.UploaderManager
}

func NewUploadHandler(app *core.AppServer, db *gorm.DB, manager *oss.UploaderManager) *UploadHandler {
	handler := &UploadHandler{db: db, uploaderManager: manager}
	handler.App = app
	return handler
}

func (h *UploadHandler) Upload(c *gin.Context) {
	fileURL, err := h.uploaderManager.GetActiveService().PutFile(c)
	if err != nil {
		resp.ERROR(c, fmt.Sprintf("文件上传失败: %s", err.Error()))
		return
	}

	resp.SUCCESS(c, fileURL)
}