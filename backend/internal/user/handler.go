package user

import (
	"errors"
	"net/http"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/auth"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/responses"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	h := &Handler{
		s: s,
	}
	return h
}

func (h *Handler) Create(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.Create(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
	} else {
		responses.Success(c, http.StatusCreated, gin.H{"success": true})
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req User

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.Login(&req); err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) || errors.Is(err, ErrInvalidCredentials) {
			responses.Fail(c, http.StatusUnauthorized, err.Error())
		} else {
			responses.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	token, err := auth.GenerateToken()
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.Success(c, http.StatusOK, gin.H{"token": token})
}
