package endpoints

import (
	"errors"
	"goapi/internal/contract"
	internalerrors "goapi/internal/internalErrors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)
	if err != nil {

		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
