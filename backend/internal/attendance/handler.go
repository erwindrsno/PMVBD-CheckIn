package attendance

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/attendee"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		slog.Error("actual error is", "error", err.Error())
		if errors.Is(err, attendee.ErrAttendeeNotExist) {
			responses.Fail(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, ErrDuplicateAttendance) {
			responses.Fail(c, http.StatusConflict, err.Error())
		} else {
			responses.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	responses.Success(c, http.StatusCreated, gin.H{"avi": *avi})
}

func (h *Handler) ReadByEventId(c *gin.Context) {
	eventIdParam := c.Param("event_id") // captures '?status=open'
	slog.Info("event id param is", "info", eventIdParam)

	atvis, err := h.s.ReadByEventId(eventIdParam)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.Success(c, http.StatusCreated, gin.H{"atvis": atvis})
}

func (h *Handler) DeleteById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		if errors.Is(err, ErrResourceNotFound) {
			responses.Fail(c, http.StatusNotFound, err.Error())
		} else {
			responses.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	if err := h.s.Delete(id); err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
}
