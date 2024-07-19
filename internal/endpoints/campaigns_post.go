package endpoints

import (
	"goapi/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	defer r.Body.Close()
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}
