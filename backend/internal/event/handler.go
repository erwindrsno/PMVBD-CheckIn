package event

import (
	"net/http"

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

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.s.Create(req.Name); err != nil {
		responses.Fail(c, http.StatusBadRequest, err.Error())
	} else {
		responses.Success(c, http.StatusCreated, gin.H{"success": true})
	}
}

func (h *Handler) Read(c *gin.Context) {
	statusQuery := c.Query("status")

	filter := EventFilter{}
	if statusQuery != "" {
		if val, ok := statusMap[statusQuery]; ok {
			filter.Status = &val
		} else {
			// Handle invalid status (optional)
			responses.Fail(c, http.StatusBadRequest, "Invalid status value")
			return
		}
	}
	// events, err := h.s.Read(filter)
	datas, err := h.s.Read(filter)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
	responses.Success(c, http.StatusOK, gin.H{"events": datas})
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
	if err := h.s.UpdateStatus(id); err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
	if err := h.s.Delete(id); err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
}
