package handler

import (
	"purekit-backend/errors"
	"purekit-backend/internal/service"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type TextHandler struct {
	textService service.TextService
}

func NewTextHandler(textService service.TextService) *TextHandler {
	return &TextHandler{
		textService: textService,
	}
}

type TextProcessRequest struct {
	Text   string `json:"text" binding:"required"`
	Action string `json:"action" binding:"required"`
}

func (h *TextHandler) Process(c *gin.Context) {
	var req TextProcessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.Error(c, errors.NewBadRequestError("Invalid request data"))
		return
	}

	result, stats, err := h.textService.Process(req.Text, req.Action)
	if err != nil {
		httputil.Error(c, errors.NewInternalServerError("Failed to process text", err))
		return
	}

	httputil.Success(c, gin.H{
		"result": result,
		"stats":  stats,
	})
}
