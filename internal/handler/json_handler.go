package handler

import (
	"purekit-backend/errors"
	"purekit-backend/internal/service"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type JsonHandler struct {
	jsonService service.JsonService
}

func NewJsonHandler(jsonService service.JsonService) *JsonHandler {
	return &JsonHandler{
		jsonService: jsonService,
	}
}

type JsonFormatRequest struct {
	Data   string `json:"data" binding:"required"`
	Indent bool   `json:"indent"`
	Action string `json:"action"` // "format", "escape", "unescape"
}

func (h *JsonHandler) Format(c *gin.Context) {
	var req JsonFormatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.Error(c, errors.NewBadRequestError("Invalid request data"))
		return
	}

	var result string
	var err error

	switch req.Action {
	case "escape":
		result, err = h.jsonService.Escape(req.Data)
	case "unescape":
		result, err = h.jsonService.Unescape(req.Data)
	default:
		// 默认为格式化
		result, err = h.jsonService.Format(req.Data, req.Indent)
	}

	if err != nil {
		httputil.Error(c, errors.NewBadRequestError("JSON process failed: "+err.Error()))
		return
	}

	httputil.Success(c, gin.H{
		"result": result,
	})
}
