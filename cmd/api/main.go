package main

import (
	"goapi/internal/domain/campaign"
	"goapi/internal/endpoints"
	databse "goapi/internal/infraStructure/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	campaignService := campaign.Service{
		Repository: &databse.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	//routes
	r.Post("/campaigns", endpoints.HandlerError(handler.CreateCampaign))
	r.Get("/campaigns", endpoints.HandlerError(handler.GetCampaign))
	log.Println("Server starting on port 3000")
	http.ListenAndServe(":3000", r)
}
