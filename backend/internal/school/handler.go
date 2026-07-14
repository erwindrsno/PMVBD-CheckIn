package school

import (
	"fmt"
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

func (h *Handler) Read(c *gin.Context) {
	datas, err := h.s.st.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	responses.Success(c, http.StatusOK, gin.H{"schools": datas})
}
