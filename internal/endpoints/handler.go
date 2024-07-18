package endpoints

import "goapi/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
