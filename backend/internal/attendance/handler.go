package attendance

import (
	"errors"
	"net/http"

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

func (h *Handler) Capture(c *gin.Context) {
	var req AttendanceCapturePayload

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	avi, err := h.s.Capture(req)
	if err != nil {
		if errors.Is(err, ErrAttendeeNotExist) {
			responses.Fail(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(err, ErrDuplicateAttendance) {
			responses.Fail(c, http.StatusConflict, err.Error())
		} else {
			responses.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	responses.Success(c, http.StatusCreated, gin.H{"avi": *avi})
}
