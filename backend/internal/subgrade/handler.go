package subgrade

import (
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

func (h *Handler) Create(c *gin.Context) {
	// var req struct {
	// 	Name string `json:"name" binding:"required"`
	// }
	//
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	responses.Fail(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	//
	// if err := h.s.Create(req.Name); err != nil {
	// 	responses.Fail(c, http.StatusBadRequest, err.Error())
	// } else {
	// 	responses.Success(c, http.StatusCreated, gin.H{"success": true})
	// }
}

func (h *Handler) Read(c *gin.Context) {
	datas, err := h.s.Read()
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
	}
	responses.Success(c, http.StatusOK, gin.H{"subgrades": datas})
}
