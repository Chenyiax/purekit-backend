package handler

import (
	"strconv"

	"purekit-backend/errors"
	"purekit-backend/internal/service"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type PasswordHandler struct {
	passwordService service.PasswordService
}

func NewPasswordHandler(passwordService service.PasswordService) *PasswordHandler {
	return &PasswordHandler{
		passwordService: passwordService,
	}
}

func (h *PasswordHandler) Generate(c *gin.Context) {
	lengthStr := c.DefaultQuery("length", "16")
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length < 4 || length > 128 {
		length = 16
	}

	includeUpper := c.DefaultQuery("upper", "true") == "true"
	includeLower := c.DefaultQuery("lower", "true") == "true"
	includeNumber := c.DefaultQuery("number", "true") == "true"
	includeSymbol := c.DefaultQuery("symbol", "false") == "true"

	password, err := h.passwordService.Generate(length, includeUpper, includeLower, includeNumber, includeSymbol)
	if err != nil {
		httputil.Error(c, errors.NewInternalServerError("Failed to generate password", err))
		return
	}

	httputil.Success(c, gin.H{
		"password": password,
	})
}
