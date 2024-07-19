package endpoints

import (
	"net/http"
)

func (h *Handler) GetCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	defer r.Body.Close()
	campaigns, err := h.CampaignService.Repository.Get()
	return campaigns, 200, err
}
